package services

import (
	"github.com/sdcoffey/techan"
	"papertrader.io/backoffice/strategies"
)

type StrategyExecutor interface {
	Execute(strategy strategies.Strategy, series *techan.TimeSeries)
}

type SimpleExecutor struct {
}

func (s *SimpleExecutor) Execute(strategy strategies.Strategy, series *techan.TimeSeries) {
	strategy.Run(series)
}
