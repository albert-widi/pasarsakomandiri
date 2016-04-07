package session

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
    _"time"
	"github.com/pasarsakomandiri/shared/tools"
)

var (

)

type Session struct {
	SecretKey string		`json:"SecretKey"`
	Name string			`json:"Name"`
	Options sessions.Options	`json:"Options"`
}

func Configure(r *gin.Engine, s Session) {
    //alterfnative, configure unique secret key every time server start up
    timeKey := tools.RandomString(20)
	store := sessions.NewCookieStore([]byte(timeKey))
	store.Options(s.Options)
	r.Use(sessions.Sessions("pasarsakomandiri", store))
}

func Instance(c *gin.Context) sessions.Session {
	session := sessions.Default(c)
    //fmt.Printf("%+v", session)
	return session
}

func Clear(s sessions.Session) {
	s.Clear()
	s.Save()
}