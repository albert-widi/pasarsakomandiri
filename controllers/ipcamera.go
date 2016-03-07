package controllers

import "github.com/gin-gonic/gin"
import (
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/pasarsakomandiri/api"
    	"github.com/pasarsakomandiri/models"
	"net/http"
	"github.com/pasarsakomandiri/shared/response"
	"github.com/pasarsakomandiri/shared/static"
	"strconv"
)

func getIpCamPicture(ipcamhost string) {
	date := time.Now()
	dateTimeName := date.Format("2006-01-02 03:04:05 PM")

	ipCamera := api.IpCamera{}
	ipCamera.Protocol = "http"
	ipCamera.Param = "Streaming/channels/1/picture"
	ipCamera.Host = ipcamhost
	ipCamera.Username = "admin"
	ipCamera.Password = "12345"
	ipCamera.GetPicture()

	
}

//taking picture from camera
func IpCamTakePictureFromDevice(c *gin.Context) {
	date := time.Now()
	dateTimeName := date.Format("2006-01-02 03:04:05 PM")
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

	ipCamera := api.IpCamera{}
	ipCamera.Protocol = "http"
	ipCamera.Param = "Streaming/channels/1/picture"
	ipCamera.Host = deviceGroup.Camera_ip
	ipCamera.Username = "admin"
	ipCamera.Password = "12345"
    	ipCamera.Picture = make(chan []byte)
	//get picture from ipcamera
	go ipCamera.GetPicture()

	pic := models.Picture{}
	pic.Filepath = "campicture"
	pic.Filename = strings.Replace(dateTimeName, ":", "", 10)
	pic.Format = "jpg"
	pic.Expired_date = date.Add(time.Hour * 24).String()
	pic.Created_by = -1
	pic.Created_date = date.String()

	//save filepath to db
	result, dbErr := models.PictureSave(pic)

	if dbErr != nil {
		log.Println(dbErr)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "System Error"))
		return
	}

	pic.Filename = pic.Filename + " - " + strconv.FormatInt(result.LastInsertId(), 10)

	//save file to filesystem
	fsErr := static.SaveFileToStaticFS(<-ipCamera.Picture, pic.PictureGetFullPath())

	if fsErr != nil {
		log.Println(fsErr)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "System Error"))
		return
	}

	c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "System Error"))
	return
}
