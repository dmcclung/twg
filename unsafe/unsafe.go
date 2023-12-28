package unsafe

import (
	"fmt"
	"unsafe"
)

func PrintUnsafe() {
	x := struct {
		a bool
		b int16
		c []int
	}{}

	_ = x

	fmt.Printf("Sizeof(x) = %v Alignof(x) = %v\n", unsafe.Sizeof(x), unsafe.Alignof(x))
	fmt.Printf("Sizeof(x.a) = %v Alignof(x.a) = %v Offsetof(x.a) = %v\n", unsafe.Sizeof(x.a), unsafe.Alignof(x.a), unsafe.Offsetof(x.a))
	fmt.Printf("Sizeof(x.b) = %v Alignof(x.b) = %v Offsetof(x.b) = %v\n", unsafe.Sizeof(x.b), unsafe.Alignof(x.b), unsafe.Offsetof(x.b))
	fmt.Printf("Sizeof(x.c) = %v Alignof(x.c) = %v Offsetof(x.c) = %v\n", unsafe.Sizeof(x.c), unsafe.Alignof(x.c), unsafe.Offsetof(x.c))
}
