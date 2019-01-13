package main

import (
	"testing"
)

func Test_func(t *testing.T) {
	err := bmpTOjpeg("1.bmp")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("ok")
}
