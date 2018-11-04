package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuanjize/FreeBookStore/model"
	"log"
	"net/http"
	"os"
	"io"
	"path/filepath"
	"github.com/yuanjize/FreeBookStore/config"
	"path"
)

func UserInfo(c *gin.Context) {
	user, ok := IsLogin(c)

	if !ok {
		c.Redirect(http.StatusTemporaryRedirect, "/login") //login
	} else {
		account := user.(*model.Account)
		//bytes, er := json.Marshal(account)
		c.HTML(http.StatusOK,"setting/index.html", gin.H{
			"Member": account,
		})
		SaveUser(c,account)
	}
}

func ModifyUserInfo(c *gin.Context) {
	user, ok := IsLogin(c)

	if !ok {
		c.Redirect(http.StatusTemporaryRedirect, "/login") //login
	} else {
		account := user.(*model.Account)
		eamil := c.PostForm("email")
		phone := c.PostForm("phone")
		description := c.PostForm("description")
		account.Description = description
		account.Phone = phone
		account.Email = eamil
		err := account.Update()
		if err == nil {
			SaveUser(c,account)
			c.JSON(http.StatusOK, gin.H{
				"errcode": "0",
				"msg":     "修改成功",
			})
		} else {
			c.Error(err)
			log.Print("[ModifyUserInfo] err:", err)
			c.JSON(http.StatusOK, gin.H{
				"errcode": "1",
				"msg":     "修改失败",
			})
		}
	}
}

func GetModifyPassword(c *gin.Context) {
	login, ok := IsLogin(c)
	if !ok || login==nil{
		c.Redirect(http.StatusTemporaryRedirect,"/login.html")
	}else{
		c.HTML(http.StatusOK, "setting/password.html",gin.H{
			"Member":login.(*model.Account),
		})
	}

}

func UploadHeader(c *gin.Context) {
	logger := log.New(os.Stderr, "[UploadHeader]", log.Flags())
	form, err := c.MultipartForm()
	result := gin.H{}
	login, ok := IsLogin(c)
	if !ok {
		logger.Println("Please login")
		result["errcode"] = "1"
		result["msg"] = "请登录"
		c.JSON(http.StatusOK, result)
		return
	}
	if err != nil {
		logger.Println("Parse File fail")
		result["errcode"] = "1"
		result["msg"] = err.Error()
	} else if len(form.File["image-file"]) == 0 {
		result["errcode"] = "1"
		result["msg"] = "服务器没有收到文件"
	} else {
		result["errcode"] = "0"
		result["msg"] = "ok"
		file := form.File["image-file"][0]
		src, err := file.Open()
		defer src.Close()
		if err != nil {
			result["errcode"] = "1"
			result["msg"] = err.Error()
		} else {
			dir := "./headers/"
			filePath,_ := filepath.Abs(dir + file.Filename)
			os.MkdirAll(dir, 0777)
			os.Chmod(dir,0777)
			openFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
			defer openFile.Close()
			if err != nil {
				result["errcode"] = "1"
				result["msg"] = err.Error()
			} else {
				count, err := io.Copy(openFile, src)
				openFile.Sync()
				if err != nil {
					result["errcode"] = "1"
					result["msg"] = err.Error()
				} else if count != file.Size {
					result["errcode"] = "1"
					result["msg"] = "文件接收失败"
				} else {
					account := login.(*model.Account)
					account.Header, _ = filepath.Abs(filePath)
					err := account.Update()
					if err != nil {
						logger.Println("Update account err:",err)
						result["errcode"] = "1"
						result["msg"] = "服务器错误"
					} else {
						SaveUser(c,account)
						result["data"] = path.Join(config.Host, "header")
					}
				}
			}
		}
	}
	c.JSON(http.StatusOK,result)
}

func GetHeader(c *gin.Context)  {
	user, ok := IsLogin(c)
	if!ok{
		c.Redirect(http.StatusTemporaryRedirect,"/login");
	}else{
		account := user.(*model.Account)
		picLocation,_ := filepath.Abs("./avatar.jpeg")
		if account.Header!=""{
			picLocation = account.Header
		}
		c.File(picLocation)
	}
}

func ModifyPassword(c *gin.Context) {
	user, ok := IsLogin(c)
	if !ok {
		c.Redirect(http.StatusTemporaryRedirect, "/login") //login
	} else {
		oldPassed := c.PostForm("oldpasswd")
		newPasswd1 := c.PostForm("newpasswd1")
		newPasswd2 := c.PostForm("newpasswd2")
		account := user.(*model.Account)
		if oldPassed != account.Password {
			c.JSON(http.StatusOK, gin.H{
				"errcode": "1",
				"msg":     "原始密码错误",
			})
			return
		}
		if newPasswd1 != newPasswd2 {
			c.JSON(http.StatusOK, gin.H{
				"errcode": "1",
				"msg":     "两次密码输入不一致",
			})
			return
		}
		account.Password = newPasswd1
		err := account.Update()
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"errcode": "1",
				"msg":     "Server Error",
			})
		} else {
			SaveUser(c,account)
			c.JSON(http.StatusOK, gin.H{
				"errcode": "0",
				"msg":     "密码更改成功",
			})
		}
	}
}
