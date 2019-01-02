package reg

import (
	"fmt"
	"testing"

	"github.com/xiongyejun/xyjgo/ucs2T0utf8"
	"github.com/xiongyejun/xyjgo/winAPI/advapi32"
)

func Test_func(t *testing.T) {
	r := New(advapi32.HKEY_CLASSES_ROOT, `.zip`)
	ret, err := r.ReadData("Content Type")
	t.Log(ret, err)

	rv, err := r.EnumValue()
	t.Log(rv, err)

	for i := range rv {
		if rv[i].Type == advapi32.REG_SZ {
			rv[i].Data, err = ucs2T0utf8.UCS2toUTF8(rv[i].Data)
			if err != nil {
				t.Log(err)
				return
			}
			fmt.Printf("%d %s %s, 0x%x \r\n", i, rv[i].ValueName, rv[i].Data, rv[i].Type)
		} else {
			fmt.Printf("%d %s %x, 0x%x \r\n", i, rv[i].ValueName, rv[i].Data, rv[i].Type)
		}

	}
	//	printhex(`%SystemRoot%\Explorer.exe`)
	//	printhex(`%GOPATH%\bin\sendMail.exe`)
}

// 注册表使用环境变量的时候，需要把类型设置为REG_EXPAND_SZ，但是regedit好像不能直接编辑
// 用.reg文件导入的时候，又需要使用hex16进制数据，用printhex输出，最后的逗号不需要
func printhex(str string) {
	b := []byte(str)
	for i := range b {
		fmt.Printf("%x,00,", b[i])
	}

	fmt.Println()
}
