package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

func LoggerStartup() {
	logFile, err := os.OpenFile("./app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error on Logging Startup:", err)
		log.Fatal(err)
	}

	// Set log output to both stdout and the log file
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	log.SetFlags(log.LstdFlags | log.Lshortfile) // Adds timestamps & file info

	log.Println("Logger initialized") // Test log to confirm it's working
}
