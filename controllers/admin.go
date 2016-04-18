package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pasarsakomandiri/shared/view"
)

func AdminPage(c *gin.Context) {
	v := view.New(c)
	v.Name = "admin"
	v.Render()
	//session := session.Instance(c)
	//c.HTML(http.StatusFound, "admin.tmpl", gin.H{"title": "Cahiser", "token": session.Get("token")})
}
