package main

import (
	"testing"
)

func Test_func(t *testing.T) {
	if ret, err := scanDir(`.`); err != nil {
		t.Error(err)
	} else {
		t.Log(ret)
	}
}
