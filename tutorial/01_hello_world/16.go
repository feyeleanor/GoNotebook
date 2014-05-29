package main
import . "fmt"

func main() {
	print("Hello", "world")
}

func print(v ...interface{}) {
	Println(v...)
}