package pdf

import (
	"bytes"
	"errors"
	"regexp"
)

func (me *PDF) Parse() (err error) {
	if me.Src == nil {
		return errors.New("me.Src == nil, 请先使用New读取PDF。")
	}
	if err = me.getTrailer(); err != nil {
		return
	}
	if err = me.getObj(); err != nil {
		return
	}
	return
}

//PDF解析流程：
//a）从trailer中找到Root关键字，Root是指向Catalog字典，Catalog是一个PDF文件的总入口，它包含Page tree，Outline hierarchy等。
func (me *PDF) getTrailer() (err error) {
	rootIndex := bytes.Index(me.Src, []byte("trailer"))

	if rootIndex == -1 {
		return errors.New("没有找到[trailer]")
	}

	b := me.Src[rootIndex:]

	if me.RootR, err = findR(b, "Root"); err != nil {
		return
	}
	return nil
}

// 用正则获取
// # # obj
// ...
// endobj
func (me *PDF) getObj() (err error) {
	var re *regexp.Regexp
	var expr string = `(?s)(\d{1,}) \d{1,} obj.*?<<(.*?)>>.*?endobj`
	if re, err = regexp.Compile(expr); err != nil {
		return
	}
	objs := re.FindAllSubmatch(me.Src, -1)
	indexs := re.FindAllSubmatchIndex(me.Src, -1)
	if objs == nil {
		return errors.New("当前文档没有找到obj。")
	}
	me.objs = make([]obj, len(objs))
	me.mObj = make(map[string]int, len(objs))
	for i := range objs {
		me.objs[i].strIndex = string(objs[i][1])
		me.mObj[me.objs[i].strIndex] = i
		me.objs[i].b = objs[i][0]
		me.objs[i].indexSrc = indexs[i][0]
	}
	return
}

//b）从Catalog中找到Pages关键字，Pages是PDF所有页面的总入口，即Page Tree Root。
//c）从Pages中找到Kids和Count关键字，Kids中包含Page子节点，Count列出该文档的总页数。到这里我们已经知道PDF文件有多少页了。
func (me *PDF) getPageTreeRoot() (err error) {
	// var bRoot []byte
	// if bRoot, err = me.GetObjByte(me.RootR); err != nil {
	// 	return
	// }

	// var strIndex string
	// if strIndex, err = findR(bRoot, "Pages"); err != nil {
	// 	return
	// }

	// var bPages []byte
	// if bPages, err = me.GetObjByte(strIndex); err != nil {
	// 	return
	// }

	return
}

//d）从Page字典中获取MediaBox、Contents、Resources等信息，MediaBox包含页面宽高信息，Contents包含页面内容，Resources包含页面所需要的资源信息。

//e）从Contents指向的内容流中获取页面内容。

//简单流程

//trailer→ Root→ Catalog→ Pages→ Page→ Contents
