// registry注册表相关操作
package reg

import (
	"errors"
	"strconv"
	"syscall"

	"github.com/xiongyejun/xyjgo/winAPI/advapi32"
	"github.com/xiongyejun/xyjgo/winAPI/win"
)

type Reg struct {
	HKey   advapi32.HKEY
	Path   string
	Subs   []*Reg
	Parent *Reg
}

//REG_SZ：字符串：文本字符串
//REG_MULTI_SZ：多字符串值：含有多个文本值的字符串
//REG_BINARY：二进制数：二进制值，以十六进制显示，
//REG_DWORD：双字值；一个32位的二进制值，显示为8位的十六进制值。

func New(hkey advapi32.HKEY, path string) (r *Reg) {
	r = new(Reg)
	r.HKey = hkey
	r.Path = path
	return
}

func CreateKey(hkey advapi32.HKEY) (r *Reg) {
	return nil
}

func AddSub(hKey advapi32.HKEY, path string) (sub *Reg, err error) {
	return
}

func DelSub(path string) (err error) { return }

type RegValue struct {
	ValueName string
	Data      []byte
	Type      advapi32.REG_TYPE
}

func (me *Reg) EnumValue() (ret []*RegValue, err error) {
	var i int32 = 0
	var icount uint32 = 0

	var phkResult uint32
	if phkResult, err = me.OpenKey(); err != nil {
		return
	}
	defer advapi32.RegCloseKey(phkResult)

	for i == 0 {
		var lpType uint32
		var lpcbData uint32 = 512
		lpData := make([]byte, lpcbData)

		var lpcchValueName uint32 = lpcbData / 2
		lpValueName := make([]uint16, lpcchValueName)

		i = advapi32.RegEnumValue(advapi32.HKEY(phkResult), icount, &lpValueName[0], &lpcchValueName, nil, &lpType, &lpData[0], &lpcbData)
		if i > 0 {
			return ret, errors.New(win.GetErrString(int(i)))
		}
		icount++
		strlpValueName := syscall.UTF16ToString(lpValueName[:lpcchValueName])

		ret = append(ret, &RegValue{strlpValueName, lpData[:lpcbData], advapi32.REG_TYPE(lpType)})
	}
	return
}

func (me *Reg) ReadData(valueName string) (ret string, err error) {
	var phkResult uint32
	if phkResult, err = me.OpenKey(); err != nil {
		return
	}
	defer advapi32.RegCloseKey(phkResult)

	var lpType uint32
	var lpcbData uint32 = 512
	lpData := make([]byte, lpcbData)
	retv := advapi32.RegQueryValueEx(advapi32.HKEY(phkResult), valueName, nil, &lpType, &lpData[0], &lpcbData)
	if retv != 0 {
		return "", errors.New("RegQueryValueEx errno=" + strconv.Itoa(int(retv)))
	}
	lpData = lpData[:lpcbData]

	return string(lpData), nil
}

func (me *Reg) OpenKey() (phkResult uint32, err error) {
	ret := advapi32.RegOpenKeyEx(me.HKey, me.Path, 0, advapi32.KEY_READ, &phkResult)
	if ret != 0 {
		return 0, errors.New("RegOpenKeyEx errno=" + strconv.Itoa(int(ret)))
	}
	return
}
