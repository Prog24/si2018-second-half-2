package repositories

import (
	"fmt"
	"os"

	"github.com/go-xorm/xorm"

	_ "github.com/go-sql-driver/mysql"
)

var engine *xorm.Engine

func init() {
	fmt.Println("init xorm Engine")
	var err error

	var hostname string
	var dbname string
	var username string
	var password string
	var port string

	hostname = os.Getenv("DB_HOSTNAME")
	dbname = os.Getenv("DB_DBNAME")
	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	port = os.Getenv("DB_PORT")

	engine, err = xorm.NewEngine("mysql", username+":"+password+"@tcp("+hostname+":"+port+")/"+dbname)

	engine.ShowSQL(true)
	if err != nil {
		panic(err)
	}
}
