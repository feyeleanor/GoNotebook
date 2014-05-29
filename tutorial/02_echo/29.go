package main
import (
	"os"
	. "flag"
	. "fmt"
	"strconv"
	. "strings"
)

var name, spacer *string
var message	string
var repeats int

func init() {
	def_n := os.Getenv("DEF_NAME")
	if len(def_n) == 0 {
		def_n = "world"
	}

	def_c, e := strconv.Atoi(os.Getenv("DEF_REPS"))
	if e != nil {
		def_c = 0
	}

	name = String("n", def_n, "n: name of person to greet")
	spacer = String("s", ",", "s: separator between name and message")
	IntVar(&repeats, "c", def_c, "c: number of times to display the message")
	Parse()
	message = Join(Args(), " ")
}

func main() {
	if len(message) > 0 {
		for ; repeats > 0; repeats-- {
			Printf("hello %v%v %v\n", *name, *spacer, message)
		}
	} else {
		for ; repeats > 0; repeats-- {
			Println("hello", *name)
		}
	}
}