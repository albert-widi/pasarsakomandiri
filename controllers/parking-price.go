package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pasarsakomandiri/models"
	"github.com/pasarsakomandiri/shared/response"
	"github.com/pasarsakomandiri/shared/session"
	"github.com/pasarsakomandiri/shared/view"
)

func PriceConfigPage(c *gin.Context) {
	v := view.New(c)
	v.Name = "parking_price"
	v.Render()
	//session := session.Instance(c)
	//c.HTML(http.StatusFound, "parking_price.tmpl", gin.H{"title": "Parking Price", "token": session.Get("token")})
}

func PriceUpdateConfigPage(c *gin.Context) {
	v := view.New(c)
	v.Name = "parking_price"
	v.Render()
	//session := session.Instance(c)
	//c.HTML(http.StatusFound, "update_parking_price.tmpl", gin.H{"title": "Update Parking Price", "token": session.Get("token")})
}

func PriceGetInfoAPI(c *gin.Context) {
	priceId, err := strconv.ParseInt(c.Query("id"), 10, 64)

	fmt.Println(priceId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Invalid Price ID"))
	}
	price, err := models.ParkingPriceGetById(priceId)

	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Parking Price Not Found"))
			return
		}
	}
	c.JSON(http.StatusOK, price)
}

func PriceGetAll(c *gin.Context) {
	price, err := models.ParkingPriceGetAllAPI()

	if err != nil {
		log.Println("ParkingPriceGetAllApi Error", err)
		return
	}
	c.JSON(http.StatusOK, price)
}

func PriceRegister(c *gin.Context) {
	vehicleId, err := strconv.Atoi(c.PostForm("vehicle_id"))
	jamPertama, err := strconv.Atoi(c.PostForm("first_hour_price"))
	jamBerikutnya, err := strconv.Atoi(c.PostForm("next_hour_price"))
	promoJamPertama, err := strconv.Atoi(c.PostForm("promo_jam_pertama"))
	promoJamBerikutnya, err := strconv.Atoi(c.PostForm("promo_jam_berikutnya"))
	biayaMax, err := strconv.Atoi(c.PostForm("maximum_price"))

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
	price.Promo_jam_pertama = promoJamPertama
	price.Promo_jam_berikutnya = promoJamBerikutnya
	price.Next_hour_price = jamBerikutnya
	price.Maximum_price = biayaMax
	price.Created_by = session.Instance(c).Get("id").(int64)
	price.Created_date = time.Now().String()

	if price.First_hour_price < 0 || price.Next_hour_price < 0 {
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

func PriceUpdateAPI(c *gin.Context) {
	priceId, err := strconv.ParseInt(c.PostForm("id"), 10, 64)
	log.Println(priceId)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Invalid Price ID"))
	}

	_, err = models.ParkingPriceGetById(priceId)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Failed Get Price ID"))
	}

	vehicleId, err := strconv.Atoi(c.PostForm("vehicle_id"))
	first_hour_price, err := strconv.Atoi(c.PostForm("first_hour_price"))
	promoJamPertama, err := strconv.Atoi(c.PostForm("promo_jam_pertama"))
	next_hour_price, err := strconv.Atoi(c.PostForm("next_hour_price"))
	promoJamBerikutnya, err := strconv.Atoi(c.PostForm("promo_jam_berikutnya"))
	maximum_price, err := strconv.Atoi(c.PostForm("maximum_price"))

	price := models.ParkingPrice{}
	price.Id = priceId
	price.Vehicle_id = vehicleId
	price.First_hour_price = first_hour_price
	price.Promo_jam_pertama = promoJamPertama
	price.Next_hour_price = next_hour_price
	price.Promo_jam_berikutnya = promoJamBerikutnya
	price.Maximum_price = maximum_price

	log.Println(price.First_hour_price)
	fmt.Println(next_hour_price)
	fmt.Println(maximum_price)

	if price.First_hour_price == 0 || price.Next_hour_price == 0 {
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Required items cannot be empty"))
		return
	}

	err = models.ParkingPriceUpdate(price)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Failed Updating Parking Price"))
		return
	}

	c.JSON(http.StatusOK, response.NewSimpleResponse("Success", "Parking price Updated"))
}
