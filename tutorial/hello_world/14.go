package main
import . "fmt"

func main() {
	greet("world")
}

func greet(name string) {
	Println("hello", name)
}