package main

import (
	"testing"
)

func Test_func(t *testing.T) {
	d.strDir = "."

	if err := d.getFiels(); err != nil {
		t.Error(err)
		return
	}

	for i := range d.files {
		if err := d.getResult(i); err != nil {
			t.Error(err)
		}
	}
}
