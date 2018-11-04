package model

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/yuanjize/FreeBookStore/util"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Account struct {
	Id          string    `json:"id" orm:"id"`
	Account     string    `json:"account" orm:"username"`
	Password    string    `json:"password" orm:"password"`
	NickName    string    `json:"nickname" orm:"nickname"`
	CreateAt    time.Time `json:"createat" orm:"createat"`
	UpdateTime  time.Time `json:"uodatetime" orm:"updatetime"`
	Email       string    `json:"email" orm:"email"`
	Role        int       `json:"role" orm:"role"`
	Phone       string    `json:"phone" orm:"phone"`
	Description string    `json:"description" orm:"description"`
	Header      string    `json:"header" orm:"header"`
}


const (
	accountInsert = "INSERT INTO accounts(id,username,password,nickname,email,role,createat,updatetime) VALUES(?,?,?,?,?,?,?,?)"
	accountQuery  = "SELECT id,username,password,nickname,email,role,createat,updatetime,phone,description,header FROM accounts WHERE %s =?"
	accountDelete = "DELETE FROM accounts WHERE id = ?"
	accountUpdate = "UPDATE accounts SET username=?,password=?,nickname=?,email=?,role=?,createat=?,updatetime=?,phone=?,description=?,header=? WHERE id=?"
)

func NewAccount() *Account {
	return &Account{}
}

func (this *Account) Insert() (err error) {
	this.Role = 1
	this.Id = util.UUID()
	this.CreateAt = time.Now()
	this.UpdateTime = time.Now()
	_, err = DB.Exec(accountInsert, this.Id, this.Account, this.Password, this.NickName, this.Email, this.Role, this.CreateAt, this.UpdateTime)
	if err != nil {
		err = errors.Wrap(err, "accountInsert account fail")
	}


	return
}

func (this *Account) Find(field string, value string) (err error) {
	typ := reflect.TypeOf(this).Elem()
	fieldType, ok := typ.FieldByName(field)
	if ok {
		ormTag := fieldType.Tag.Get("orm")
		queryState := fmt.Sprintf(accountQuery, ormTag)
		row := DB.QueryRow(queryState, value)
		err = row.Scan(&this.Id, &this.Account, &this.Password, &this.NickName, &this.Email, &this.Role, &this.CreateAt, &this.UpdateTime,&this.Description,&this.Phone,&this.Header)
	} else {
		err = fmt.Errorf("field %s can`t be find\n", field)
	}
	return

}

func (this *Account) Delete() (err error) {
	_, err = DB.Exec(accountDelete, this.Id)
	return

}

func (this *Account) Update() (err error) {
	this.UpdateTime = time.Now()
	log.Println("Update Accpunt %#v",*this)
	_, err = DB.Exec(accountUpdate, this.Account, this.Password, this.NickName, this.Email, this.Role, this.CreateAt, this.UpdateTime,this.Phone,this.Description,this.Header,this.Id)
	return
}

func (this *Account) Login(account, passwd string) error {
	row, _ := find(gin.H{
		"username": this.Account,
		"password": this.Password,
	})
	acc := this
	err := row.Scan(&acc.Id, &acc.Account, &acc.Password, &acc.NickName, &acc.Email, &acc.Role, &acc.CreateAt, &acc.UpdateTime,&acc.Phone,&acc.Description,&acc.Header)
	return err
}

// except time type
func find(param map[string]interface{}) (*sql.Row, error) {
	if len(param) == 0 {
		return nil, errors.New("find param cannot be nil")
	}
	statement := "SELECT id,username,password,nickname,email,role,createat,updatetime,phone,description,header FROM accounts WHERE "
	slice := make([]string, 0, len(param))
	for k, v := range param {

		switch v.(type) {
		case int:
			{
				slice = append(slice, fmt.Sprintf("%s=%s", k, "\""+strconv.Itoa(v.(int))+"\""))
			}
		case string:
			slice = append(slice, fmt.Sprintf("%s=%s", k, "\""+v.(string))+"\"")
		}
	}
	statement += strings.Join(slice, " and ")
	log.Println("[find] statement:", statement)
	row := DB.QueryRow(statement)
	return row, nil
}
