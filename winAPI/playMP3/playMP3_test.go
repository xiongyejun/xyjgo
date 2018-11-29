package playMP3

import (
	"testing"
)

func Test_play(t *testing.T) {
	defer Free()
	if err := Play(`C:\Users\Administrator\Desktop\voice.mp3`, true); err != nil {
		t.Error(err)
		return
	}
}
