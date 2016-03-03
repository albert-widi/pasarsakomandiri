/*
This auth middlewares automatically check user sessions and redirect the users if not authenticated
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"net/http"
	"github.com/pasarsakomandiri/models"
	"fmt"
	"strings"
)

func AllowUsingToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string

		if c.Request.Method == "POST" {
			token = c.PostForm("token")
		} else if c.Request.Method == "GET" {
			token = c.Query("token")
		}

		if token == "" {
			c.Redirect(http.StatusFound, "/redirected")
			c.Next()
			return
		}

		session := sessions.Default(c)


		if session.Get("id") == nil {
			c.Redirect(http.StatusFound, "/redirected")
			c.Next()
			return
		}

		if !strings.EqualFold(session.Get("token").(string), token) {
			c.Redirect(http.StatusFound, "/redirected")
			c.Next()
			return
		}

		c.Next()
		return
	}
}

func AllowOnlyCashier() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get sessions
		session := sessions.Default(c)

		if session.Get("id") == nil {
			//default redirect
			c.Redirect(http.StatusFound, "/redirected")
			c.Next()
			return
		}

		if session.Get("level") != models.Role_level_cashier {
			c.Redirect(http.StatusFound, "/redirected")
			c.Next()
			return
		}

		c.Next()
		return
	}
}

func AllowOnlyAdministrator() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get sessions
		session := sessions.Default(c)

		if session.Get("id") == nil {
			//default redirect
			c.Redirect(http.StatusFound, "/redirected")
			c.Next()
			return
		}

		if session.Get("level").(int) < models.Role_level_administrator {
			c.Redirect(http.StatusFound, "/redirected")
			c.Next()
			return
		}

		c.Next()
		return
	}
}

//only super user
func AllowOnlySuperUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get sessions
		session := sessions.Default(c)

		if session.Get("id") == nil {
			//default redirect
			c.Redirect(http.StatusFound, "/redirected")
			c.Next()
			return
		}

		if session.Get("level").(int) != models.Role_level_superuser {
			c.Redirect(http.StatusFound, "/redirected")
			c.Next()
			return
		}

		c.Next()
		return
	}
}

//disallow when session is on
func DisallowAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get sessions
		session := sessions.Default(c)

		if session.Get("id") != nil {
			//default redirect
			c.Redirect(http.StatusFound, "/user/user_auth")
		}

		c.Next()
		return
	}
}

//disallow anonymous
func DisAllowAnon() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get sessions
		session := sessions.Default(c)
		if session.Get("id") == nil {
			//default redirect
			fmt.Println("redirected kok gan")
			c.Redirect(http.StatusFound, "/redirected")
		}
		c.Next()

		return
	}
}

/*//allow auth in this level range
func AllowInRange(levelmin int64, levelmax int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		//get sessions
		session := sessions.Default(c)
		c.Next()

		if session.Get("id") == nil {
			//default redirect
			c.Redirect(http.StatusFound, "/")
		}

		if session.Get("level") < levelmin || session.Get("level") > levelmax {
			c.Redirect(http.StatusFound, "/")
		}

		return
	}
}

//allow auth below this level range
func AllowBelow(level int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		//get sessions
		session := sessions.Default(c)
		c.Next()

		if session.Get("id") == nil {
			//default redirect
			c.Redirect(http.StatusFound, "/")
		}

		if session.Get("level") >= level {
			c.Redirect(http.StatusFound, "/")
		}

		return
	}
}

func AllowUpFrom(level int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		//get sessions
		session := sessions.Default(c)
		c.Next()

		if session.Get("id") == nil {
			//default redirect
			c.Redirect(http.StatusFound, "/")
		}

		if session.Get("level") < level {
			c.Redirect(http.StatusFound, "/")
		}

		return
	}
}*/

