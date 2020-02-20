package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	epub "github.com/bmaupin/go-epub"
)

func (me *EpubSet) create() (err error) {
	ep := epub.NewEpub(me.Name)
	// 封面图片
	if me.CoverPicFile != "" {
		var picpath string
		if picpath, err = ep.AddImage(me.CoverPicFile, me.CoverPicFile); err != nil {
			return
		}
		ep.SetCover(picpath, "")
	}

	var entrys []os.FileInfo
	if entrys, err = ioutil.ReadDir(me.SrcFolderPath); err != nil {
		return
	}
	for i := range entrys {
		if !entrys[i].IsDir() {
			var b []byte
			if b, err = ioutil.ReadFile(me.SrcFolderPath + entrys[i].Name()); err != nil {
				return
			}
			// 只需要提取小说内容
			var str string = string(b)
			if me.DivID != "" {
				if str, err = getHtmlFromDivId(str, me.DivID); err != nil {
					return
				}
			} else {
				for j := range me.SplitInfos {
					str = strings.Split(str, me.SplitInfos[j].Sep)[me.SplitInfos[j].RetIndex]
				}
			}

			for j := range me.ReplaceExpr {
				var reg *regexp.Regexp
				if reg, err = regexp.Compile(me.ReplaceExpr[j]); err != nil {
					return
				}
				str = reg.ReplaceAllString(str, "")
			}

			// 为了创建epub格式
			str = `<h2>` + entrys[i].Name()[6:] + "</h2>" + str

			if _, err = ep.AddSection(str, entrys[i].Name()[6:], "", ""); err != nil {
				return
			}
			fmt.Println(i, "epub ok")
		}
	}

	ep.SetAuthor(me.Author)
	return ep.Write(me.Name + ".epub")
}
