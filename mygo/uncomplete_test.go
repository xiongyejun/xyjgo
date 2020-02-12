package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func Test_func(t *testing.T) {
	var err error
	err = testEpub()

	if err != nil {
		t.Error(err)
	}
}

func testEpub() (err error) {
	if err = createEpub("/Users/xiongyejun/01-GitHub/08-go/src/github.com/xiongyejun/xyjgo/tmp/srcHtml/", "fan.epub", "凡人修仙传"); err != nil {
		return
	}
	return

}
func testDown() (err error) {
	var b []byte
	b, err = ioutil.ReadFile("dirjosn.txt")
	if err != nil {
		return
	}
	var di dirInfos = dirInfos{}
	if err = json.Unmarshal(b, &di); err != nil {
		return
	}
	di.downXS()
	<-ch
	<-ch
	return
}
func testGetDir() (err error) {
	var b []byte
	b, err = ioutil.ReadFile("dir.html")
	if err != nil {
		return
	}

	di, err := getDir(b)
	if err != nil {
		return
	}

	err = di.saveJsonTxt("dirJson.txt")
	if err != nil {
		return
	}
	return
}
