package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/pasarsakomandiri/shared/session"
	"strconv"
	"github.com/pasarsakomandiri/models"
	"database/sql"
	"log"
	"github.com/pasarsakomandiri/shared/response"
	"time"
)

func PriceConfigPage (c *gin.Context){
	session := session.Instance(c)
	c.HTML(http.StatusFound, "parking_price.tmpl", gin.H{"title" : "Parking Price", "token":session.Get("token")})
}

func PriceGetAll(c *gin.Context) {
	price, err := models.ParkingPriceGetAllAPI()

	if err != nil{
		log.Println("ParkingPriceGetAllApi Error", err)
		return
	}
	c.JSON(http.StatusOK, price)
}

func PriceRegister(c *gin.Context) {
	vehicleId, err := strconv.Atoi(c.PostForm("vehicle_id"))
	jamPertama, err := strconv.Atoi(c.PostForm("jam_pertama"))
	jamBerikutnya, err := strconv.Atoi(c.PostForm("jam_berikutnya"))
	biayaMax, err := strconv.Atoi(c.PostForm("biaya_max"))

	vehicle, err := models.ParkingVehicleGetByID(vehicleId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Vehicle not found"))
			return
		}
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "System fatal error"))
		return
	}

	_, err = models.ParkingPriceGetByVehicleId(vehicle.Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
			c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "System fatal error"))
			return
		}
	}

	if err == nil {
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Price for this vehicle already exists"))
		return
	}


	price := models.ParkingPrice{}
	price.Id = 0
	price.Vehicle_id = vehicle.Id
	price.Vehicle_type = vehicle.Vehicle_type
	price.First_hour_price = jamPertama
	price.Next_hour_price = jamBerikutnya
	price.Maximum_price = biayaMax
	price.Created_by = session.Instance(c).Get("id").(int64)
	price.Created_date = time.Now().String()

	if price.First_hour_price == 0 || price.Next_hour_price == 0 {
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Required items cannot be empty"))
		return
	}

	err = models.ParkingPriceCreateNew(price)

	if err != nil {
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Failed create parking price"))
		return
	}

	c.JSON(http.StatusOK, response.NewSimpleResponse("Success", "Parking price created"))
}