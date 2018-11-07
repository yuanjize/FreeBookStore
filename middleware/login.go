package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
	"log"
	"strings"
)

func LoginCheck(c *gin.Context)  {
	if strings.Contains(c.Request.URL.Path,"login"){
		c.Next()
		return
	}
	session := sessions.Default(c)
	user := session.Get("user")
	log.Println("Login Check:",user)
	if user==nil{
		c.Redirect(http.StatusSeeOther,"/login")
		return
	}
	c.Set("user",user)
	c.Next()

}
