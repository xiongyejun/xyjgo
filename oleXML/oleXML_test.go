package oleXML

import (
	"io/ioutil"
	"testing"
)

func Test_func(t *testing.T) {
	x := New()
	err := x.InitXML()

	if err != nil {
		t.Error(err)
		return
	}
	defer x.UnInit()

	b, _ := ioutil.ReadFile(`C:\Users\Administrator\Desktop\customUI.xml`)

	err = x.Validate(b, `E:\04-github\08-go\src\github.com\xiongyejun\xyjgo\ribbon\CustomUI.xsd`, `http://schemas.microsoft.com/office/2006/01/customui`)

	if err != nil {
		t.Error(err)
		return
	}

}
