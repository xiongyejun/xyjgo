package main

import (
	"io/ioutil"
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

	if _, err = me.scanDir(me.SrcFolderPath, "", 0); err != nil {
		return
	}

	if me.nav != "" {
		if err = ioutil.WriteFile("nav.txt", []byte(me.nav), 0666); err != nil {
			return
		}
		println("有多级目录，请用[nav.txt]手动修改[nav.xhtml]文件")
	}

	me.ep.SetAuthor(me.Author)
	return me.ep.Write(me.Name + ".epub")
}

func (me *EpubSet) scanDir(dirPath, dirName string, iLevel int) (nav string, err error) {
	var retDirs, retFiles []string
	if retDirs, retFiles, err = scanDir(dirPath); err != nil {
		return
	}
	// 如果有子文件夹，那就是多层次情况，找到最底层的文件夹为止
	if len(retDirs) > 0 {
		for i := range retDirs {
			space := strings.Repeat("  ", iLevel)
			var navDir string
			if navDir, err = me.ep.AddSection(retDirs[i], retDirs[i], "", ""); err != nil {
				return
			}
			navDir = space + "<li>\n" + space + " <a href=\"xhtml/" + navDir + "\">" + retDirs[i] + "</a>\n"

			if nav, err = me.scanDir(dirPath+retDirs[i]+strPathSeparator, retDirs[i], iLevel+1); err != nil {
				return
			}

			me.nav = me.nav + "\n\n" + navDir + nav + "\n" + space + "</li>"
		}
	} else {

		// 这种时候才真正需要添加所需要的文件
		navs := make([]string, len(retFiles))
		for i := range retFiles {
			if navs[i], err = me.addSection(dirPath, retFiles[i], iLevel); err != nil {
				return
			}
		}
		nav = strings.Join(navs, "\n")
	}

	return
}

func (me *EpubSet) addSection(dirPath, fileName string, iLevel int) (nav string, err error) {
	var b []byte
	if b, err = ioutil.ReadFile(dirPath + fileName); err != nil {
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
		if reg, err = regexp.Compile(me.ReplaceExpr[j].Expr); err != nil {
			return
		}
		str = reg.ReplaceAllString(str, me.ReplaceExpr[j].New)
	}

	// 为了创建epub格式
	str = `<h2>` + fileName[6:] + "</h2>" + str

	if nav, err = me.ep.AddSection(str, fileName[6:], "", ""); err != nil {
		return
	}

	space := strings.Repeat("  ", iLevel+1)
	//	      <li>
	//          <a href="xhtml/section0001.xhtml">离胜利只差一步1</a>
	//        </li>
	nav = space + "<li>\n" + space + " <a href=\"xhtml/" + nav + "\">" + fileName[6:] + "</a>\n" + space + "</li>"

	return
}
