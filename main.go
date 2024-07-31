package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lmath
#include "math.h"
*/
import "C"
import (
	"fmt"
)

// Add wraps the C add function
func Add(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}

func main() {
	result := Add(5, 3)
	fmt.Printf("Result of 5 + 3 = %dn", result)
}
