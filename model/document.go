package model

import (
	"fmt"
	"github.com/yuanjize/FreeBookStore/util"
	"reflect"
	"time"
	"github.com/yuanjize/FreeBookStore/config"
	"log"
	"database/sql"
)

type Document struct {
	Id            string   `json:"id" orm:"id"`
	Bookname      string   `json:"bookname" orm:"bookname"`
	SectionCount  int      `json:"sectioncount" orm:"sectioncount"`
	Owner         *Account `json:"owner" orm:"owner"`
	ReadCount     int      `json:"readcount" orm:"readcount"`
	FavoriteCount int      `json:"favoritecount" orm:"favoritecount"`
	Description   string   `json:"description" orm:"description"`
	Score         int      `json:"score" orm:"score"`
	Url           string   `json:"url" orm:"url"`
	PrivatelyOwned int     `json:"privately_owned" orm:"privately_owned"`
	Identify      string   `json:"identify" orm:"identify"`
	Picture       string   `json:"picture" orm:"picture"`
	CreateAt       time.Time   `json:"createat" orm:"createat"`
	UpdateAt       time.Time   `json:"last_modify_text" orm:"last_modify_text"`
	HeaderUrl      string	   `json:"headerurl"`

}

const (
	documentInsert = "INSERT INTO document(id,bookname,sectioncount,owner,readcount,favoritescount,description,score,url,privately_owned,identify,picture,createat,last_modify_text) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	documentQuery  = "SELECT id,bookname,sectioncount,owner,readcount,favoritescount,description,score,url,privately_owned,identify,picture,createat,last_modify_text FROM document WHERE %s=?"
	documentUpdate = "UPDATE document SET id=?,bookname=?,sectioncount=?,owner=?,readcount=?,favoritescount=?,description=?,score=?,url=?,privately_owned=?,identify=?,picture=?,createat=?,last_modify_text=? WHERE id = ?"
	documentSelectAll = "SELECT id,bookname,sectioncount,owner,readcount,favoritescount,description,score,url,privately_owned,identify,picture,createat,last_modify_text FROM document WHERE privately_owned = ?"
)

func NewDocument() *Document {
	return &Document{}
}

func (this *Document) Insert() (err error) {
	this.Id = util.UUID()
	this.UpdateAt = time.Now()
	this.CreateAt = this.UpdateAt
	if this.Score == 0 {
		this.Score = 4
	}
	_, err = DB.Exec(documentInsert, this.Id, this.Bookname, this.SectionCount, this.Owner.Id, this.ReadCount, this.FavoriteCount, this.Description, this.Score, this.Url,this.PrivatelyOwned,this.Identify,this.Picture,this.CreateAt,this.UpdateAt)
	return
}

func (this *Document) Update() (err error) {
	this.UpdateAt = time.Now()
	_, err = DB.Exec(documentUpdate, this.Id, this.Bookname, this.SectionCount, this.Owner.Id, this.ReadCount, this.FavoriteCount, this.Description, this.Score, this.Url,this.PrivatelyOwned,this.Identify,this.Picture,this.CreateAt,this.UpdateAt,this.Id)
	return
}

func (this *Document) Find(fieldName, value string) (err error) {
	typ := reflect.TypeOf(this).Elem()
	fieldType, ok := typ.FieldByName(fieldName)
	if ok {
		ormTag := fieldType.Tag.Get("orm")
		queryState := fmt.Sprintf(documentQuery, ormTag)
		log.Println("tags",fieldType.Tag," queryState:",queryState," Value:",value," ormTag:",ormTag)
		row := DB.QueryRow(queryState, value)
		ownerId := ""
		err = row.Scan(&this.Id, &this.Bookname, &this.SectionCount, &ownerId, &this.ReadCount, &this.FavoriteCount, &this.Description, &this.Score, &this.Url,&this.PrivatelyOwned,&this.Identify,&this.Picture,&this.CreateAt,&this.UpdateAt)
		if err != nil {
			return
		}
		account := &Account{}
		account.Find("Id", ownerId)
		//if e != nil {
		//	err = errors.Wrap(err, e.Error())
		//	return
		//}
		this.Owner = account
	} else {
		err = fmt.Errorf("field %s can`t be find\n", fieldName)
	}
	return
}

func QueryAllDocument(user *Account,pri int)(err error,documents []*Document){
	documents = make([]*Document,0,10)
	var rows *sql.Rows
	rows, err = DB.Query(documentSelectAll,pri)
	if err!=nil{
		return
	}
	ownerId := ""
	for rows.Next(){
		this := NewDocument()
		err = rows.Scan(&this.Id, &this.Bookname, &this.SectionCount, &ownerId, &this.ReadCount, &this.FavoriteCount, &this.Description, &this.Score, &this.Url,&this.PrivatelyOwned,&this.Identify,&this.Picture,&this.CreateAt,&this.UpdateAt)
		if err!=nil{
			return
		}
		this.Owner = user
		this.HeaderUrl = config.Host+"document/header?id="+this.Id
		documents = append(documents,this)
	}
	return nil,documents
}
