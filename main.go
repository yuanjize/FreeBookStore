package main

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/yuanjize/FreeBookStore/router"
	"net/http"
	"encoding/gob"
	"github.com/yuanjize/FreeBookStore/model"
	"github.com/yuanjize/FreeBookStore/middleware"
)

func init() {
	gob.Register(&model.Account{})
}

func main() {
	engine := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions("msession", store),middleware.LoginCheck)
	engine.StaticFS("/static", http.Dir("./static"))
	engine.LoadHTMLGlob("./views/*/*")

	//initTempls()

	//engine.LoadHTMLGlob("./views/account/register.html")
	engine.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello Guest")
	})
	engine.GET("/Token", router.Token)
	engine.GET("/auth", router.Auths)
	engine.GET("/vertify", router.GenerateCaptchaHandler)
	engine.GET("/register", router.GetRegister)
	engine.POST("/register", router.Register)
	engine.POST("/login", router.Login)
	engine.GET("/login", router.GetLogin)
	engine.GET("/captcha", router.Captcha)
	engine.GET("/document/header",router.DocumentPic)
	engine.GET("/book/:identify",router.ReadDocument)
	engine.POST("/uploaddocument/:id",router.UploadDocument)
	engine.POST("/adddocument",router.AddBook)
	engine.GET("/logout",router.Logout)


	setting := engine.Group("/setting/")
	setting.GET("/myproject", router.Myproject)
	setting.GET("/info", router.UserInfo)
	setting.POST("/info", router.ModifyUserInfo)
	setting.POST("/password", router.ModifyPassword)
	setting.GET("/password", router.GetModifyPassword)
	setting.POST("/uploadheader", router.UploadHeader)

	engine.GET("/header", router.GetHeader)

	engine.Run(":8080")
}
func initTempls(engin *gin.Engine) {
	//matches, _ := filepath.Glob("")

	engin.LoadHTMLGlob("./views/*/*.html")
}
