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

func Init(message any) {
	fmt.Printf("%s[INIT]%s %v\n", blue, reset, message)
}

func Info(message any) {
	fmt.Printf("%s[INFO]%s %v\n", green, reset, message)
}

func Warn(message any) {
	fmt.Printf("%s[WARNING]%s %v\n", orange, reset, message)
}

func Error(message any) {
	fmt.Printf("%s[ERROR]%s %v\n", red, reset, message)
}

func GetRequest(message any) {
	fmt.Printf("%s[GET]%s %v\n", cyan, reset, message)
}

func PostRequest(message any) {
	fmt.Printf("%s[POST]%s %v\n", magenta, reset, message)
}

func OptionsRequest(message any) {
	fmt.Printf("%s[OPTIONS]%s %v\n", gray, reset, message)
}

func DeleteRequest(message any) {
	fmt.Printf("%s[DELETE]%s %v\n", yellow, reset, message)
}
