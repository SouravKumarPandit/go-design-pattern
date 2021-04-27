package main

import "fmt"

var errorLogger = NewErrorLogger(ERROR)
var fileLogger = NewFileLogger(DEBUG)
var consoleLogger = NewConsoleLogger(INFO)

func main() {
	fmt.Println("Logger ")

	abstractLogger:=getChainOfLoggers()

	abstractLogger.logMessage(INFO,"This Is an INFO Information")
	abstractLogger.logMessage(DEBUG,"This Is an DEBUG Level Information")
	abstractLogger.logMessage(ERROR,"This Is an ERROR Information")


}

func getChainOfLoggers() Logger {
	errorLogger.setNextLogger(&fileLogger)
	fileLogger.setNextLogger(&consoleLogger)
	return &errorLogger
}
