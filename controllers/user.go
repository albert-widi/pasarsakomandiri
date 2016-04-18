package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
	"github.com/pasarsakomandiri/models"
	"github.com/pasarsakomandiri/shared/response"
	"github.com/pasarsakomandiri/shared/session"
	"golang.org/x/crypto/bcrypt"
)

type UserResponse struct {
	Response response.SimpleResponse
	Data     models.User
}

func UserRegisterPages(c *gin.Context) {
	session := session.Instance(c)
	c.HTML(http.StatusFound, "user_register.tmpl", gin.H{"title": "Register User", "token": session.Get("token")})
}

//determine redirect of users

func UserListPages(c *gin.Context) {
	session := session.Instance(c)
	c.HTML(http.StatusFound, "user_list.tmpl", gin.H{"title": "User List", "token": session.Get("token")})
}
func UserEditPages(c *gin.Context) {
	session := session.Instance(c)
	c.HTML(http.StatusFound, "update_user.tmpl", gin.H{"title": "Edit User", "token": session.Get("token")})
}

func UserSessionRedirect(c *gin.Context) {
	session := session.Instance(c)
	userlevel := session.Get("Level").(int)

	//redirect user to cashier
	if userlevel == models.Role_level_cashier {
		c.Redirect(http.StatusFound, "/cashier")
	}

	//default
	c.Redirect(http.StatusFound, "/admin")

	return
}

func UserGetAllRoleLimitLevel(c *gin.Context) {
	session := session.Instance(c)
	executor := session.Get("Level").(int)

	userRoles, err := models.RoleGetAllLimitByLevel(executor)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Now rows returned"))
			return
		}
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, userRoles)
	return
}

func UserGetAllRoleAPI(c *gin.Context) {
	userRoles, err := models.RoleGetAllRole()

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Now rows returned"))
			return
		}
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, userRoles)
	return
}

func UserGetInfoAPI(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Query("id"), 10, 64)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Invalid user id"))
	}

	user, err := models.UserGetByID(userID)

	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "User not found"))
			return
		}
	}

	c.JSON(http.StatusOK, user)
}

func UserUpdateAPI(c *gin.Context) {

	userId, err := strconv.ParseInt(c.PostForm("id"), 10, 64)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Invalid user id"))
	}

	_, err = models.UserGetByID(userId)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Invalid user id"))
	}

	username := c.PostForm("username")
	password := c.PostForm("password")
	confirmpass := c.PostForm("confirmpassword")
	role := c.PostForm("role")
	description := c.PostForm("description")

	if username == "" || role == "" {
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Required item cannot be null"))
		return
	}

	if password != "" {
		if password != confirmpass {
			c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Password and confirm password does not match"))
			return
		}
	}

	bytePass := []byte(password)
	hashedPass, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)

	user := models.User{}
	user.Id = userId
	user.Username = username
	user.Password = string(hashedPass)
	user.Role = role
	user.Description = description

	err = models.UserUpdate(user)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Failed to update user"))
		return
	}

	c.JSON(http.StatusOK, response.NewSimpleResponse("Success", "Update Success"))
}

func UsertGetAllLimitLevel(c *gin.Context) {
	//session
	session := session.Instance(c)
	executor := session.Get("Level").(int)

	users, err := models.UserGetAllLimitByLevel(executor)

	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, users)
}

func UserGetAllAPI(c *gin.Context) {
	users, err := models.UserGetAll()

	if err != nil {
		log.Println("UserGetALlAPI Error: ", err)
		return
	}
	c.JSON(http.StatusOK, users)
}

//---- experimental code

func UserSinceAPI(c *gin.Context) {
	username := c.Query("username")

	user, err := models.UserGetByUsername(username)

	if err != nil {
		log.Println("Error: ", err)
	}

	//fmt.Println("Date created: ", user.Created_date)
	//fmt.Println("Date: ", now.MustParse(user.Created_date))
	date := now.MustParse(user.Created_date)
	delta := time.Now().Sub(date)

	fmt.Println("Delta: ", delta)
}
