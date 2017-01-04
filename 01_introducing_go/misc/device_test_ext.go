package main

/*
#cgo CFLAGS: -framework opencl
#cgo LDFLAGS: -framework opencl

#include "OpenCL/cl.h"
*/

import "C"

import "fmt"

func main() {
  var platform  C.cl_platform_id

  if err := clGetPlatformIDs(1, &platform, NULL); err < 0 {
    fmt.Errorf("Couldn't find any platforms")
  }
}