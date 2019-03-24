package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-gorp/gorp"
)

func NewDbMap() *gorp.DbMap {
	u := os.Getenv("MYSQL_USER")
	p := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	if u == "" || p == "" {
		panic(errors.New("user or password is empty"))
	}
	source := fmt.Sprintf("%s:%s@tcp(%s:%v)/money_transfer?parseTime=true", u, p, host, port)
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	return &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
}
