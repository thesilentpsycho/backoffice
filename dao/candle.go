package dao

import "time"

type Candle struct {
	TimeStamp time.Time `db:"timestamp"`
	Open      float64   `db:"open"`
	High      float64   `db:"high"`
	Low       float64   `db:"low"`
	Close     float64   `db:"close"`
	Volume    int64     `db:"volume"`
}
