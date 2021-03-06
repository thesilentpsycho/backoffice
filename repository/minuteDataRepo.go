package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"papertrader.io/backoffice/dao"
)

const (
	MINUTE_DATA_TABLE_NAME = "minute_data"
)

type minuteDataRepo struct {
	dbConn    *sqlx.DB
	tableName string
}

func NewMinuteDataRepo(dbConn *sqlx.DB) *minuteDataRepo {
	return &minuteDataRepo{dbConn: dbConn, tableName: MINUTE_DATA_TABLE_NAME}
}

func (r *minuteDataRepo) GetAll(scripId string) ([]dao.Candle, error) {
	candles := make([]dao.Candle, 0)

	query := "SELECT open, high, low, close, volume, timestamp FROM minute_data where scrip_id = ?"
	result, err := r.dbConn.Queryx(query, scripId)
	if err != nil {
		fmt.Println(err.Error())
		return candles, err
	}

	for result.Next() {
		var candle dao.Candle
		scanErr := result.StructScan(&candle)
		if scanErr != nil {
			fmt.Println("Scan error", result)
		}
		candles = append(candles, candle)
	}
	return candles, nil
}
