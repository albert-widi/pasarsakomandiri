package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/pasarsakomandiri/models"
	"database/sql"
	"net/http"
	"github.com/pasarsakomandiri/shared/response"
	"log"
	"github.com/pasarsakomandiri/shared/session"
)

func DeviceGroupPage(c *gin.Context) {
	session := session.Instance(c)
	c.HTML(http.StatusFound, "d_group.tmpl", gin.H{"title":"Device Group", "token":session.Get("token")})
}

func DeviceGroupGetAllAPI (c *gin.Context){
	groupDevice, err := models.DeviceGroupGetAllAPI()

	if err != nil {
		log.Println("DevicegroupGetAllAPI error : ", err)
		return
	}
	c.JSON(http.StatusOK, groupDevice)
}

func DeviceGroupRegisterAPI(c *gin.Context) {
	raspberryId, err := strconv.ParseInt(c.PostForm("raspberry_id"), 10, 64)
	cameraId, err := strconv.ParseInt(c.PostForm("camera_id"), 10, 64)
	vehicleId, err := strconv.Atoi(c.PostForm("vehicle_id"))
	gateName := c.PostForm("gate_name")

	raspberry, err := models.DeviceGetByID(raspberryId)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Raspberry not found"))
			return
		}

		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", string(err.Error())))
		return
	}

	if raspberry.Device_type != "Raspberry" {
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Device type of Raspberry is not Raspberry"))
		return
	}

	camera, err := models.DeviceGetByID(cameraId)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Camera not found"))
			return
		}

		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", string(err.Error())))
		return
	}

	if camera.Device_type != "Camera" {
		//log.Println(camera.Device_type)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Device type of IPCamera is not IPCamera"))
		return
	}

	vehicle, err := models.ParkingVehicleGetByID(vehicleId)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Vehicle not found"))
			return
		}

		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "System fatal error"))
		return
	}

	_, err = models.DeviceGroupGetByHost("raspberry_ip", raspberry.Host)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
			c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "System fatal error"))
			return
		}
	}

	if err == nil {
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Device group already exists"))
		return
	}

	_, err = models.DeviceGroupGetByHost("camera_ip", camera.Host)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
			c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "System fatal error"))
			return
		}
	}

	if err == nil {
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Device group already exists"))
		return
	}

	/*log.Println("Camera IP: ", group.Camera_ip)
	log.Println("Camera host: ", camera.Host)

	if group.Camera_ip == camera.Host {
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Camera already grouped in another group"))
		return
	}*/

	deviceGroup := models.DeviceGroup{}
	deviceGroup.Id = 0
	deviceGroup.Raspberry_id = raspberryId
	deviceGroup.Raspberry_ip = raspberry.Host
	deviceGroup.Camera_id = cameraId
	deviceGroup.Camera_ip = camera.Host
	deviceGroup.Vehicle_id = vehicle.Id
	deviceGroup.Vehicle_type = vehicle.Vehicle_type
	deviceGroup.Gate_name = gateName

	err = models.DeviceGroupCreateNew(session.Instance(c).Get("id").(int64), deviceGroup)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Failed to create device group"))
		return
	}

	c.JSON(http.StatusOK, response.NewSimpleResponse("Success", "Device group created"))
}

func DeviceGroupDeleteAPI (c *gin.Context){
	deviceId, err := strconv.ParseInt(c.PostForm("id"),10, 64)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Invalid ID"))
		return
	}

	err = models.DeviceGroupDeleteById(deviceId)

	if err != nil {
		log.Println("Device Group DELETE Error : ", err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Failed to Delete Device group"))
		return
	}
	c.JSON(http.StatusOK, response.NewSimpleResponse("Success", "Device Group Deleted"))
}
