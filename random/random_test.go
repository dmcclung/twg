package random

import (
	"math/rand"
	"testing"
)

func TestPick(t *testing.T) {
	var seed int64 = 1702101523435974748
	t.Logf("Seed: %d", seed)

	r := rand.New(rand.NewSource(seed))
	arg := make([]int, 5)
	for i := 0; i < 5; i++ {
		arg[i] = r.Int()
	}

	got := Pick(arg)

	for _, v := range arg {
		if got == v {
			return
		}
	}

	t.Errorf("Pick(seed=%d, %d) = %d; not in slice", seed, arg, got)
}
