package main

import (
	"archive/zip"
	"errors"
	"io"
	"io/ioutil"

	"github.com/xiongyejun/xyjgo/vbaProject"
)

// 读取、改写office文件
type iOffice interface {
	readFile(fileName string) (b []byte, err error)
	reWriteFile() error
}

var iof iOffice
var of *officeFile

type officeFile struct {
	fileName string
	vp       *vbaProject.VBAProject

	b []byte
}
type file03 struct {
}

func (me *file03) readFile(fileName string) (b []byte, err error) {
	if b, err = ioutil.ReadFile(fileName); err != nil {
		return
	}
	return
}
func (me *file03) reWriteFile() (err error) {
	return
}

type file07 struct {
	vbaProjectIndex int
}

func (me *file07) readFile(fileName string) (b []byte, err error) {
	var reader *zip.ReadCloser
	if reader, err = zip.OpenReader(fileName); err != nil {
		return
	}
	defer reader.Close()

	for i, f := range reader.File {
		if f.Name == "xl/vbaProject.bin" {
			var rc io.ReadCloser
			if rc, err = f.Open(); err != nil {
				return
			}

			me.vbaProjectIndex = i
			if b, err = ioutil.ReadAll(rc); err != nil {
				return
			}

			return
		}
	}
	return nil, errors.New("err: 没有找到 vbaProject.bin")
}

func (me *file07) reWriteFile() (err error) {
	return
}
