package controllers
import (
	"github.com/pasarsakomandiri/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/pasarsakomandiri/shared/response"
	"log"
	"github.com/gin-gonic/contrib/sessions"
)

func VehicleRegister(c *gin.Context) {
	vehicleType := c.PostForm("vehicle_type")

	if !models.ParkingCreateNewVehicle(sessions.Default(c).Get("id").(int64), vehicleType) {
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Vehicle creation failed"))
		return
	}

	c.JSON(http.StatusOK, response.NewSimpleResponse("Success", "Vehicle Created"))
}

func VehicleGetAll(c *gin.Context) {
	vehicle, err := models.ParkingVehicleGetAll()

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "No data available"))
			return
		}

		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "System fatal error"))
		return
	}

	c.JSON(http.StatusOK, vehicle)
}