package model

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/yuanjize/FreeBookStore/util"
	"reflect"
	"time"
)

//id char(32),
//data varchar(100) NOT NULL,
//ownerid char(32),
//createat DATETIME NOT NULL ,
//PRIMARY KEY (id)

type Comment struct {
	Id       string    `json:"id" orm:"id"`
	Data     string    `json:"data" orm:"data"`
	Owner    *Account  `json:"owner" orm:"owner"`
	CreateAt time.Time `json:"createat" orm:"createat"`
}

const (
	commentInsert = "INSERT INTO comments(id,data,owner,createat) VALUES(?,?,?,?)"
	commentQuery  = "SELECT id,data,owner,createat FROM label WHERE %s=?"
)

func NewComment() *Comment {
	return &Comment{}
}

func (this *Comment) Insert() (err error) {
	this.Id = util.UUID()
	this.CreateAt = time.Now()
	_, err = DB.Exec(commentInsert, this.Id, this.Data, this.Owner.Id, this.CreateAt)
	return
}

func (this *Comment) Find(fieldName, value string) (err error) {
	typ := reflect.TypeOf(this).Elem()
	fieldType, ok := typ.FieldByName(fieldName)
	if ok {
		ormTag := fieldType.Tag.Get("orm")
		queryState := fmt.Sprintf(commentQuery, ormTag)
		row := DB.QueryRow(queryState, value)
		ownerId := ""
		err = row.Scan(&this.Id, &this.Data, &ownerId, &this.CreateAt)
		if err != nil {
			return
		}
		account := &Account{}
		e := account.Find("id", ownerId)
		if e != nil {
			err = errors.Wrap(err, e.Error())
			return
		}
		this.Owner = account
	} else {
		err = fmt.Errorf("field %s can`t be find\n", fieldName)
	}
	return
}
