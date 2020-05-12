package main

import (
	"flag"
	"go.uber.org/zap"
	"papertrader.io/backoffice/config"
	"papertrader.io/backoffice/infra"
	"papertrader.io/backoffice/logging"
	"papertrader.io/backoffice/repository"
)

func main() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "config.yml", "absolute path to the configuration file")
	flag.Parse()
	configuration := config.Loadconfig(configFilePath)

	logger := logging.InitLogger(configuration)
	defer logger.Sync()

	undo := zap.ReplaceGlobals(logger)
	defer undo()
	dbConn := infra.DbConn(configuration)
	minuteDataRepo := repository.NewMinuteDataRepo(dbConn)

	candles, err := minuteDataRepo.GetAll("2952193")
	if err != nil {
		return
	}
	zap.L().Info("Number of candles received : " , zap.Int("count", len(candles)))
	return
}
