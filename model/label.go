package model

import (
	"fmt"
	"github.com/yuanjize/FreeBookStore/util"
	"reflect"
)

type Label struct {
	Id   string `json:"id" orm:"id"`
	Name string `json:"name" orm:"name"`
}

const (
	labelInsert = "INSERT INTO label(id,name) VALUES(?,?)"
	labelQuery  = "SELECT id,name FROM label WHERE %s=?"
)

func NewLabel() *Label {
	return &Label{}
}

func (this *Label) Insert() error {
	this.Id = util.UUID()
	_, err := DB.Exec(labelInsert, this.Id, this.Name)
	return err
}

func (this *Label) Find(fieldName, value string) (err error) {
	typ := reflect.TypeOf(this).Elem()
	fieldType, ok := typ.FieldByName(fieldName)
	if ok {
		ormTag := fieldType.Tag.Get("orm")
		queryState := fmt.Sprintf(labelQuery, ormTag)
		row := DB.QueryRow(queryState, value)
		err = row.Scan(&this.Id, &this.Name)
	} else {
		err = fmt.Errorf("field %s can`t be find\n", fieldName)
	}
	return
}
