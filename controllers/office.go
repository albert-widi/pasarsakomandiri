package controllers
import (
	"github.com/gin-gonic/gin"
	"github.com/pasarsakomandiri/shared/session"
	"net/http"
)


func OfficePage(c *gin.Context) {
	session := session.Instance(c)
	c.HTML(http.StatusFound, "office.tmpl", gin.H{"title":"Office", "token":session.Get("token")})
}
