package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// connect mysql
var DB *sql.DB

func init() {

	//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]

	dataSource := fmt.Sprintf("%s:%s@%s(%s)/%s?parseTime=true", "root", "123", "tcp", "localhost", "bookstore")
	log.Println("[model] DataSource: ", dataSource)
	log.Println("[model] Connecting database....")
	var err error
	DB, err = sql.Open("mysql", dataSource)
	if err != nil {
		log.Panicln("connect database fail ", err)
	}
	log.Println("[model] Connect database OK!")
}
