package unsafe

import (
	"fmt"
	"unsafe"
)

var x struct {
	a bool
	b int16
	c []int
}

func PrintUnsafe() {
	fmt.Printf("Sizeof(x) = %v Alignof(x) = %v\n", unsafe.Sizeof(x), unsafe.Alignof(x))
	fmt.Printf("Sizeof(x.a) = %v Alignof(x.a) = %v Offsetof(x.a) = %v\n", unsafe.Sizeof(x.a), unsafe.Alignof(x.a), unsafe.Offsetof(x.a))
	fmt.Printf("Sizeof(x.b) = %v Alignof(x.b) = %v Offsetof(x.b) = %v\n", unsafe.Sizeof(x.b), unsafe.Alignof(x.b), unsafe.Offsetof(x.b))
	fmt.Printf("Sizeof(x.c) = %v Alignof(x.c) = %v Offsetof(x.c) = %v\n", unsafe.Sizeof(x.c), unsafe.Alignof(x.c), unsafe.Offsetof(x.c))
}

func UpdateStruct() {
	fmt.Printf("x.b = %v\n", x.b)
	var xb = (*int16)(unsafe.Pointer(
		uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*xb = 42
	fmt.Printf("x.b = %v\n", x.b)
}

func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}
