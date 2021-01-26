package main

import (
	"log"
	"os"
)

var (
	Logger *log.Logger
)

func mylogger() {
	// set location of log file
	// To write logs in file
	var errorLogpath = "error.log"
	var errorFile, err = os.OpenFile(errorLogpath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println("Error", err)
	}
	Logger = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	mylogger()
	Logger.Print("This is to verify my message date and time.")
}
