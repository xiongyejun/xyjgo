package keyboard

import (
	"testing"
	"time"
)

func Test_press(t *testing.T) {

	for i := 0; i < 5; i++ {
		Press(VK_LEFT, time.Second)
	}

}
