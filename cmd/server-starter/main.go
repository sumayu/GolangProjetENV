// main.go
package main

import (
	"log"
	"main/backend/backendCFG"
	"main/envchecker"
	"main/logger" // Импортируйте пакет logger
	"main/myDB"
	"main/router"
	"net/http"
)

func main() {
	pathToYaml := "C:/gol/configYML/config.yaml"
	a, err := configs.LoadConfigs(pathToYaml)
	

	// Инициализируйте логгер
	logger.InitLogger(envchecker.Envchecker())

	if err == nil {
		logger.Logger.Debug("configs -> starter(main.go) | TRUE |")
	} else {
		msg := "configs -> starter(main.go) | ERROR |"
		log.Fatalf(msg, a, err)
	}
	log.Println("server start on", envchecker.Envchecker(), "env")
	
mydb.Database()	



















	errorSERVER := http.ListenAndServe(":8080", router.Router())
if err != nil {
	logger.Logger.Info("error with create server API (router)", errorSERVER)
}
}

