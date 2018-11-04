package router

import (
	"crypto/tls"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/pkg/errors"
	"github.com/yuanjize/FreeBookStore/auth"
	"github.com/yuanjize/FreeBookStore/config"
	"github.com/yuanjize/FreeBookStore/model"
	"github.com/yuanjize/FreeBookStore/util"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func Token(context *gin.Context) {
	code := context.Query("code")
	log.Println("[/Token] code:", code)
	headers := map[string]string{"Accept": gin.MIMEJSON}
	request := util.RequestGet(config.GetAuthHost(), gin.H{
		"client_id":     config.GetClientId(),
		"client_secret": config.GetClientSecret(),
		"code":          code,
	}, headers)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Do(request)
	if err != nil {
		log.Panicln("[/Token] response body err:", err)
	}
	body, err := ioutil.ReadAll(response.Body)
	log.Println("[/Token] response body:", string(body))
	token := &auth.GithubOAuth2{}
	json.Unmarshal(body, token)

}
func Auths(context *gin.Context) {
	request := util.RequestGet(config.GetCodeHost(), gin.H{
		"client_id":    config.GetClientId(),
		"redirect_uri": "http://localhost:8080/Token",
	}, nil)

	context.Redirect(http.StatusTemporaryRedirect, request.URL.String())
}

func GetLogin(c *gin.Context) {
	captchaSwitch := config.CaptchaSwitch()
	c.HTML(http.StatusOK, "login.html", gin.H{
		"CaptchaOn": captchaSwitch,
		"captcha":   template.URL(genCaptcha(c)),
	})
}

func Login(c *gin.Context) {
	account := c.PostForm("account")
	passwd := c.PostForm("passwd")
	captcha := c.PostForm("captcha")
	session := sessions.Default(c)
	mCaptchaId := session.Get("captchaId")

	if !checkCaptcha(mCaptchaId.(string), captcha) {
		log.Printf("login fail,the captchashould be:%s ,but is %s\n", checkCaptcha, captcha)
		c.JSON(http.StatusOK, gin.H{
			"msg":     "登录失败：验证码错误",
			"errcode": "1",
		})
		return
	}
	user := &model.Account{Account: account, Password: passwd}
	err := user.Login(account, passwd)

	if user.Id != "" {
		if err!=nil{
			log.Println("Login err:", err)
		}
		SaveUser(c,user)
		c.JSON(http.StatusOK, gin.H{
			"msg":     "登陆成功",
			"errcode": "0",
		})
	} else {
		log.Println("Login err:", err)
		c.JSON(http.StatusOK, gin.H{
			"msg":     "登录失败：请检查用户名和密码,",
			"errcode": "1",
		})
	}
}

func GetRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"captcha": template.URL(genCaptcha(c)),
	})
}

func Register(c *gin.Context) {
	userName := c.PostForm("username")
	password1 := c.PostForm("password1")
	password2 := c.PostForm("password2")
	email := c.PostForm("email")
	nickName := c.PostForm("nickname")
	captcha := c.PostForm("captcha")
	session := sessions.Default(c)
	captchaId := session.Get("captchaId")

	if captchaId != nil {
		if !checkCaptcha(captchaId.(string), captcha) {
			err := errors.New("Verfiy Code Failed")
			c.Error(err)
			c.JSON(http.StatusOK, gin.H{
				"msg":      err.Error(),
				"errcode:": "1",
			})
			return
		}
	}

	if password1 != password2 {
		err := errors.New("twice input of password is not equal")
		c.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg":      err.Error(),
			"errcode:": "1",
		})
		return
	}
	if !checkEmail(email) {
		err := errors.New("Invalid email")
		c.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg":      err.Error(),
			"errcode:": "1",
		})
		return
	}

	if !checkRepeatField("Email", email) {
		err := errors.New("Email has been registered")
		c.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg":      err.Error(),
			"errcode:": "1",
		})
		return
	}

	if !checkRepeatField("NickName", nickName) {
		err := errors.New("nickname has existed")
		c.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg":      err.Error(),
			"errcode:": "1",
		})
		return
	}

	account := model.NewAccount()
	account.Email = email
	account.Password = password1
	account.Account = userName
	account.NickName = nickName
	err := account.Insert()
	if err != nil {
		err := errors.Wrap(err, "create account failed")
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":      err.Error(),
			"errcode:": "1",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":      "register success",
			"errcode:": "0",
			"location": "http://localhost:8080/login",
		})
	}
}

func GenerateCaptchaHandler(c *gin.Context) {
	// get session
	//session := sessions.Default(c)
	//captchaConfig := util.GetCaptchaConfig()
	////create base64 encoding captcha
	////创建base64图像验证码
	//config := captchaConfig.ConfigCharacter
	////GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	//captchaId, digitCap := base64Captcha.GenerateCaptcha(captchaConfig.Id, config)
	//base64Png := base64Captcha.CaptchaWriteToBase64Encoding(digitCap)
	//session.Set("captchaId", captchaId)
	//return
}

func genCaptcha(c *gin.Context) string {
	// get session
	session := sessions.Default(c)
	session.Options(sessions.Options{
		Path: "/",
	})
	captchaConfig := util.GetCaptchaConfig()
	//create base64 encoding captcha
	//创建base64图像验证码
	config := captchaConfig.ConfigCharacter
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	captchaId, digitCap := base64Captcha.GenerateCaptcha(captchaConfig.Id, config)
	base64Png := base64Captcha.CaptchaWriteToBase64Encoding(digitCap)
	session.Set("captchaId", captchaId)
	session.Save()
	//log.Println("gen captch:",base64Png)
	return base64Png
}

func Captcha(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"errcode": "0",
		"captcha": genCaptcha(c),
	})
}

func checkCaptcha(captchaId, captcha string) bool {
	return util.VerfiyCaptcha(captchaId, captcha)

}

func checkEmail(email string) (matched bool) {
	log.Println("[check email]:", email)
	matched, err := regexp.MatchString("^\\w+@\\w+\\.\\w{2,4}$", email)
	if err != nil {
		log.Println("[checkEmail] err:", err)
	}
	return
}

//func checkPasswd(passwd string)bool{
//	if len(passwd)
//	return true
//}
func checkRepeatField(fieldName, value string) bool {
	if len(fieldName) == 0 {
		return false
	}
	account := model.NewAccount()
	account.Find(fieldName, value)
	return account.Id == ""
}
