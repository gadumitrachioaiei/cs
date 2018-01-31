package debug

import (
	"fmt"
	"log"
	"unsafe"
)

type Representation int

// types of bytes representation
const (
	Hexa = iota
	Binary
)

// ShowBytes prints the bytes in the memory pointed to by x
func ShowBytes(x interface{}, size int, representation Representation) {
	var p unsafe.Pointer
	switch y := x.(type) {
	case *int:
		p = unsafe.Pointer(y)
	case *int64:
		p = unsafe.Pointer(y)
	case *float64:
		p = unsafe.Pointer(y)
	case *int32:
		p = unsafe.Pointer(y)
	case *int16:
		p = unsafe.Pointer(y)
	case *int8:
		p = unsafe.Pointer(y)
	default:
		log.Fatalf("Unknown type: %T", y)
	}
	for i := 0; i < size; i++ {
		switch representation {
		case Hexa:
			fmt.Printf("%x ", *(*byte)(unsafe.Pointer((uintptr(p) + uintptr(i)))))
		case Binary:
			fmt.Printf("%b ", *(*byte)(unsafe.Pointer((uintptr(p) + uintptr(i)))))
		}

	}
	fmt.Println()
}
