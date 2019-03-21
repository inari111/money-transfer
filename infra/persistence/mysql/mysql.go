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
	host := "db"
	port := 3306
	if u == "" || p == "" {
		panic(errors.New("user or password is empty"))
	}
	// TODO: ホスト名なども環境変数で制御する
	source := fmt.Sprintf("%s:%s@tcp(%s:%v)/money_transfer?parseTime=true", u, p, host, port)
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	return &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
}
