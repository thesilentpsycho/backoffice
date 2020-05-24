package services

import (
	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
	"papertrader.io/backoffice/dao"
	"time"
)

func GetSeriesFromCandles(candles []dao.Candle) *techan.TimeSeries {
	series := techan.NewTimeSeries()

	for _, candle := range candles {
		techanCandle := techan.NewCandle(techan.NewTimePeriod(candle.TimeStamp, time.Minute*1))
		techanCandle.ClosePrice = big.NewDecimal(candle.Close)
		techanCandle.OpenPrice = big.NewDecimal(candle.Open)
		techanCandle.MaxPrice = big.NewDecimal(candle.High)
		techanCandle.MinPrice = big.NewDecimal(candle.Low)
		techanCandle.Volume = big.NewDecimal(float64(candle.Volume))
		series.AddCandle(techanCandle)
	}

	return series
}
