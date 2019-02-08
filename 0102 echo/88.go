package main
import "flag"
import . "fmt"
import "os"
import . "strings"

leanpub-start-insert
const VERSION = "0.0.12"
leanpub-end-insert

var suppress_newline bool
leanpub-start-insert
var display_version_info bool
leanpub-end-insert

func init() {
  flag.BoolVar(&suppress_newline, "n", false, "n: suppress printing of newline")
leanpub-start-insert
  flag.BoolVar(&display_version_info, "version", false, "version: display version information")
leanpub-end-insert
}

func main() {
  flag.Parse()
leanpub-start-insert
  if display_version_info {
    Println(VERSION)
    os.Exit(0)
  }
leanpub-end-insert
  Print(Join(flag.Args(), " "))
  if !suppress_newline {
    Println()
  }
}
