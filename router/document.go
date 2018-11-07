package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/yuanjize/FreeBookStore/model"
	"github.com/yuanjize/FreeBookStore/config"
	"log"
	"os"
	"io"
	"path/filepath"
	"strconv"
)
var bookPic = "./avatar.jpeg"

func init()  {

}


type AddDocumentRequest struct {
	Bookname    string `form:"book_name"`
	Description string `form:"description"`
	PrivatelyOwned int `form:"privately_owned"`
	Identify string `form:"identify"`

}

//book_name:1212
//identify:dsdasdqqqqqQQQ
//description:dsadasdqwsqsA
//privately_owned:1


func AddBook(c *gin.Context){
	//user, ok := IsLogin(c)
	user ,_:= c.Get("user")

	account := user.(*model.Account)
	result := gin.H{}
	result["errcode"]="1"

	request := &AddDocumentRequest{}
	err := c.Bind(request)
	if err!=nil{
		result["msg"]=err.Error()
		c.JSON(http.StatusOK,result)
		return
	}

	document := model.NewDocument()
	document.Find("identify",request.Identify)
	if document.Id != ""{
		result["msg"]="项目标志已存在"
		c.JSON(http.StatusOK,result)
		return
	}
	document.Description = request.Description
	document.Bookname = request.Bookname
	document.Identify = request.Identify
	document.PrivatelyOwned = request.PrivatelyOwned
	document.Owner = account
	document.Score = 0
	document.FavoriteCount = 0
	document.ReadCount = 0
	document.SectionCount = 0
	document.Picture = bookPic
	err = document.Insert()
	if err!=nil{
		log.Println("添加书籍失败:",err)
		result["msg"]="添加书籍失败"
		c.JSON(http.StatusOK,result)
		return
	}
	result["msg"]="书籍已添加"
	result["errcode"]="0"
	result["data"]=document
	c.JSON(http.StatusOK,result)
}

func Myproject(c *gin.Context){
	user ,_:= c.Get("user")

	private := c.Query("private")
	pri := 0
	if private!=""{
		pri ,_ = strconv.Atoi(private)
	}
	err, documents := model.QueryAllDocument(user.(*model.Account),pri)
	if err != nil{
		documents = nil
	}
	c.HTML(http.StatusOK,"book/index.html",gin.H{
		"Member":user,
		"Private":pri,
		"BaseUrl":config.Host,
		"Result" :documents,
	})
}

func DocumentPic(c *gin.Context) {
	id := c.Query("id")
	document := model.NewDocument()
	err := document.Find("Id", id)
	if err!=nil{
		log.Println("cannot find document pic:",err)
	}else{
		c.File(document.Picture)
	}
}

func ReadDocument(c *gin.Context)  {
	documentPath := c.Param("identify")
	document := model.NewDocument()
	document.Find("Identify",documentPath)
	c.File(document.Url)
}

func UploadDocument(c *gin.Context) {

	user ,_:= c.Get("user")
	documentId := c.Param("id")
	form, _ := c.MultipartForm()
	files := form.File
	result := gin.H{"msg":"ok","errcode":0}
	docs,ok := files["doc"]
	if !ok || len(docs)==0{
		result["errcode"] = 1
		result["msg"] = "上传文件失败"
		c.JSON(http.StatusOK,result)
		return
	}

	document := model.NewDocument()
	err := document.Find("Id", documentId)
	if err!=nil{
		log.Println("no this document:",err)
		result["errcode"] = 1
		result["msg"] = "接收文件失败"
		c.JSON(http.StatusOK,result)
		return
	}


	account := user.(*model.Account)
	doc := docs[0]
	dir := "./docs/"+account.Id+"/"
	os.MkdirAll(dir,0777)
	fileName,_ := filepath.Abs(dir+document.Bookname)
	log.Println("upload filename:",fileName)
	os.MkdirAll("./docs",0777)
	src, err := doc.Open()
	if err!=nil{
		log.Println("doc open fail:",err)
		result["errcode"] = 1
		result["msg"] = "接收文件失败"
		c.JSON(http.StatusOK,result)
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err!=nil{
		log.Println("Create FIle err:",err)
		result["errcode"] = 1
		result["msg"] = "创建文件失败"
		c.JSON(http.StatusOK,result)
		return
	}
	defer dst.Close()
	written, err := io.Copy(dst, src)
	if written!=doc.Size || err!=nil{
		log.Println("Copy FIle err:",err)
		result["errcode"] = 1
		result["msg"] = "创建文件失败"
		return
	}

	document.Url = fileName
	err = document.Update()
	if err!=nil{
		log.Println("document.Update Fail:",err)
		result["errcode"] = 1
		result["msg"] = "接收文件失败"
		c.JSON(http.StatusOK,result)
		return
	}
	c.JSON(http.StatusOK,result)
	log.Println("upload file ok")
	return
}

