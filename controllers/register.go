package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/pasarsakomandiri/models"
	"database/sql"
	"github.com/pasarsakomandiri/shared/response"
	"time"
	"github.com/pasarsakomandiri/shared/session"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func RegisterUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	role := c.PostForm("role")
	description := c.PostForm("description")

	if username == "" || password == "" || role == "" {
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Required items cannot be null"))
		return
	}

	level, err := models.UserLevelByRole(role)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Invalid role"))
			return
		}

		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "System fatal error"))
		return
	}

	bytePass := []byte(password)
	hashedPass, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Failed to register user"))
		return;
	}

	//executor, err := strconv.ParseInt(session.Instance(c).Get("id"), 10, 64)
	user := models.User{}
	user.Username = username
	user.Password = string(hashedPass)
	user.Level = level
	user.Role = role
	user.Description = description
	user.Created_date = time.Now().String()
	user.Created_by = session.Instance(c).Get("id").(int64)

	result, err := models.UserCreateNew(user)

	if err != nil {
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Failed to register user"))
		return;
	}

	user.Id, _ = result.LastInsertId()
	c.JSON(http.StatusOK, response.NewSimpleResponse("Success", "User created"))
	return
}
