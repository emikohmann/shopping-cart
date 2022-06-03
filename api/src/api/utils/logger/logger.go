package logger

import "fmt"

func Info(msg string) {
	fmt.Println(fmt.Sprintf("INFO: %s", msg))
}

func Warn(msg string) {
	fmt.Println(fmt.Sprintf("WARN: %s", msg))
}

func Error(msg string, err error) {
	fmt.Println(fmt.Sprintf("ERR: %s: %v", msg, err))
}

func Panic(msg string, err error) {
	fmt.Println(fmt.Sprintf("PANIC: %s: %v", msg, err))
	panic(err)
}
