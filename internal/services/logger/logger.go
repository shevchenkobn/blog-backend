package logger

import (
	"fmt"
	"os"
)

type Logger struct{}

func NewLogger() *Logger {
	return new(Logger)
}

func (l *Logger) Printf(format string, values ...interface{}) {
	fmt.Printf(format+"\n", values...)
}

func (l *Logger) Print(values ...interface{}) {
	fmt.Println(values...)
}

func (l *Logger) Errorf(format string, values ...interface{}) {
	_, err := fmt.Fprintf(os.Stderr, format+"\n", values...)
	if err != nil {
		printError(err)
	}
}

func (l *Logger) Error(values ...interface{}) {
	_, err := fmt.Fprint(os.Stderr, values...)
	if err != nil {
		printError(err)
	}
}

func printError(err error) {
	fmt.Printf("PRINT_ERROR: %s", err)
}
