package logging

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "[RGOAUTH] ", log.LstdFlags|log.Lshortfile)
}

func Info(message string) {
	logger.Println("[INFO]", message)
}

func Error(message string, err error) {
	logger.Println("[ERROR]", message, err)
}

func DB(message string, err error) {
	logger.Println("[DB]", message, err)
}
