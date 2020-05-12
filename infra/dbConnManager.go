package infra

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
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
	zap.L().Info("pinging database...")
	err = db.Ping()
	if err != nil {
		zap.L().Fatal("Connection could not be established with db")
	} else {
		zap.L().Info("ping success")
	}
	return db
}
