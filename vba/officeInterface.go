package main

import (
	"archive/zip"
	"errors"
	"io"
	"io/ioutil"
	"os"

	"github.com/xiongyejun/xyjgo/vbaProject"
)

// 读取、改写office文件
type iOffice interface {
	readFile(fileName string) (b []byte, err error)
	reWriteFile(oldFileName string, saveFileName string, newByte []byte) error
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
	b, err = ioutil.ReadFile(fileName)
	return
}
func (me *file03) reWriteFile(oldFileName string, saveFileName string, newByte []byte) (err error) {
	err = ioutil.WriteFile(saveFileName, newByte, 0666)
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

func (me *file07) reWriteFile(oldFileName string, saveFileName string, newByte []byte) (err error) {
	var zipReader *zip.ReadCloser
	// 读取zip文件
	if zipReader, err = zip.OpenReader(oldFileName); err != nil {
		return
	}
	defer zipReader.Close()

	// 创建新文件
	var fw *os.File
	if fw, err = os.OpenFile(saveFileName, os.O_WRONLY|os.O_CREATE, 0666); err != nil {
		return
	}
	defer fw.Close()
	// 创建zip writer
	zipWriter := zip.NewWriter(fw)
	defer zipWriter.Close()
	// 循环zip文件中的文件
	for _, f := range zipReader.File {

		var b []byte
		// 如果是vba，就用改写了的流
		if f.Name == "xl/vbaProject.bin" {
			b = newByte
		} else {
			// 打开子文件
			var fr io.ReadCloser
			if fr, err = f.Open(); err != nil {
				return
			}
			defer fr.Close()
			// 读取子文件流
			if b, err = ioutil.ReadAll(fr); err != nil {
				return
			}
		}
		// 在zipwriter中创建新文件
		var wr io.Writer
		if wr, err = zipWriter.Create(f.Name); err != nil {
			return
		}
		// 写入新文件的数据
		n := 0
		if n, err = wr.Write(b); err != nil {
			return
		}
		if n < len(b) {
			return errors.New("写入不完整")
		}
	}
	err = zipWriter.Flush()

	return
}
