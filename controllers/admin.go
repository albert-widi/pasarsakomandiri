package controllers
import (
	"github.com/gin-gonic/gin"
	"github.com/pasarsakomandiri/shared/session"
	"net/http"
)

func AdminPage(c *gin.Context) {
	session := session.Instance(c)
	c.HTML(http.StatusFound, "admin.tmpl", gin.H{"title":"Cahiser", "token":session.Get("token")})
}