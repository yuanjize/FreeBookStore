package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yuanjize/FreeBookStore/model"
	"log"
)

func IsLogin(c *gin.Context) (interface{}, bool) {
	s := c.Request.Cookies()
	log.Println("cookies:", s)
	session := sessions.Default(c)
	user := session.Get("user")
	log.Println("user:",user)
	if user == nil {
		return user, false
	} else {
		account := user.(*model.Account)
		return account, true
	}
}

func SaveUser(c *gin.Context,account *model.Account)  {
	session := sessions.Default(c)
	session.Clear()
	session.Set("user",account)
	session.Save()
}
