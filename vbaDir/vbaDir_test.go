package vbaDir

import (
	"io/ioutil"
	"testing"
)

func Test_func(t *testing.T) {
	b, _ := ioutil.ReadFile(`../rleVBA/1.txt`)

	ret, err := GetModuleInfo(b)
	if err != nil {
		t.Error(err)
		return
	}

	for i := range ret {
		t.Logf("%d %s\n", i, ret[i].Name)
	}
}
