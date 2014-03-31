package main
import (
	. "fmt"
	"os"
)

func main() {
	Printf("hello world, %v\n", os.Args[1:])
}