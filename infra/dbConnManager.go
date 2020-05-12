package infra

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"papertrader.io/backoffice/config"
	"time"
)

func DbConn(config *config.GeneralConfig) (db *sqlx.DB) {
	db, err := sqlx.Open(config.Database.Driver, config.Database.Host)
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Second)
	db.SetMaxIdleConns(0)

	//try ping
	fmt.Println("pinging database...")
	err = db.Ping()
	if err != nil {
		fmt.Println("Connection could not be established with db")
		panic(err.Error())
	} else {
		fmt.Println("success")
	}
	return db
}
