package main
import (
	"os"
	. "flag"
	"fmt"
	"io/ioutil"
	"strconv"
	. "strings"
)

var name, spacer *string
var message, text_block	string
var repeats int

func init() {
	def_n := os.Getenv("DEF_NAME")
	if len(def_n) == 0 {
		def_n = "world"
	}

	def_c, _ := strconv.Atoi(os.Getenv("DEF_REPS"))

	name = String("n", def_n, "n: name of person to greet")
	spacer = String("s", ",", "s: separator between name and message")
	file := String("f", "", "f: name of a file containing a block of text to display")
	IntVar(&repeats, "c", def_c, "c: number of times to display the message")
	Parse()
	if text, err := ioutil.ReadFile(*file); err == nil {
		text_block = string(text)
	}
	message = Join(Args(), " ")
}

func main() {
	switch {
	case len(message) > 0:
		message = fmt.Sprintf("hello %v%v %v", *name, *spacer, message)
		if len(text_block) > 0 {
			message += "\n" + text_block
		}
	case len(text_block) > 0:
		message = fmt.Sprintf("hello %v\n%v", *name, text_block)
	default:
		message = fmt.Sprintf("hello %v", *name)
	}

	for ; repeats > 0; repeats-- {
		fmt.Println(message)
	}
}