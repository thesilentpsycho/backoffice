package strategies

import (
	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
	"go.uber.org/zap"
)

type Strategy interface {
	Run(series *techan.TimeSeries)
	GetTimeOut()
}

type SimpleStrategy struct {
}

func (s *SimpleStrategy) Run(series *techan.TimeSeries) {
	record := techan.NewTradingRecord()

	positionSize := 20
	closePrices := techan.NewClosePriceIndicator(series)
	emaIndicator := techan.NewEMAIndicator(closePrices, 10)

	// Is satisfied when the price ema moves above 30 and the current position is new
	entryRule := techan.And(techan.NewCrossUpIndicatorRule(closePrices, emaIndicator),
		techan.PositionNewRule{})
	// Is satisfied when the price ema moves below 10 and the current position is open
	exitRule := techan.And(
		techan.NewCrossDownIndicatorRule(emaIndicator, closePrices),
		techan.PositionOpenRule{})

	strategy := techan.RuleStrategy{
		UnstablePeriod: 25,
		EntryRule:      entryRule,
		ExitRule:       exitRule,
	}

	for i := 0; i <= series.LastIndex(); i++ {
		if strategy.ShouldEnter(i, record) {
			record.Operate(techan.Order{
				Side:          techan.BUY,
				Security:      "",
				Price:         series.Candles[i].ClosePrice,
				Amount:        big.NewDecimal(float64(positionSize)),
				ExecutionTime: series.Candles[i].Period.End,
			})
			zap.L().Info("Entering at: ", zap.Any("candle", series.Candles[i]),
				zap.Float64("current EMA", emaIndicator.Calculate(i).Float()),
				zap.Float64("current Price", closePrices.Calculate(i).Float()),
			)
		} else if strategy.ShouldExit(i, record) {
			record.Operate(techan.Order{
				Side:          techan.SELL,
				Security:      "",
				Price:         series.Candles[i].ClosePrice,
				Amount:        big.NewDecimal(float64(positionSize)),
				ExecutionTime: series.Candles[i].Period.End,
			})
			zap.L().Info("Exiting at: ", zap.Any("candle", series.Candles[i]),
				zap.Float64("current EMA", emaIndicator.Calculate(i).Float()),
				zap.Float64("current Price", closePrices.Calculate(i).Float()),
			)
		}
	}

	totalProfit := techan.TotalProfitAnalysis{}.Analyze(record)
	zap.L().Info("Calculating Profits: ", zap.Float64("Total Profit", totalProfit))
	zap.L().Info("Done")
}

func (s SimpleStrategy) GetTimeOut() {
	panic("implement me")
}
