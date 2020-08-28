package main

import (
	"os"

	"github.com/tuotoo/qrcode"
)

func decode(picpath string) (ret string, err error) {
	var f *os.File
	f, err = os.Open(picpath)
	if err != nil {
		return
	}
	defer f.Close()

	var qrmatrix *qrcode.Matrix
	qrmatrix, err = qrcode.Decode(f)
	if err != nil {
		return
	}
	ret = qrmatrix.Content

	return
}
