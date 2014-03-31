package main
import . "fmt"

func main() {
	Println(message())
}

func message() (string, string) {
	return "hello", "world"
}