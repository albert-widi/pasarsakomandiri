package testing

import (
	"github.com/gin-gonic/gin"
	"time"
	"strings"
	"github.com/pasarsakomandiri/models"
	"github.com/pasarsakomandiri/shared/response"
	"database/sql"
	"net/http"
	"log"
	"github.com/pasarsakomandiri/api"
	"io/ioutil"
	"strconv"
	"github.com/pasarsakomandiri/shared/tools"
	"github.com/pasarsakomandiri/shared/database"
)

type ParkingResponse struct {
	Response *response.SimpleResponse
	Data models.ParkingTicket
	DeltaTime string
}

func TestingParkingCheckin(c *gin.Context) {
	date := time.Now()
	slicedIp := strings.Split(c.ClientIP(), ":")
	deviceIp := slicedIp[0]
	deviceToken := c.PostForm("token")


	device, err := models.DeviceGetByHost(deviceIp)
	parkingResponse := ParkingResponse{}
	response := new(response.SimpleResponse)
	parkingResponse.Response = response

	if err != nil {
		if err ==  sql.ErrNoRows {
			response.Status = "Failed"; response.Message = "Device not found"
			c.JSON(http.StatusOK, parkingResponse)
			return
		}

		log.Println("Db error: ", err)
		return
	}

	if !strings.EqualFold("Raspberry", device.Device_type){
		response.Status = "Failed"; response.Message = "Device is not valid"
		c.JSON(http.StatusOK, parkingResponse)
		return
	}

	if !strings.EqualFold(deviceToken, device.Token) {
		response.Status = "Failed"; response.Message = "Token not match"
		c.JSON(http.StatusOK, parkingResponse)
		return
	}

	deviceGroup, err := models.DeviceGroupGetByHost("raspberry_ip", deviceIp)

	if err == sql.ErrNoRows {
		response.Status = "Failed"; response.Message = "Device group not found"
		c.JSON(http.StatusOK, parkingResponse)
		return
	}

	//set ipcamera param
	ipCamera := api.IpCamera{}
	ipCamera.Protocol = "http"
	ipCamera.Param = "Streaming/channels/1/picture"
	ipCamera.Host = deviceGroup.Camera_ip
	ipCamera.Username = "admin"
	ipCamera.Password = "12345"

	//take picture
	picture, err := api.IpCamGetPicture(ipCamera)

	if err != nil {
		log.Println(err)
	}

	//saving picture
	dateTimeName := date.Format("2006-01-02 03:04:05 PM")
	pictureName := strings.Replace(dateTimeName, ":", "", 10)
	//save picture path to database
	pic := models.Picture{}
	pic.Filepath = "campicture/"
	pic.Filename = pictureName
	pic.Expired_date = date.Add(time.Hour * 240).String()
	pic.Created_by = -1
	pic.Created_date = date.String()
	result, err := models.PictureSave(pic)

	if err != nil {
		log.Println(err)
	}
	//save picture bytes to a jpg image
	err = ioutil.WriteFile("campicture/"+pictureName+".jpg", picture, 0666)

	if err != nil {
		log.Println(err)
	}

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
	ticketExists := true
	for ticketExists {
		parkingTicket.Ticket_number = tools.RandomString(8)
		parkingTicket.Ticket_number =  parkingTicket.Ticket_number
		ticketExists = isTicketNumberExists(c, parkingTicket.Ticket_number)
	}

	//save ticket and picture to database
	//ticket saved to db = print, esle don't print
	result, err = models.ParkingCreateNewTicket(database.DbInstance(c), parkingTicket)

	if err != nil {
		response.Status = "Failed"; response.Message = "Cannot create ticket to db"
		c.JSON(http.StatusOK, parkingResponse)
		return
	}

	parkingTicket.Id, _ = result.LastInsertId()
	parkingTicket.Ticket_number = "*" + parkingTicket.Ticket_number + "*"
	parkingResponse.Data = parkingTicket
	c.JSON(http.StatusOK, parkingResponse)
	return
}

func isTicketNumberExists(c *gin.Context, ticketNumber string) bool {
	_, err := models.ParkingTicketExistsByNumber(database.DbInstance(c), ticketNumber)

	if err == sql.ErrNoRows {
		return false
	}

	log.Println(err)
	return true
}