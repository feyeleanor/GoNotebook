package main
import . "fmt"

const (
  ROWS = 60
  COLS = 240
  MAX_WORK = 99
)

func main() {
  var row float32
  for ; row < ROWS; row++ {
    var buffer string
    var col float32
    for ; col < COLS; col++ {
      var x, y float32
      var i int

      for ; x * x + y * y <= 2 && i < MAX_WORK; i++ {
        x, y = x * x - y * y + 2 * col / COLS - 1.5, 2 * x * y + 2 * row / ROWS - 1
      }
      if i == MAX_WORK {
        buffer += "O"
      } else {
        buffer += " "
      }
    }
  Println(buffer)
  }
}