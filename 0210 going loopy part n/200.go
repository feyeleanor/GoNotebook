package main

import (
  "fmt"
  "reflect"
//  "github.com/feyeleanor/raw"
//  "unsafe"
)

func duplicate(v reflect.Value) interface{} {
  rv := reflect.New(v.Type()).Elem()
  rv.Set(v)
  return rv.Interface()
}

/*
func makeAddressable(v reflect.Value) (r reflect.Value) {
  if r = v; !r.CanAddr() {
    ptr := reflect.PtrTo(v.Type())
    r = reflect.New(ptr)
fmt.Println(r)
fmt.Println(r.Elem().UnsafeAddr())
    r = reflect.NewAt(v.Type(), unsafe.Pointer(r.Elem().UnsafeAddr()))

//    reflect.Indirect(r.Elem()).Set(v)
  }
  return
}
*/

func CanAddress(v interface{}) interface{} {
  fmt.Printf("%v: %v\n", v, reflect.ValueOf(v).CanAddr())
  return v
}

var A struct { A, B, C int } = struct{ A, B, C int }{ 1, 3, 5}
var B *struct { A, B, C int }

var i interface{}

func main() {
  fmt.Printf("%v: %v\n", A, reflect.ValueOf(A).CanAddr())
  fmt.Printf("%v: %v\n", &A, reflect.ValueOf(&A).CanAddr())

/*
  A = struct{ A, B, C int }{ 3, 5, 7}
  CanAddress(A)
  C := A
  CanAddress(C)

  B = &struct{ A, B, C int }{ 4, 6, 8}
  CanAddress(B)

  CanAddress(raw.NewPoint(1, 3, 5))
  CanAddress(*raw.NewPoint(2, 4, 6))

  x := raw.NewPoint(3, 5, 7)
  CanAddress(x)
  CanAddress(&x)
*/
//  x := *raw.NewPoint(3, 5, 7)
//  y := raw.NewPoint(2, 4, 6)

//  rx := reflect.ValueOf(x)
//  ry := reflect.ValueOf(y)

//  z := duplicate(rx)
//  fmt.Printf("%#v =?= %#v: %v\n", x, z, x == z)

//  z = duplicate(ry)
//  fmt.Printf("%#v =?= %#v: %v\n", y, z, y == z)

//  rx.Set(makeAddressable(reflect.ValueOf(y)))
//  fmt.Printf("%#v =?= %#v: %v\n", x, y, x == y)
}
