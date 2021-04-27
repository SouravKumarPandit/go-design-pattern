package main

import "fmt"

const (
	INFO = iota
	DEBUG
	ERROR
)

type Logger interface {
	logMessage(level int,message string)
}

type AbstractLogger struct {
	INFO  int
	DEBUG int
	ERROR int
	level int
	Level      int
	NextLogger Logger
	write func(message string)
}


func (a *AbstractLogger) logMessage(level int,message string)  {
	if a.level<=level {
		a.write(message)
	}
	if a.NextLogger!=nil {
		a.NextLogger.logMessage(level,message)
	}

}
func (a *AbstractLogger) setNextLogger(logger Logger)  {
	a.NextLogger=logger
}


type ConsoleLogger struct {
	AbstractLogger
}

func NewConsoleLogger(level int) ConsoleLogger {
	return ConsoleLogger{
		AbstractLogger{
			Level:      level,
			write: func(message string) {
				fmt.Println("Standard Console::Logger:: "+message)
			},
		},
	}

}
type ErrorLogger struct {
	AbstractLogger
}

func NewErrorLogger(level int) ErrorLogger {
	return ErrorLogger{
		AbstractLogger{
			Level:      level,
			write: func(message string) {
				fmt.Println("Error Console::Logger: : "+message)
			},
		},
	}

}


type FileLogger struct {
	AbstractLogger
}

func NewFileLogger(level int) FileLogger {
	return FileLogger{
		AbstractLogger{
			Level:      level,
			write: func(message string) {
				fmt.Println("File :: Logger : "+message)
			},
		},
	}

}






































