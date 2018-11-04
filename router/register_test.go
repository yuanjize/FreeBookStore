package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/yuanjize/FreeBookStore/model"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckEmail(t *testing.T) {
	t.Log(checkEmail("1224500079@qq.com"))
	t.Log(checkEmail("1224500079@gamil.com"))
	t.Log(checkEmail("dnsakjd.@gamil.com"))
}

func TestAddAccount(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "http://localhost:8080/register", nil)
	request.ParseForm()
	request.PostForm.Add("username", "yuanjize222")
	request.PostForm.Add("password1", "456")
	request.PostForm.Add("password2", "456")
	request.PostForm.Add("email", "fuck@123.com")
	request.PostForm.Add("nickname", "hello1222")

	engine := setupRouter()
	engine.ServeHTTP(recorder, request)
	bytes, _ := ioutil.ReadAll(recorder.Body)
	t.Log("hello:", string(bytes))
}

func TestDeleteAccount(t *testing.T) {
	account := model.Account{Id: "c95cc84779de418dae87944472af8c5f"}
	account.Delete()
}

func TestUpdateAccount(t *testing.T) {
	account := model.Account{}
	account.Find("Id", "b72f65dc7cbb4bc9a96ea85f17e9b3b0")
	account.NickName = "kitty"
	account.Update()
}

func setupRouter() *gin.Engine {
	engine := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions("mysession", store))
	engine.GET("/Token", Token)
	engine.GET("/auth", Auths)
	engine.GET("/vertify", GenerateCaptchaHandler)
	engine.POST("/register", Register)
	engine.POST("/login", Login)

	return engine
}

func TestCheckNickName(t *testing.T) {
	account := model.NewAccount()
	account.Find("NickName", "hello1")
	t.Error(account)
}
func TestLogin(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "http://localhost:8080/login", nil)
	request.ParseForm()
	request.PostForm.Add("account", "yuanjize")
	request.PostForm.Add("passwd", "456")

	engine := setupRouter()
	engine.ServeHTTP(recorder, request)
	bytes, _ := ioutil.ReadAll(recorder.Body)
	t.Log("hello:", string(bytes))
}
