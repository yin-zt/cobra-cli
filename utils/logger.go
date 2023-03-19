package utils

import (
	log "github.com/cihub/seelog"
	"github.com/yin-zt/cobra-cli/config"
	"os"
)

var (
	Logger   log.LoggerInterface
	response any
)

func GetLog() log.LoggerInterface {
	os.MkdirAll("/var/log/", 0777)
	os.MkdirAll("/var/lib/cli", 0777)

	logger, err := log.LoggerFromConfigAsBytes([]byte(config.LogConfigStr))

	if err != nil {
		log.Error(err)
		response = "init log fail"
		panic(response)
	}
	Logger = logger
	return Logger
}
