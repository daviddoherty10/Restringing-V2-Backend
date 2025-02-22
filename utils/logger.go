package utils

import (
	"fmt"
	"log"
	"os"
)

func LoggerStartup() {
	logFile, err := os.OpenFile("../app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error on Logging Startup: " + err.Error())
		panic(1)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
}
