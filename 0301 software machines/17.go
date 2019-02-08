package main
import "fmt"

func main() {
  m := make(map[string] interface{})
  m["apple"] = "rosy"
  fmt.Printf("%v\n", m["apple"])

  m["blueberry"] = "sweet"
  fmt.Printf("%v\n", m["blueberry"])

  m["cherry"] = "pie"
  fmt.Printf("%v\n", m["cherry"])

  m["cherry"] = "tart"
  fmt.Printf("%v\n", m["cherry"])

  fmt.Printf("%v\n", m["tart"])
}
