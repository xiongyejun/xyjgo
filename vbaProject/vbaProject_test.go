package vbaProject

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_GetModuleCode(t *testing.T) {
	vp := New()
	if b, err := ioutil.ReadFile("vbaProject.bin"); err != nil {
		fmt.Println(err)
	} else {
		if err := vp.Parse(b); err != nil {
			fmt.Println(err)
		} else {
			if str, err := vp.GetModuleCode("模块1"); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(str)
			}
		}
	}

}
