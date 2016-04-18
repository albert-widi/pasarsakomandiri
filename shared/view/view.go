package view

import (
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/pasarsakomandiri/shared/session"
)

var (
	viewInfo View
)

//View struct to render
type View struct {
	Extension string
	Folder    string
	Name      string
	Variables map[string]interface{}
	Context   *gin.Context
}

//Configure view when server up
func Configure(vi View) {
	viewInfo = vi
}

//New view
func New(c *gin.Context) *View {
	vi := &View{}
	vi.Variables = make(map[string]interface{})
	vi.Variables["AuthLevel"] = "Guest"
	vi.Variables["Username"] = "Guest"

	//Configuring auth level
	sess := session.Instance(c)
	if tmpLevel := sess.Get("Level").(int); tmpLevel != 0 {

		if tmpLevel == 2 {
			vi.Variables["AuthLevel"] = "Office"
		}

		if tmpLevel > 2 {
			vi.Variables["AuthLevel"] = "Admin"
		}

		log.Println(vi.Variables["AuthLevel"])

		//vi.Variables["AuthLevel"] = sess.Get("Level")
		vi.Variables["Username"] = sess.Get("Username")
		vi.Variables["Token"] = sess.Get("Token")
	}

	vi.Context = c
	vi.Extension = viewInfo.Extension

	return vi
}

//Render the view
func (v *View) Render() {
	v.Context.HTML(http.StatusOK, v.Name+"."+v.Extension, v.Variables)
}
