package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

var ch chan int
var herf string

func (me *DownSet) down() (err error) {
	var b []byte

	if b, err = ioutil.ReadFile(me.DirInfoJsonFile); err != nil {
		return
	}

	herf = me.PreHerf

	var di dirInfos = dirInfos{}
	if err = json.Unmarshal(b, &di); err != nil {
		return
	}
	di.down()

	for i := range di.DirInfos {
		<-ch
		i++
	}
	return
}

//  下载小说
func (me *dirInfos) down() {
	ch = make(chan int, 50)

	fmt.Printf("开始下载，总共%d个。\n", len(me.DirInfos))
	for i := range me.DirInfos {
		// if i == 2 {
		// 	return
		// }
		go me.DirInfos[i].down(i)
	}

	return
}

func (me *DirInfo) down(j int) {
	var b []byte
	var err error
	var count int = 0
	filename := strconv.Itoa(j)
	filename = ("000000" + filename)[len(filename):]
	filename += me.Name
	filename = "srcHtml/" + filename

	// 不管下载是否成功都要往ch里写入数据
	defer func() {
		ch <- j
	}()

	// 存在的话就不需要再下载
	if exists(filename) {
		return
	}
	fmt.Printf("开始下载：%d, %s, %s,  %s\n", j, herf+me.Href, me.Name, filename)

	// 尝试20次下载
	for count < 20 {
		if b, err = httpGet(herf + me.Href); err != nil {
			count++
		} else {
			break
		}
	}
	if err != nil {
		me.PrintErr(err)
		return
	}

	// 保存下载内容

	if err = ioutil.WriteFile(filename, b, 0666); err != nil {
		me.PrintErr(err)
		return
	}
	fmt.Printf("完成下载：%d, %s\n", j, herf+me.Href)

	return
}

func (me *DirInfo) PrintErr(err error) {
	fmt.Printf("Err: herf: %s, name: %s\n%s\n", herf+me.Href, me.Name, err)
}

type DirInfo struct {
	Href string
	Name string
}

type dirInfos struct {
	DirInfos []DirInfo
}

// 根据正则提取目录中记录的每一个章节的地址和名称
func (me *DirSet) getDirInfo() (ret dirInfos, err error) {
	var reg *regexp.Regexp
	if reg, err = regexp.Compile(me.Expr); err != nil {
		return
	}
	var bHtml []byte
	if bHtml, err = ioutil.ReadFile(me.DirHtmlFile); err != nil {
		return
	}
	bbb := reg.FindAllSubmatch(bHtml, -1)
	for i := range bbb {
		tmp := DirInfo{string(bbb[i][1]), string(bbb[i][2])}
		ret.DirInfos = append(ret.DirInfos, tmp)
	}
	return
}

func (me *dirInfos) saveJsonTxt(path string) (err error) {
	var b []byte
	if b, err = json.MarshalIndent(me, "", "\t"); err != nil {
		return
	}
	return ioutil.WriteFile(path, b, 0666)
}
