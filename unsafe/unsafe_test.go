package unsafe_test

import (
	"github.com/dmcclung/twg/unsafe"
)

func ExamplePrintUnsafe() {
	unsafe.PrintUnsafe()

	// Output: Sizeof(x) = 32 Alignof(x) = 8
	// Sizeof(x.a) = 1 Alignof(x.a) = 1 Offsetof(x.a) = 0
	// Sizeof(x.b) = 2 Alignof(x.b) = 2 Offsetof(x.b) = 2
	// Sizeof(x.c) = 24 Alignof(x.c) = 8 Offsetof(x.c) = 8
}
