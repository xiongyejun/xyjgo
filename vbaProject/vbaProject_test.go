package vbaProject

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/axgle/mahonia"
)

func Test_getModuleInfo(t *testing.T) {
	vp := New()
	if b, err := ioutil.ReadFile("vbaProject.bin"); err != nil {
		fmt.Println(err)
	} else {
		if err := vp.Parse(b); err != nil {
			fmt.Println(err)
		} else {
			for i := range vp.Module {
				fmt.Println(vp.Module[i].Name)
				decoder := mahonia.NewDecoder("gbk")

				str := decoder.ConvertString(string(vp.Module[i].Code))

				fmt.Println(str)

			}
		}
	}

}
