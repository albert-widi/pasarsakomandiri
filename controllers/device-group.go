package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pasarsakomandiri/models"
	"github.com/pasarsakomandiri/shared/response"
	"github.com/pasarsakomandiri/shared/session"
	"github.com/pasarsakomandiri/shared/view"
)

func DeviceGroupPage(c *gin.Context) {
	v := view.New(c)
	v.Name = "d_group"
	v.Render()
	//session := session.Instance(c)
	//c.HTML(http.StatusFound, "d_group.tmpl", gin.H{"title": "Device Group", "token": session.Get("token")})
}

func DeviceGroupGetAllAPI(c *gin.Context) {
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
	groupName := c.PostForm("gate_name")
	groupType := c.PostForm("group_type")

	raspberry, err := models.DeviceGetByID(raspberryId)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Raspberry not found"))
			return
		}

		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", string(err.Error())))
		return
	}

	//log.Println("Device type: ", raspberry.Device_type)
	log.Println("Compare, ", strings.Compare("Raspberry", raspberry.Device_type))
	log.Println("Compare, ", strings.Compare("Cashier", raspberry.Device_type))

	if raspberry.Device_type != "Raspberry" && raspberry.Device_type != "Cashier" {
		//if strings.Compare("Cashier", raspberry.Device_type) != 0 && strings.Compare("Raspberry", raspberry.Device_type) != 0 {
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Device type is not Raspberry or Cashier"))
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

	deviceGroup := models.DeviceGroup{}
	deviceGroup.Raspberry_id = raspberryId
	deviceGroup.Raspberry_ip = raspberry.Host
	deviceGroup.Camera_id = cameraId
	deviceGroup.Camera_ip = camera.Host
	deviceGroup.Group_name = groupName
	deviceGroup.Group_type = groupType

	if raspberry.Device_type == "Raspberry" {
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

		deviceGroup.Vehicle_id = vehicle.Id
		deviceGroup.Vehicle_type = vehicle.Vehicle_type
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

	err = models.DeviceGroupCreateNew(session.Instance(c).Get("id").(int64), deviceGroup)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Failed to create device group"))
		return
	}

	c.JSON(http.StatusOK, response.NewSimpleResponse("Success", "Device group created"))
}

func DeviceGroupDeleteAPI(c *gin.Context) {
	deviceId, err := strconv.ParseInt(c.PostForm("id"), 10, 64)

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
