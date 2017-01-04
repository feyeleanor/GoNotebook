package main
import . "fmt"
import "reflect"

func main() {
	v := reflect.ValueOf(map[int] int{0: 0, 1: 2, 2: 4, 3: 6, 4: 8})
	if v.Kind() == reflect.Map {
		for i, k := range v.MapKeys() {
			Printf("%v: %v\n", i, v.MapIndex(k).Interface())
			v.SetMapIndex(k, k)
		}
		for i, k := range v.MapKeys() {
			Printf("%v: %v\n", i, v.MapIndex(k).Interface())
		}
	}
}