package controllers
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/pasarsakomandiri/shared/session"
)

func CashierPage(c *gin.Context) {
	session := session.Instance(c)
	c.HTML(http.StatusFound, "cashier.tmpl", gin.H{"title":"Cahiser", "token":session.Get("token")})
}