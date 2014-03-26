package main
import . "fmt"

const Hello = "hello"
var world = "world"

func main() {
	world += "!"
	Println(Hello, world)
}