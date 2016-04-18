package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pasarsakomandiri/models"
	"github.com/pasarsakomandiri/shared/session"
)

func RedirectUserSession(c *gin.Context) {
	sessions := session.Instance(c)
	department := sessions.Get("department").(string)
	level := sessions.Get("Level").(int)

	//go to super user portal
	if level == models.Role_level_superuser {

	}

	if department == models.DepartmentParking {
		redirectParking(c, level)
	} else if department == models.DepartmentToycar {
		redirectToyCar(c, level)
	}
}

func redirectParking(c *gin.Context, level int) {
	if level == models.Role_level_cashier {
		c.Redirect(http.StatusFound, "/cashier")
	}

	if level == models.Role_level_office {

	}

	if level == models.Role_level_administrator {
		c.Redirect(http.StatusFound, "/admin")
	}
}

func redirectToyCar(c *gin.Context, level int) {

}
