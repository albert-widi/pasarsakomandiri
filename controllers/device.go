package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pasarsakomandiri/models"
	"net/http"
	"github.com/pasarsakomandiri/shared/response"
	"database/sql"
	"strconv"
	"github.com/pasarsakomandiri/shared/session"
	"log"
	_"fmt"
	"time"
	"github.com/pasarsakomandiri/api"
)

func DeviceListPages(c *gin.Context)  {
	session := session.Instance(c)
	c.HTML(http.StatusFound, "device_list.tmpl", gin.H{"title":"Cahiser", "token":session.Get("token")})
}

func DeviceListEditPages(c *gin.Context)  {
	session := session.Instance(c)
	c.HTML(http.StatusFound, "edit_devicelist.tmpl", gin.H{"tittle":"Edit Device List", "token":session.Get("token")})
}
func GetDeviceListAPI(c *gin.Context) {
	devices, err := models.DeviceGetAll()

	if err != nil {
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", err.Error()))
		return
	}

	c.JSON(http.StatusOK, devices)
	return
}

func GetDeviceType(c *gin.Context)  {
	devices, err := models.DeviceGetAllDeviceType()

	if err != nil{
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", err.Error()))
		return
	}

	c.JSON(http.StatusOK, devices)
	return
}

//get device by device type
//for example you can get device data by giving device_type parameter in your get to get all devices by "Raspberry" type
//get: device_name='Raspberry', return []Devices in JSON
func DeviceGetFromTypeAPI(c *gin.Context) {
	deviceType := c.Query("device_type")

	devices, err := models.DevicesGetByType(deviceType)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "No data"))
			return
		}

		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "System fatal error"))
		return
	}

	c.JSON(http.StatusOK, devices)
}

//device information
func DeviceGetDeviceInfoAPI(c *gin.Context) {

	deviceID, err := strconv.ParseInt(c.Query("device_id"), 10, 64)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Invalid form of ID"))
		return
	}

	device, err := models.DeviceGetByID(deviceID)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Device not found"))
			return
		}

		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", err.Error()))
		return
	}

	c.JSON(http.StatusOK, device)
}

//register devices and groups
func DeviceRegisterAPI(c *gin.Context) {
	device := models.Device{}
	device.Id = 0
	device.Device_name = c.PostForm("name")
	device.Device_type = c.PostForm("type")
	device.Host = c.PostForm("host")
	device.Token = c.PostForm("token")
	device.Description = c.PostForm("description")

	if device.Device_name == "" || device.Host == "" || device.Token == "" {
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Required items cannot be empty"))
		return
	}

	dtype, err := models.DeviceGetDeviceTypeByName(device.Device_type)
	if err == sql.ErrNoRows {
		//log.Println("Device type: ", device.Device_type)
		//log.Println("error bos", err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Invalid device"))
		return
	}

	//set device type id
	device.Device_type_id = dtype.Id
	device.Created_by = session.Instance(c).Get("id").(int64)
	device.Created_date = time.Now().String()

	_, err = models.DeviceGetByHostType(device.Host, device.Device_type)

	if err != nil {
		if err != sql.ErrNoRows {
			c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", err.Error()))
			return
		}
	}

	err = models.DeviceCreateNew(device)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Failed to create device"))
		return
	}

	c.JSON(http.StatusOK, response.NewSimpleResponse("Success", "Device created"))
}

func DeviceContactCheckOut(c *gin.Context) {
	api.ContactClientCheckOut()
}


//DELETE
func DeleteDeviceList(c *gin.Context)  {

	devid, err := strconv.ParseInt(c.PostForm("device_id"), 10, 64)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Invalid get of ID"))
	}

	_, err = models.DeviceGetByID(devid)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Invalid get of ID"))
	}

	devlistid := models.Device{}
	devlistid.Id = devid

	err = models.DeleteDeviceById(devlistid)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Invalid get of ID"))
	}

	c.JSON(http.StatusOK, response.NewSimpleResponse("Success", "Delete Success"))
}

//EDIT
func EditDeviceList(c *gin.Context)  {

	iddev, err := strconv.ParseInt(c.PostForm("device_id"), 10, 64)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Invalid get of ID"))
	}

	_, err = models.DeviceGetByID(iddev)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Invalid get of ID"))
	}

	deviceType := c.PostForm("device_type")
	deviceName := c.PostForm("device_name")
	deviceHost := c.PostForm("host")
	deviceToken := c.PostForm("token")
	deviceDescription := c.PostForm("description")

	devlist := models.Device{}
	devlist.Id = iddev
	devlist.Device_type = deviceType
	devlist.Device_name = deviceName
	devlist.Host = deviceHost
	devlist.Token = deviceToken
	devlist.Description = deviceDescription

	err = models.DeviceEdit(devlist)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Invalid get of ID"))
	}

	c.JSON(http.StatusOK, response.NewSimpleResponse("Success", "Edit Success"))

}