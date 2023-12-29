package unsafe_test

import (
	"fmt"

	"github.com/dmcclung/twg/unsafe"
)

func ExamplePrintUnsafe() {
	unsafe.PrintUnsafe()

	// Output: Sizeof(x) = 32 Alignof(x) = 8
	// Sizeof(x.a) = 1 Alignof(x.a) = 1 Offsetof(x.a) = 0
	// Sizeof(x.b) = 2 Alignof(x.b) = 2 Offsetof(x.b) = 2
	// Sizeof(x.c) = 24 Alignof(x.c) = 8 Offsetof(x.c) = 8
}

func ExampleFloat64bits() {
	var x uint64 = unsafe.Float64bits(1.0)
	fmt.Printf("%#016x\n", x)

	// Output: 0x3ff0000000000000
}

func ExampleUpdateStruct() {
	unsafe.UpdateStruct()

	// Output: x.b = 0
	// x.b = 42
}
