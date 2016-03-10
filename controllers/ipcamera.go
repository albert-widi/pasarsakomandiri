package controllers

import "github.com/gin-gonic/gin"
import (
	"log"
	"strings"
	"time"

	"github.com/pasarsakomandiri/api"
    	"github.com/pasarsakomandiri/models"
	"net/http"
	"github.com/pasarsakomandiri/shared/response"
	"github.com/pasarsakomandiri/shared/static"
	"strconv"
	"github.com/fatih/structs"
	"database/sql"
)

var timeSaveStandard time.Duration = time.Duration(24)
var nilInterface map[string]interface{}

func getIpCamPicture(ipcamhost string) api.IpCamera {
	ipCamera := api.IpCamera{}
	ipCamera.Protocol = "http"
	ipCamera.Param = "Streaming/channels/1/picture"
	ipCamera.Host = ipcamhost
	ipCamera.Username = "admin"
	ipCamera.Password = "sakomandiri1"
	ipCamera.GetPicture()

	return ipCamera
}

func saveIpCamPicture(date time.Time, ipCamera api.IpCamera) (models.Picture, error) {
	dateTimeName := date.Format("2006-01-02 03:04:05 PM")
	var err error
	pic := models.Picture{}
	pic.Filepath = "campicture"
	pic.Filename = strings.Replace(dateTimeName, ":", "", 10)
	pic.Format = "jpg"
	pic.Expired_date = date.Add(time.Hour * timeSaveStandard).String()
	pic.Created_by = -1
	pic.Created_date = date.String()

	//save filepath to db
	result, dbErr := models.PictureSave(pic)
	err = dbErr

	if err != nil {
		return pic, err
	}

	pictureId, _ := result.LastInsertId()

	//append picture id to filaname
	pic.Filename = pic.Filename + " P" + strconv.FormatInt(pictureId, 10)
	//update picture name
	err = models.PictureUpdateName(pic.Filename, pictureId)
	if err != nil {
		log.Println("Update picture name error, ", err)
	}

	//save file to filesystem
	pic.PictureFullPath = pic.GetFullPath()
	fsErr := static.SaveFileToStaticFS(ipCamera.Picture, pic.PictureFullPath)
	err = fsErr
	if err != nil {
		return pic, fsErr
	}

	return pic, err
}

//taking picture by ip
func IpCamTakePictureByIP(c *gin.Context) {
	date := time.Now()
	camIp := c.Query("camip")
	ipCamera := getIpCamPicture(camIp)
	picture, err := saveIpCamPicture(date, ipCamera)

	pictureMap := structs.Map(picture)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Error: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewDataResponse("Success", "Here is your picture", pictureMap))
	return
}

//taking picture from device
func IpCamTakePictureFromDevice(c *gin.Context) {
	date := time.Now()
	slicedIp := strings.Split(c.ClientIP(), ":")
	deviceIp := slicedIp[0]

	device, err := models.DeviceGetByHost(deviceIp)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Device not found"))
		return
	}

	//default condition is raspberry
	condition := "raspberry_ip"
	deviceGroup, err := models.DeviceGroupGetByHost(condition, device.Host)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Device not in device group"))
			return
		}

		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Please try again later"))
		return
	}

	//get picture from ipcamera
	ipCamera := getIpCamPicture(deviceGroup.Camera_ip)
	picture, err := saveIpCamPicture(date, ipCamera)

	pictureMap := structs.Map(picture)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "System Error: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewDataResponse("Success", "Here is your picture", pictureMap))
	return
}
