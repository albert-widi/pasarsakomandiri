package controllers

import "github.com/gin-gonic/gin"
import (
	"github.com/pasarsakomandiri/api"
	"io/ioutil"
	"log"
	"time"
	"strings"
)

func IpCamTakePictureFromDevice(c *gin.Context) {
	/*slicedIp := strings.Split(c.ClientIP(), ":")
	deviceIp := slicedIp[0]

	deviceGroup := models.DeviceGroupGetByHost()*/

	ipCamera := api.IpCamera{}
	ipCamera.Protocol = "http"
	ipCamera.Param = "Streaming/channels/1/picture"
	ipCamera.Host = "192.168.0.187"
	ipCamera.Username = "admin"
	ipCamera.Password = "12345"

	//take picture
	picture, err := api.IpCamGetPicture(ipCamera)
	//var b bytes.Buffer
	//archieveWriter := gzip.NewWriter()

	if err != nil {
		log.Println(err)
	}

	//fmt.Printf("%v", picture)

	dateTimeName := time.Now().Format("2006-01-02 03:04:05 PM")
	pictureName := strings.Replace(dateTimeName, ":", "", 10)
	err = ioutil.WriteFile("campicture/" + pictureName + ".jpg", picture, 0666)

	if err != nil {
		log.Println(err)
	}
}
