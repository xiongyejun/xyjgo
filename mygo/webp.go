package main

import (
	"image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
)

func webp2jpgDir(dir string) (err error) {
	println(dir)
	return filepath.Walk(dir, walkfunc)
}

func walkfunc(path string, info os.FileInfo, err error) error {
	if strings.HasSuffix(strings.ToLower(path), ".webp") {
		if err = webp2jpg(path); err != nil {
			return err
		}
	}
	return nil
}

func webp2jpg(file string) (err error) {
	println(file)
	var b []byte
	if b, err = ioutil.ReadFile(file); err != nil {
		return
	}
	var m *webp.RGBImage
	if m, err = webp.DecodeRGB(b); err != nil {
		return
	}

	var w *os.File
	if w, err = os.Create(file + ".jpg"); err != nil {
		return
	}
	if err = jpeg.Encode(w, m, &jpeg.Options{100}); err != nil {
		return
	}

	return
}
