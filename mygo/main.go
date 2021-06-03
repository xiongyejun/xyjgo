//一些自己偶尔使用的功能
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	epub "github.com/bmaupin/go-epub"
)

var strPathSeparator = string(os.PathSeparator)

func main() {
	if len(os.Args) == 1 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "xsdir":
		fDirSet()
	case "xsdown":
		fDownSet()
	case "xsepub":
		fEpubSet()
	case "xsset":
		xsTemplateSet()
	case "dl":
		if len(os.Args) != 4 {
			printHelp()
			return
		}
		if err := download(os.Args[2], os.Args[3]); err != nil {
			fmt.Println(err)
		}
		fmt.Println("download ok")

	case "lp":
		if len(os.Args) != 3 {
			printHelp()
			return
		}
		if spath, err := exec.LookPath(os.Args[2]); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(spath)
		}
	case "wj":
		if len(os.Args) != 3 {
			printHelp()
			return
		}

		if err := webp2jpgDir(os.Args[2]); err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("未设置的命令。")
		printHelp()
	}
}

type DirSet struct {
	DirHtmlFile string // dir.html路径
	Expr        string // 提取的正则表达式
}

// 解析目录，获取每一个章节的信息
func fDirSet() {
	var ds *DirSet = new(DirSet)
	var b []byte
	var err error
	if b, err = ioutil.ReadFile("dir.set"); err != nil {
		fmt.Println(err)
		return
	}
	if err = json.Unmarshal(b, ds); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", b)

	var r dirInfos
	if r, err = ds.getDirInfo(); err != nil {
		fmt.Println(err)
		return
	}

	if err = r.saveJsonTxt("dirJson.txt"); err != nil {
		fmt.Println(err)
		return
	}
}

type DownSet struct {
	DirInfoJsonFile string // dirJson.txt路径
	PreHerf         string // dirJson.txt里记录的herf需要的前缀路径
	SleepSecond     uint32
}

// 读取dirJson.txt信息，下载每一个章节的原始网页
func fDownSet() {
	var ds *DownSet = new(DownSet)
	var b []byte
	var err error
	if b, err = ioutil.ReadFile("down.set"); err != nil {
		fmt.Println(err)
		return
	}
	if err = json.Unmarshal(b, ds); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", b)

	os.RemoveAll("srcHtml")
	os.Mkdir("srcHtml", 0666)
	if err = ds.down(); err != nil {
		fmt.Println(err)
		return
	}
}

type SplitInfo struct {
	Sep      string
	RetIndex int // 分割后需要获取数组的下标
}
type ReplaceInfo struct {
	Expr string
	New  string
}

type EpubSet struct {
	Name          string // 创建的文件名称
	Author        string // 书的作者
	CoverPicFile  string // 封面图片的路径
	SrcFolderPath string // 每一个章节的原始网页保存路径

	DivID   string // 按DivID提取内容，如果为空，则应该按照SplitSep符号来分割
	CharSet string

	SplitInfos  []SplitInfo
	ReplaceExpr []ReplaceInfo // 有些广告需要替换的正则表达式

	ep  *epub.Epub
	nav string // 多层次的情况要自己创建目录，改写nav.xhtml
}

// 创建epub格式的文件
func fEpubSet() {
	var es *EpubSet = new(EpubSet)
	var b []byte
	var err error
	if b, err = ioutil.ReadFile("epub.set"); err != nil {
		fmt.Println(err)
		return
	}
	if err = json.Unmarshal(b, es); err != nil {
		fmt.Println(err)
		return
	}

	if !strings.HasSuffix(es.SrcFolderPath, strPathSeparator) {
		es.SrcFolderPath += strPathSeparator
	}
	fmt.Printf("%s\n", b)

	if err = es.create(); err != nil {
		fmt.Println(err)
		return
	}
}

// 输出set的模板格式
func xsTemplateSet() {
	var b []byte
	var err error

	ds := DirSet{`C:\dir.html`, `https://www.url/(.*?)"aaa(.*?) //必须有2个(.*?)，注意有些符号的转义`}
	if b, err = json.MarshalIndent(&ds, "", "\t"); err != nil {
		fmt.Println(err)
		return
	}
	if err = ioutil.WriteFile(`dir.set`, b, 0666); err != nil {
		fmt.Println(err)
		return
	}

	downs := DownSet{`C:\dirJson.txt //dirJson.txt路径`, `https://www.url/ //dirJson.txt里记录的herf需要的前缀路径`, 0}
	if b, err = json.MarshalIndent(&downs, "", "\t"); err != nil {
		fmt.Println(err)
		return
	}
	if err = ioutil.WriteFile(`down.set`, b, 0666); err != nil {
		fmt.Println(err)
		return
	}

	es := EpubSet{
		Name:          "创建的文件名称",
		Author:        "书的作者",
		CoverPicFile:  "封面图片的路径",
		SrcFolderPath: "每一个章节的原始网页保存路径",

		DivID:   `按DivID提取内容，如果为空，则应该按照SplitSep符号来分割 <div id=\"book_text\">`,
		CharSet: "UTF-8",

		SplitInfos: []SplitInfo{
			SplitInfo{
				Sep:      "分隔符1",
				RetIndex: 1, //  分割后需要获取数组的下标
			},
			SplitInfo{
				Sep:      "分隔符2",
				RetIndex: 0, //  分割后需要获取数组的下标
			},
		},

		ReplaceExpr: []ReplaceInfo{
			{
				Expr: "有些广告需要替换的正则表达式1",
				New:  "替换为，一般为空",
			},
		},
	}

	if b, err = json.MarshalIndent(&es, "", "\t"); err != nil {
		fmt.Println(err)
		return
	}
	if err = ioutil.WriteFile(`epub.set`, b, 0666); err != nil {
		fmt.Println(err)
		return
	}

}

func printHelp() {
	fmt.Println(`
 mygo xsdir--解析小说目录，获取每一个章节的信息，会生成dirJson.txt，使用前先把<dir.set>的信息设置好，其中DirHtmlFile手动下载保存
 mygo xsdown--根据dirJson.txt下载小说，使用前先把<down.set>的信息设置好，固定下载在srcHtml文件夹内
 mygo xsepub--创建epub，使用前先把<epub.set>的信息设置好，如果有folder里有子folder，则构建多层次目录
 mygo xsset--输出set的模板格式
 mygo dl <url> <savename>--download 下载资源
 mygo lp <exe> --LookPath 查找程序的路径
 mygo wj <folder> --webp2jpg
	`)
}
