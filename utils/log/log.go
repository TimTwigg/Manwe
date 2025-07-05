package logger

import "fmt"

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var blue = "\033[34m"
var magenta = "\033[35m"
var cyan = "\033[36m"
var gray = "\033[37m"
var orange = "\033[38;5;208m"

func Init(message ...any) {
	fmt.Printf("%s[INIT]%s", blue, reset)
	for _, m := range message {
		fmt.Printf(" %v", m)
	}
	fmt.Println()
}

func Info(message ...any) {
	fmt.Printf("%s[INFO]%s", green, reset)
	for _, m := range message {
		fmt.Printf(" %v", m)
	}
	fmt.Println()
}

func Warn(message ...any) {
	fmt.Printf("%s[WARNING]%s", orange, reset)
	for _, m := range message {
		fmt.Printf(" %v", m)
	}
	fmt.Println()
}

func Error(message ...any) {
	fmt.Printf("%s[ERROR]%s", red, reset)
	for _, m := range message {
		fmt.Printf(" %v", m)
	}
	fmt.Println()
}

func GetRequest(message ...any) {
	fmt.Printf("%s[GET]%s", cyan, reset)
	for _, m := range message {
		fmt.Printf(" %v", m)
	}
	fmt.Println()
}

func PostRequest(message ...any) {
	fmt.Printf("%s[POST]%s", magenta, reset)
	for _, m := range message {
		fmt.Printf(" %v", m)
	}
	fmt.Println()
}

func OptionsRequest(message ...any) {
	fmt.Printf("%s[OPTIONS]%s", gray, reset)
	for _, m := range message {
		fmt.Printf(" %v", m)
	}
	fmt.Println()
}

func DeleteRequest(message ...any) {
	fmt.Printf("%s[DELETE]%s", yellow, reset)
	for _, m := range message {
		fmt.Printf(" %v", m)
	}
	fmt.Println()
}
