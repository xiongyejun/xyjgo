package sms

import (
	"testing"
)

func Test_func(t *testing.T) {
	ret, _ := GetMsg("155")
	for i := range ret {
		t.Log(ret[i].NO, ret[i].From, ret[i].Content, ret[i].Time)
	}
	// ret, err := GetInfos()

	// if err != nil {
	// 	t.Error(err)
	// } else {
	// 	if err = SaveToJson("json.txt", ret); err != nil {
	// 		t.Error(err)
	// 	}
	// }

	// b, _ := ioutil.ReadFile("1.html")

	// t.Logf("%s\n", b)
	// ret, _ := getInfo(b)

	// for i := range ret {
	// 	t.Log(ret[i].EM, ret[i].Phone, ret[i].SMSContent)
	// }
}
