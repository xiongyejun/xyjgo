package pressKey

import (
	"testing"
	"time"
)

func Test_func(t *testing.T) {
	k := New()
	k.Add('A', 3000, 0)
	k.Add('T', 3000, 0)

	go func() {
		time.Sleep(10 * time.Second)
		k.Stop()
	}()

	k.Start()
}
