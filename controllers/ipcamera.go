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
)

var timeSaveStandard time.Duration = time.Duration(24)

func getIpCamPicture(ipcamhost string) api.IpCamera {
	ipCamera := api.IpCamera{}
	ipCamera.Protocol = "http"
	ipCamera.Param = "Streaming/channels/1/picture"
	ipCamera.Host = ipcamhost
	ipCamera.Username = "admin"
	ipCamera.Password = "12345"
	ipCamera.GetPicture()

	return ipCamera
}

func saveIpCamPicture(date time.Time, ipCamera api.IpCamera) (string, error) {
	dateTimeName := date.Format("2006-01-02 03:04:05 PM")
	var err error
	pic := models.Picture{}
	pic.Filepath = "campicture"
	pic.Filename = strings.Replace(dateTimeName, ":", "", 10)
	pic.Format = "jpg"
	pic.Expired_date = date.Add(time.Hour * timeSaveStandard).String()
	pic.Created_by = -1
	pic.Created_date = date.String()
	pic.PictureFullPath = pic.GetFullPath()

	//save file to filesystem
	fsErr := static.SaveFileToStaticFS(ipCamera.Picture, pic.PictureFullPath)
	err = fsErr
	if err != nil {
		return "", fsErr
	}

	//save filepath to db
	_, dbErr := models.PictureSave(pic)
	err = dbErr

	if err != nil {
		return "", err
	}

	return pic.PictureFullPath, err
}

//taking picture by ip
func IpCamTakePictureByIP(c *gin.Context) {
	date := time.Now()
	camIp := c.Query("camip")
	ipCamera := getIpCamPicture(camIp)
	filePath, err := saveIpCamPicture(date, ipCamera)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "System Error: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewSimpleResponse("Success", filePath))
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
	//get picture from ipcamera
	ipCamera := getIpCamPicture(deviceGroup.Camera_ip)
	filePath, err := saveIpCamPicture(date, ipCamera)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "System Error: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewSimpleResponse("Success", filePath))
	return
}
