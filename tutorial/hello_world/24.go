package main
import (
	"flag"
	. "fmt"
	. "strings"
)

var name, spacer *string
var message	string

func init() {
	name = flag.String("n", "world", "n: name of person to greet")
	spacer = flag.String("s", ",", "s: separator between name and message")
	flag.Parse()
	message = Join(flag.Args(), " ")
}

func main() {
	if len(message) > 0 {
		Printf("hello %v%v %v\n", *name, *spacer, message)
	} else {
		Println("hello", *name)
	}
}