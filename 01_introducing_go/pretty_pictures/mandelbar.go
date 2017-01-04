package main
import . "fmt"

/*
#60.times{ |a|
#	puts((0..240).map{ |b|
#		x = y = i = 0
#		until  x * x + y * y > 4 || i == 99
#			x, y, i = x * x - y * y + b / 120.0 - 1.5, 2 * x * y + a / 30.0 - 1, i + 1
#		end
#		i == 99 ? '#' : '.'
#	} * '')
#}
*/

const (
	HEIGHT = 60
	WIDTH = 240
	ITERATIONS = 99
)

func main() {
	for a := 0; a < HEIGHT; a++ {
		var buffer	string
		for p := 0; p < 240; p++ {
			var x, y	float32
			var i		int

			for ; x * x + y * y <= 4 && i < ITERATIONS; i++ {
				 x, y = x * x - y * y + (2 * float32(p) / WIDTH) - 1.5, 2 * x * y + (2 * float32(a) / HEIGHT) - 1
			}
			if i == ITERATIONS {
				buffer += "#"
			} else {
				buffer += " "
			}
		}
		Println(buffer)
	}
}