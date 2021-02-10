package wifi

import (
	"testing"
)

func Test_func(t *testing.T) {
	if ssid, err := GetSSID(); err != nil {
		t.Error(err)
	} else {
		t.Log(ssid)

		if psd, err := GetPsw(ssid); err != nil {
			t.Error(err)
		} else {
			t.Log(psd)
		}
	}
}
