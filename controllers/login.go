package controllers

import (
	"database/sql"
	_ "database/sql"
	_ "encoding/json"
	_ "fmt"
	"log"
	"net/http"
	"strings"
	"time"
	_ "time"

	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
	_ "github.com/jmoiron/sqlx"
	"github.com/pasarsakomandiri/models"
	_ "github.com/pasarsakomandiri/models"
	"github.com/pasarsakomandiri/shared/session"
	"github.com/pasarsakomandiri/shared/token"
	"golang.org/x/crypto/bcrypt"
)

type LoginReturn struct {
	Status   bool
	Username string
	Message  string
}

func Redirected(c *gin.Context) {
	c.HTML(http.StatusFound, "dir_login.tmpl", gin.H{
		"title": "redirected"})
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusFound, "login.tmpl", gin.H{
		"title": "testing"})
}

func flushBan(c *gin.Context) bool {
	slicedIp := strings.Split(c.ClientIP(), ":")
	ipAddress := slicedIp[0]

	err := models.BanFlushAttempt(ipAddress)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func isBanned(c *gin.Context) bool {
	slicedIp := strings.Split(c.ClientIP(), ":")
	ipAddress := slicedIp[0]

	ban, err := models.BanGetInfoByHost(ipAddress)

	if ban.Ban_time == "0000-00-00 00:00:00" {
		return false
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}

		log.Println(err)
		return false
	}

	banTime := now.MustParse(ban.Ban_time)

	if time.Now().Before(banTime) {
		return true
	}

	return false
}

func accumulateAttempt(c *gin.Context) {
	slicedIp := strings.Split(c.ClientIP(), ":")
	ipAddress := slicedIp[0]

	ban, err := models.BanGetInfoByHost(ipAddress)

	if err != nil {
		if err == sql.ErrNoRows {
			//create ban line
			if err = models.BanCreateNewAddress(ipAddress); err != nil {
				log.Println(err)
			}
			return
		}
		log.Println(err)
		return
	}

	ban.Attempt += 1
	//log.Println("Delta: ",now.MustParse(ban.Ban_time).Sub(time.Now()).Seconds())

	if ban.Attempt%5 == 0 {
		ban.Ban_time = time.Now().Add(time.Minute * 5).String()

		if ban.Attempt == 25 {
			ban.Ban_time = time.Now().Add(time.Hour * 24).String()
			//reset the attempt
			ban.Attempt = 0
		}
	}

	if err = models.BanUpdateAddress(ban); err != nil {
		log.Println(err)
	}
}

func LoginAPI(c *gin.Context) {
	session := session.Instance(c)
	var login = LoginReturn{}

	if isBanned(c) {
		login.Status = false
		login.Message = "Please wait to login, you're being restricted"
		c.JSON(http.StatusOK, login)
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")
	login.Username = username
	login.Status = false

	if username == "" || password == "" {
		login.Message = "Username or password cannot be empty"
		c.JSON(http.StatusFound, login)
		return
	}

	user, err := models.UserGetByUsername(username)

	if err != nil {
		if err == sql.ErrNoRows {
			login.Message = "Wrong username or password"
			accumulateAttempt(c)
			c.JSON(http.StatusOK, login)
			return
		}

		login.Message = "System fatal error"
		c.JSON(http.StatusOK, login)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		login.Message = "Wrong username or password"
		accumulateAttempt(c)
		c.JSON(http.StatusOK, login)
		return
	}

	if user.Status == 1 {
		login.Status = false
		login.Message = "User is already online, multiple login attempt failed"
		c.JSON(http.StatusOK, login)
		return
	}

	/*if !strings.EqualFold(password, user.Password) {
		login.Message = "Wrong username or password"
		accumulateAttempt(c)
		c.JSON(http.StatusOK, login)
		return
	}*/

	login.Status = true
	login.Message = "Login Success"

	//create jwt token
	tokenString, err := token.CreateTokenByUserId(user.Id)

	if err != nil {
		log.Println(err)
	}

	//set user status to online
	err = models.UpdateUserStatus(user.Id, 1)

	if err != nil {
		log.Println(err)
	}

	session.Set("id", user.Id)
	session.Set("level", user.Level)
	session.Set("role", user.Role)
	session.Set("token", tokenString)
	session.Save()

	go flushBan(c)

	c.JSON(http.StatusOK, login)
}

func LogoutAPI(c *gin.Context) {
	s := session.Instance(c)

	id := s.Get("id").(int64)

	//set user status to offline
	err := models.UpdateUserStatus(id, 0)

	if err != nil {
		log.Println(err)
	}

	token.ClearTokenSession(s.Get("token").(string))
	session.Clear(s)

	c.Redirect(http.StatusFound, "/")
	return
}

func GetBannedList() {

}
