package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	epub "github.com/bmaupin/go-epub"
)

func createEpub(folder, savePath, title string) (err error) {
	var divID string = "content"
	// <div id="content">
	divID = `<div id="` + divID + `">`

	ep := epub.NewEpub(title)
	if _, err = ep.AddCSS("/Users/xiongyejun/01-GitHub/08-go/src/github.com/bmaupin/go-epub/testdata/cover.css", "epub.css"); err != nil {
		return
	}
	var entrys []os.FileInfo
	if entrys, err = ioutil.ReadDir(folder); err != nil {
		return
	}
	for i := range entrys {
		// if i == 2 {
		// 	break
		// }
		if !entrys[i].IsDir() {
			var b []byte
			if b, err = ioutil.ReadFile(folder + entrys[i].Name()); err != nil {
				return
			}
			// 只需要提取小说内容
			var str string
			if str, err = getHtmlFromDivId(string(b), divID); err != nil {
				fmt.Printf("%s\n%s", entrys[i].Name(), err)
				return
			}

			str = strings.ReplaceAll(str, "&nbsp;", "")
			// 为了创建epub格式
			str = `<h1>` + entrys[i].Name() + "</h1>" + str
			// fmt.Println(str)

			if _, err = ep.AddSection(str, entrys[i].Name(), "", ""); err != nil {
				return
			}
			fmt.Println(i)
		}
	}

	return ep.Write(savePath)
}
