package controllers
import (
	"github.com/gin-gonic/gin"
	"github.com/pasarsakomandiri/models"
	"database/sql"
	"net/http"
	"github.com/pasarsakomandiri/shared/response"
	"strings"
	"github.com/pasarsakomandiri/shared/tools"
	"time"
	"github.com/pasarsakomandiri/api"
	"github.com/jinzhu/now"
	"github.com/pasarsakomandiri/shared/session"
	"log"
	"strconv"
	"errors"
	"io/ioutil"
	"math"
	"os"
    "sync"
)

type ParkingResponse struct {
	Response *response.SimpleResponse
	Data models.ParkingTicket
	DeltaTime string
	Picture_path_in string
	Picture_path_out string
}

var mux sync.Mutex

func ParkingTicketPage(c *gin.Context) {
    session := session.Instance(c)
	c.HTML(http.StatusFound, "ticket_test.tmpl", gin.H{"title":"Ticket Test", "token":session.Get("token")})
}

func ParkingTransTgl(c *gin.Context)  {
	created_date := c.Query("created_date")

	tglparking, err := models.ParkingTransGetByTgl(created_date)

	if err != nil{
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("failed", err.Error()))
		return
	}

	c.JSON(http.StatusOK, tglparking)
	return
}

func ParkingCheckIn(c *gin.Context) {
	//init others
	date := time.Now()
	slicedIp := strings.Split(c.ClientIP(), ":")
	deviceIp := slicedIp[0]
	deviceToken := c.PostForm("token")

	//get raspberry device
	device, err := models.DeviceGetByHost(deviceIp)
	//create parking response struct
	parkingResponse := ParkingResponse{}
	response := new(response.SimpleResponse)
	parkingResponse.Response = response
	//check if device exists
	if err != nil {
		if err ==  sql.ErrNoRows {
			response.Status = "Failed"; response.Message = "Device not found"
			c.JSON(http.StatusOK, parkingResponse)
			return
		}

		log.Println("Db error: ", err)
		return
	}
	//check if device is a raspberry
	if !strings.EqualFold("Raspberry", device.Device_type){
		response.Status = "Failed"; response.Message = "Device is not valid"
		c.JSON(http.StatusOK, parkingResponse)
		return
	}
	//check if device token match
	if !strings.EqualFold(deviceToken, device.Token) {
		response.Status = "Failed"; response.Message = "Token not match"
		c.JSON(http.StatusOK, parkingResponse)
		return
	}
	//get device group of raspberry
	deviceGroup, err := models.DeviceGroupGetByHost("raspberry_ip", deviceIp)
	//check if device group not found
	if err == sql.ErrNoRows {
		response.Status = "Failed"; response.Message = "Device group not found"
		c.JSON(http.StatusOK, parkingResponse)
		return
	}
	//set ipcamera param
	camChan := make(chan []byte)
	ipCamera := api.IpCamera{}
	ipCamera.Protocol = "http"
	ipCamera.Param = "Streaming/channels/1/picture"
	ipCamera.Host = deviceGroup.Camera_ip
	ipCamera.Username = "admin"
	ipCamera.Password = "sakomandiri1"
	go ipCamera.GetPictureWithChannel(camChan)

	//parking ticket struct
	parkingTicket := models.ParkingTicket{}
	parkingTicket.Created_by = device.Id
	parkingTicket.Created_date = date.String()
	year := strconv.Itoa(date.Year())
	year = year[2:4]
	month := date.Month().String()
	month = month[0:3]
	parkingTicket.Vehicle_id = deviceGroup.Vehicle_id
	parkingTicket.Vehicle_type = deviceGroup.Vehicle_type

	//generate ticket number
	//loop until parking ticket number is not exists
    var result sql.Result
	ticketExists := true
	for ticketExists {
        mux.Lock()
		parkingTicket.Ticket_number = tools.RandomString(10)
		parkingTicket.Ticket_number =  parkingTicket.Ticket_number
		ticketExists = isTicketNumberExists(c, parkingTicket.Ticket_number)
        
        if !ticketExists {
            result, err = models.ParkingCreateNewTicket(parkingTicket)
            
            if err != nil {
                response.Status = "Failed"; response.Message = "Cannot create ticket to db"
                c.JSON(http.StatusOK, parkingResponse)
                return
            }
        }
        
        mux.Unlock()
	}

	//save ticket and picture to database
	//ticket saved to db = print, esle don't print
	/*result, err := models.ParkingCreateNewTicket(parkingTicket)

	if err != nil {
		response.Status = "Failed"; response.Message = "Cannot create ticket to db"
		c.JSON(http.StatusOK, parkingResponse)
		return
	}*/

	parkingTicket.Id, _ = result.LastInsertId()
    //taking picture from camera with goroutines, reducing latency
	go saveCameraPicture(<-camChan, date,parkingTicket.Id)
	parkingTicket.Ticket_number = "*" + parkingTicket.Ticket_number + "*"
	parkingResponse.Data = parkingTicket

	c.JSON(http.StatusOK, parkingResponse)
	return
}

func isTicketNumberExists(c *gin.Context, ticketNumber string) bool {
	_, err := models.ParkingTicketExistsByNumber(ticketNumber)

	if err == sql.ErrNoRows {
		return false
	}

	log.Println(err)
	return true
}

func pictureFullPath(pic models.Picture) string {
	return pic.Filepath+string(os.PathSeparator)+pic.Filename+"."+pic.Format
}

func saveCameraPicture(picture []byte, date time.Time, ticketId int64) {
	//saving picture
	dateTimeName := date.Format("2006-01-02 15:04:05 PM")
	pictureName := strings.Replace(dateTimeName, ":", "", 10) + "-T" + strconv.FormatInt(ticketId, 10)//+ "-" + string(ticketId)
	//save picture path to database
	pic := models.Picture{}
	pic.Filepath = "campicture"
	pic.Filename = pictureName
	pic.Format = "jpg"
	pic.Expired_date = date.Add(time.Hour * 240).String()
	pic.Created_by = -1
	pic.Created_date = date.String()
	result, err := models.PictureSave(pic)
	fileFullPath := pictureFullPath(pic)

	if err != nil {
		log.Println(err)
	}
	//save picture bytes to a jpg image
	err = ioutil.WriteFile(fileFullPath, picture, 0666)

	if err != nil {
		log.Println(err)
	}

	pictureId, _ := result.LastInsertId()

	err = models.ParkingTicketUpdatePictureIn(pictureId, ticketId)

	if err != nil {
		log.Println(err)
	}
}

//parking ticket info API can only be retrieved from cashier computer
func ParkingGetTicketInfo(c *gin.Context) {
	//var
	var isMember bool
	currentTime := time.Now()
	parkingResponse := ParkingResponse{}
	response := new(response.SimpleResponse)
	var err error
	parkingResponse.Response = response

	ticketNumber := c.Query("ticket_number")
	vehicleNumber := c.Query("vehicle_number")

	parkingTicket, err := models.ParkingGetTicketByTicketNumber(ticketNumber)

	if err != nil {
		if err == sql.ErrNoRows {
			response.Status = "Failed"; response.Message = "Ticket not found"
			c.JSON(http.StatusOK, parkingResponse)
			return
		}

		log.Println(err)
		response.Status = "Failed"; response.Message = "System fatal error"
		c.JSON(http.StatusOK, parkingResponse)
		return
	}

	pictureIn, err := models.PictureGetById(parkingTicket.Picture_in_id)

	if err != nil {
		log.Println(err)
	}

	pictureInPath := pictureFullPath(pictureIn)
	parkingResponse.Picture_path_in = pictureInPath

	//get member
	isMember = true
	_, err = models.MemberGetByPoliceNumber(vehicleNumber)

	if err != nil {
		isMember = false

		if err != sql.ErrNoRows {
			log.Println(err)
		}
	}

	//get parking price
	parkingPrice, err := models.ParkingPriceGetByVehicleId(parkingTicket.Vehicle_id)
	if err != nil {
		log.Println(err)
		response.Status = "Failed"; response.Message = "Parking price not found"
		c.JSON(http.StatusOK, parkingResponse)
		return
	}

	deltaTime := currentTime.Sub(now.MustParse(parkingTicket.Created_date))
	//parking cost
	totalCost := parkingPrice.First_hour_price
	totalCost += (int(math.Ceil(deltaTime.Hours())) - 1) * parkingPrice.Next_hour_price
	if parkingPrice.Maximum_price != 0 {
		if totalCost > parkingPrice.Maximum_price {
			totalCost = parkingPrice.Maximum_price
		}
	}
	parkingCost := totalCost
	//set to zero if member
	if isMember {
		parkingCost = 0
	}

	//------------
	parkingTicket.Parking_cost = int(parkingCost)
	parkingTicket.Out_date = currentTime.Format("2006-01-02 15:04:05")
	deltaHours := int(deltaTime.Hours())
	deltaMin := int(deltaTime.Minutes()) - (deltaHours * 60)
	deltaSecs := int(deltaTime.Seconds()) - (int(deltaTime.Minutes()) * 60)
	parkingResponse.DeltaTime = strconv.Itoa(deltaHours) + "h " + strconv.Itoa(deltaMin) + "m " + strconv.Itoa(deltaSecs) + "s"
	//log.Println("%+v", parkingResponse)
	response.Status = "Success"; response.Message = "Thank you!"

	parkingResponse.Data = parkingTicket
	c.JSON(http.StatusOK, parkingResponse)
}

func ParkingCheckOut(c *gin.Context) {
	ticketId, err := strconv.ParseInt(c.PostForm("ticket_id"), 10, 64)
	ticketNumber := c.PostForm("ticket_number")
	vehicleNumber := c.PostForm("vehicle_number")
	dateOut := c.PostForm("ticket_date_out")
	parkingCost, err := strconv.Atoi(c.PostForm(("parking_cost")))
	pictureOutId, err := strconv.ParseInt(c.PostForm("picture_out_id"), 10, 64)
    deltaTime := c.PostForm("parking_duration") 

	response := new(response.SimpleResponse)
	parkingResponse := ParkingResponse{}
	parkingResponse.Response = response

	//session
	session := session.Instance(c)
	executor := session.Get("id").(int64)

	_, err = models.PictureGetById(pictureOutId)

	if err != nil {
		log.Println("Invalid picture, err: ", err)
	}

	cashier, err := hostIsCashier(c, c.ClientIP())

	/*if err != errors.New("") {
		if err == sql.ErrNoRows {
			response.Status = "Failed"; response.Message = "Invalid host"
			c.JSON(http.StatusOK, parkingResponse)
			return
		}
		log.Println(err)
		response.Status = "Failed"; response.Message = "System fatal error"
		c.JSON(http.StatusOK, parkingResponse)
		return
	}*/

	if err != nil {
		response.Status = "Failed"; response.Message = "Not a cashier host"
		c.JSON(http.StatusOK, parkingResponse)
		return
	}

	parkingTicket, err := models.ParkingGetTicketByNumberAndId(ticketId, ticketNumber)

	if err == sql.ErrNoRows {
		response.Status = "Failed"; response.Message = "Ticket number not found"
		c.JSON(http.StatusOK, parkingResponse)
		return
	}

	if parkingTicket.Out_date != "" {
		response.Status = "Failed"; response.Message = "Invalid ticket"
		c.JSON(http.StatusOK, parkingResponse)
		return
	}

	//picture, err := models.PictureGetById(pictureOutId)

	if err != nil {
		log.Println(err)
	}
    
    raspberryPi := &api.RaspberryPi{}
    raspberryPi.Protocol="http"
    raspberryPi.Host = cashier.Host
    raspberryPi.Port = "8888"
    raspberryPi.Token = cashier.Token
    raspberryPi.Param = ""
    raspberryPi.RaspberryPrintTicketOut()

	parkingTicket.Parking_cost = parkingCost
	parkingTicket.Vehicle_number = vehicleNumber
	parkingTicket.Verified_by = executor
	parkingTicket.Out_date = dateOut
	parkingTicket.Last_update_date = dateOut
	parkingTicket.Updated_by = executor
	parkingTicket.Picture_out_id = pictureOutId

	err = models.ParkingUpdateTicket(parkingTicket)

	if err != nil {
		response.Status = "Failed"; response.Message = "Cannot update parking ticket"
		c.JSON(http.StatusOK, parkingResponse)
		return
	}

	//send httppost to raspberry to print ticket
    parkingResponse.DeltaTime = deltaTime 
	response.Status = "Failed"; response.Message = "Thank you"
	parkingResponse.Data = parkingTicket
	c.JSON(http.StatusOK, parkingResponse)
}

//update picture after out
func updatePictureOut(pictureId int64, ticketId int64) {

}

func hostIsCashier(c *gin.Context, ip string) (models.Device, error) {
    var err error
	slicedIp := strings.Split(ip, ":")
	deviceIp := slicedIp[0]

	device, err := models.DeviceGetByHost(deviceIp)

	if err != nil {
		return device, err
	}

	if !strings.EqualFold("Cashier", device.Device_type) {
		return device, errors.New("Not a cashier")
	}

	return device, err
}
