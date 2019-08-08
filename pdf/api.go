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
	// trailer 说明文件尾trailer对象的开始
	rootIndex := bytes.Index(me.Src, []byte("trailer"))

	if rootIndex == -1 {
		return errors.New("没有找到[trailer]")
	}

	b := me.Src[rootIndex:]
	// /Root # # R说明根对象的对象号为#
	if me.RootR, err = findR(b, "Root"); err != nil {
		return
	}

	// Startxref
	// ### 说明交叉引用表的偏移地址

	// xref 交叉引用表
	// #1 #2 说明下面各行所描述的对象号是从#1开始，并且有#2个对象
	// 0000000000 65535 f 一般每个PDF文件都是以这一行开始交叉应用表的，说明对象0的起始地址为0000000000，产生号（generation number）为65535，也是最大产生号，不可以再进行更改，而且最后对象的表示是f,表明该对象为free, 这里，大家可以看到，其实这个对象可以看作是文件头
	// …………
	// 0000000322 00000 n  对象n的偏移地址为322
	// …………
	return nil
}

// 用正则获取对象
// ------------------------------------
// # # obj
// ...
// endobj
// ------------------------------------
// 第一个数字#称为对象号，来唯一标识一个对象的
// 第二个#是产生号，是来表明它在被创建后的第几次修改
// 所有新创建的PDF文件的对象号应该都是0，即第一次被创建以后没有被修改过
func (me *PDF) getObj() (err error) {
	var re *regexp.Regexp
	// (?标记)               在组内设置标记，非捕获，标记影响当前组后的正则表达式
	// s              让 . 匹配 \n (默认为 false)
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
		me.objs[i].strIndex = string(objs[i][1]) // 记录对象号
		me.mObj[me.objs[i].strIndex] = i         // 记录到map中，key是对象号
		me.objs[i].b = objs[i][0]                // 对象的byte（在<< 和>>之间的）
		me.objs[i].indexSrc = indexs[i][0]       // 对象的byte在src中出现的位置
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
