package main

import (
	"io/ioutil"
)

func download(url, savename string) (err error) {
	var b []byte
	if b, err = httpGet(url); err != nil {
		return
	}

	return ioutil.WriteFile(savename, b, 0666)
}
