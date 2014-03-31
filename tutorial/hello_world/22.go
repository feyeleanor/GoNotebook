package main
import (
	. "fmt"
	"os"
	"strings"
)

func main() {
	Printf("hello world, %v\n", strings.Join(os.Args[1:], " "))
}