package main
import . "fmt"

type Text string

func main() {
  var name Text = "Ellie"
  var pointer_to_name *Text

  pointer_to_name = &name
  Printf("name = %v stored at %v\n", name, pointer_to_name)
  Printf("pointer_to_name references %v\n", *pointer_to_name)
}
