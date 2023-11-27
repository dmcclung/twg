package sleep

import (
	"fmt"
	"testing"
	"time"
)

func TestTmpExecutable(t *testing.T) {
	time.Sleep(time.Minute)
	fmt.Println("hello")
}
