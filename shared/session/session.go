package session

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (

)

type Session struct {
	SecretKey string		`json:"SecretKey"`
	Name string			`json:"Name"`
	Options sessions.Options	`json:"Options'`
}

func Configure(r *gin.Engine, s Session) {
	store := sessions.NewCookieStore([]byte(s.SecretKey))
	store.Options(s.Options)
	r.Use(sessions.Sessions("pasarsakomandiri", store))
}

func Instance(c *gin.Context) sessions.Session {
	session := sessions.Default(c)
	return session
}

func Clear(s sessions.Session) {
	s.Clear()
	s.Save()
}