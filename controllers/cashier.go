package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pasarsakomandiri/shared/view"
)

func CashierPage(c *gin.Context) {
	v := view.New(c)
	v.Name = "cashier"
	v.Render()
	//session := session.Instance(c)
	//c.HTML(http.StatusFound, "cashier.tmpl", gin.H{"title":"Cahiser", "token":session.Get("token")})
}
