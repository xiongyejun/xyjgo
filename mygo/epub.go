package main

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"

	epub "github.com/bmaupin/go-epub"
)

func (me *EpubSet) create() (err error) {
	me.ep = epub.NewEpub(me.Name)
	// 封面图片
	if me.CoverPicFile != "" {
		var picpath string
		if picpath, err = me.ep.AddImage(me.CoverPicFile, me.CoverPicFile); err != nil {
			return
		}
		me.ep.SetCover(picpath, "")
	}

	path := me.SrcFolderPath
	if path, err = filepath.Abs(path); err != nil {
		return
	}
	index := strings.LastIndex(path, strPathSeparator)
	name := path[index+1:]
	path = path[:index] + strPathSeparator
	var fi *FileInfo
	if fi, err = scanDir(path, name, 0); err != nil {
		return
	}

	if err = me.addSection(fi, 0); err != nil {
		return
	}

	me.ep.SetAuthor(me.Author)
	return me.ep.Write(me.Name + ".epub")
}

func (me *EpubSet) addSection(fi *FileInfo, iLevel int) (err error) {
	if fi.IsDir {
		if err = me.addSubDir(fi); err != nil {
			panic(err)
		}

		if fi.HasFile {
			for i := range fi.Subs {
				var b []byte
				if b, err = ioutil.ReadFile(fi.Subs[i].Path + fi.Subs[i].Name); err != nil {
					panic(err)
				}
				// 只需要提取小说内容
				var str string = string(b)
				if me.DivID != "" {
					if str, err = getHtmlFromDivId(str, me.DivID); err != nil {
						panic(err)
					}
				} else {
					for j := range me.SplitInfos {
						str = strings.Split(str, me.SplitInfos[j].Sep)[me.SplitInfos[j].RetIndex]
					}
				}

				for j := range me.ReplaceExpr {
					var reg *regexp.Regexp
					if reg, err = regexp.Compile(me.ReplaceExpr[j].Expr); err != nil {
						panic(err)
					}
					str = reg.ReplaceAllString(str, me.ReplaceExpr[j].New)
				}

				// 为了创建epub格式
				str = `<h2>` + fi.Subs[i].Name[6:] + "</h2>" + str
				str += `<br /><a href="` + fi.Name + `.xhtml">` + fi.Name + `</a>` // 跳转到目录
				str += `<br /><a href="目录.xhtml">目录</a>`                           // 跳转到目录

				if _, err = me.ep.AddSection(str, fi.Subs[i].Name[6:], fi.Subs[i].Name+".xhtml", ""); err != nil {
					panic(err)
				}
			}

		} else {
			// 子目录，只能是没有文件的目录才是包含子目录的
			for i := range fi.Subs {
				if err = me.addSection(fi.Subs[i], iLevel+1); err != nil {
					panic(err)
				}
			}
		}

	}

	return
}

// 把Subs创建目录
func (me *EpubSet) addSubDir(fi *FileInfo) (err error) {
	// 目录
	var bodys []string = make([]string, len(fi.Subs))
	for i := range fi.Subs {
		var filename string = fi.Subs[i].Name
		if !fi.Subs[i].IsDir {
			filename = filename[6:]
		}
		bodys[i] = `  <li><a href="` + fi.Subs[i].Name + `.xhtml">` + filename + `</a></li>`
	}

	if _, err = me.ep.AddSection("<ol>\n"+strings.Join(bodys, "\n")+"\n</ol>"+
		`<br /><a href="目录.xhtml">目录</a>`,
		fi.Name, fi.Name+".xhtml", ""); err != nil {
		return
	}

	return
}
