package advapi32

import (
	"testing"
)

func Test_func(t *testing.T) {
	var phkResult uint32

	ret := RegOpenKeyEx(HKEY_CLASSES_ROOT, `.zip`, 0, KEY_READ, &phkResult)
	t.Log(ret, phkResult)

	var lpType uint32
	var lpcbData uint32 = 128
	lpData := make([]byte, lpcbData)
	ret = RegQueryValueEx(HKEY(phkResult), "Content Type", nil, &lpType, &lpData[0], &lpcbData)
	t.Log(ret, lpType, lpcbData)
	lpData = lpData[:lpcbData]
	t.Logf("%s\r\n", lpData)
	ret = RegCloseKey(phkResult)
	t.Log(ret)
}
