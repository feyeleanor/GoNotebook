package main
import (
	"os"
	. "flag"
	"fmt"
	"io/ioutil"
	"strconv"
	. "strings"
)

const (
	OK = iota
	NO_SUCH_FILE
	INVALID_REPS
)

type OSError interface {
	fmt.Stringer
	Code() int
}

type NoSuchFile	string
func (n NoSuchFile) String() string {
	return "no such file: " + string(n)
}

func (n NoSuchFile) Code() int {
	return NO_SUCH_FILE
}

type InvalidReps string
func (i InvalidReps) String() string {
	return "Reps must be an integer: " + string(i)
}

func (i InvalidReps) Code() int {
	return INVALID_REPS
}

var name, spacer *string
var message, text_block	string
var repeats int

func init() {
	var file string

	defer abortOnError()

	name = String("n", defaultName(), "n: name of person to greet")
	spacer = String("s", ",", "s: separator between name and message")
	StringVar(&file, "f", "", "f: name of a file containing a block of text to display")
	IntVar(&repeats, "c", defaultRepeats(), "c: number of times to display the message")
	Parse()
	text_block = loadMessage(file)
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

func defaultName() (r string) {
	if r = os.Getenv("DEF_NAME"); len(r) == 0 {
		r = "world"
	}
	return
}

func defaultRepeats() (r int) {
	if d := os.Getenv("DEF_REPS"); len(d) != 0 {
		var e error
		if r, e = strconv.Atoi(d); e != nil {
			panic(InvalidReps(d))
		}
	}
	return
}

func loadMessage(filename string) (r string) {
	if len(filename) > 0 {
		if text, err := ioutil.ReadFile(filename); err == nil {
			r = string(text)
		} else {
			panic(NoSuchFile(filename))
		}
	}
	return
}

func abortOnError() {
	switch e := recover().(type) {
	case OSError:
		fmt.Println(e)
		os.Exit(e.Code())
	}
}