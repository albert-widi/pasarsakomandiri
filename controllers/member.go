package controllers
import (
	"github.com/pasarsakomandiri/shared/session"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/pasarsakomandiri/models"
	"github.com/pasarsakomandiri/shared/response"
	"log"
	"strconv"
	"database/sql"
	"time"
	"strings"
)

func MemberPages(c *gin.Context)  {
	session := session.Instance(c)
	c.HTML(http.StatusFound, "member.tmpl", gin.H{"tittle":"Member Pages", "token":session.Get("token")})
}

func MemberCreateNew(c *gin.Context)  {

	vehicleId, err := strconv.Atoi(c.PostForm("vehicle_id"))

	member := models.Member{}
	//member.Id =
	member.Vehicle_type = c.PostForm("vehicle_type")
	member.Police_number = strings.ToUpper(c.PostForm("police_number"))
	member.Description = c.PostForm("description")
	member.Created_by = -1
	member.Created_date = time.Now().String()

	vehicle, err := models.ParkingVehicleGetByID(vehicleId)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Vehicle not found"))
			return
		}

		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", string(err.Error())))
		return
	}

	member.Vehicle_id = vehicle.Id
	member.Vehicle_type = vehicle.Vehicle_type

	err = models.MemberCreateNew(member)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Failed to create member"))
		return
	}

	c.JSON(http.StatusOK, response.NewSimpleResponse("Success", "Member created"))


}


func MemberGetAll(c *gin.Context)  {
	member, err := models.MemberGetAll()

	if err != nil {
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", err.Error()))
		return
	}

	c.JSON(http.StatusOK, member)
}

func MemberDelete(c *gin.Context)  {

	memberId, err := strconv.Atoi(c.PostForm("member_id"))

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Invalid ID"))
		return
	}

	err = models.MemberDelete(memberId)

	if err != nil {
		log.Println("member DELETE Error : ", err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Failed to Delete Member"))
		return
	}
	c.JSON(http.StatusOK, response.NewSimpleResponse("Success", "Member Deleted"))
}