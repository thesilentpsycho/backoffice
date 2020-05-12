package main

import (
	"flag"
	"fmt"
	"papertrader.io/backoffice/config"
	"papertrader.io/backoffice/infra"
	"papertrader.io/backoffice/repository"
)

func main() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "config.yml", "absolute path to the configuration file")
	flag.Parse()
	configuration := config.Loadconfig(configFilePath)

	dbConn := infra.DbConn(configuration)
	minuteDataRepo := repository.NewMinuteDataRepo(dbConn)

	candles, err := minuteDataRepo.GetAll("2952193")
	if err != nil {
		return
	}
	fmt.Println(len(candles))
	return
}
