// Portable Document Format的简称，意为“便携式文档格式”
// https://www.jianshu.com/p/51eb811ba935 简单讲解
// https://blog.csdn.net/steve_cui/article/details/81910632 详细讲解
package pdf

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"regexp"
	"strconv"
)

// PDF页面内容类型
type PDF_Resources int

const (
	RSC_Text  PDF_Resources = 0x01 // 0000 0000 0000 0001
	RSC_Image               = 0x02 // 0000 0000 0000 0010
)

// 页面信息
type pageInfo struct {
	ResourcesObjIndex int
	ContentsObjIndex  int
	ResourcesType     PDF_Resources
}

type PDF struct {
	f     *os.File
	fSize int64

	trailerOffset int64   // trailer偏移未知
	xrefOffset    int64   // 交叉引用表偏移位置
	objsOffset    []int64 // 对象的偏移位置，通过交叉引用表读取
	objStartIndex int     // xref 下面1行的2个数字
	objCount      int

	RootR        string // trailer里Root对应的obj
	rootObjIndex int

	pagesObjIndex int
	Pages         int

	kidsObj   []int // 记录的是page的obj index
	pagesInfo []pageInfo
	Header    []byte // %PDF-1.0
}

//PDF解析流程：
//a）从trailer中找到Root关键字
// Root是指向Catalog字典，Catalog是一个PDF文件的总入口，它包含Page tree，Outline hierarchy等。
func (me *PDF) getTrailer() (err error) {
	// trailer 说明文件尾trailer对象的开始
	var b []byte
	var nOffset int64 = 512 // 相对文件尾部
	if b, err = readFileEndByte(me.f, nOffset); err != nil {
		return
	}
	n := int64(bytes.Index(b, []byte("trailer")))
	if n == -1 {
		return errors.New("没有找到[trailer]")
	}
	me.trailerOffset = n + me.fSize - nOffset

	b = b[n:]
	// /Root # # R说明根对象的对象号为#
	if me.rootObjIndex, err = findR(b, "Root"); err != nil {
		return
	}

	return me.getXref(b)
}

// 读取交叉引用表
// 交叉引用表记录的是每个对象的偏移地址
func (me *PDF) getXref(bTrailer []byte) (err error) {
	// startxref
	// ### 说明交叉引用表的偏移地址
	n := bytes.Index(bTrailer, []byte("startxref"))
	if n == -1 {
		return errors.New("没有找到[xrefOffset]")
	}
	me.xrefOffset = int64(n) + me.trailerOffset
	// xref 交叉引用表
	// #1 #2 说明下面各行所描述的对象号是从#1开始，并且有#2个对象
	// 0000000000 65535 f 一般每个PDF文件都是以这一行开始交叉应用表的，说明对象0的起始地址为0000000000，产生号（generation number）为65535，也是最大产生号，不可以再进行更改，而且最后对象的表示是f,表明该对象为free, 这里，大家可以看到，其实这个对象可以看作是文件头
	// …………
	// 0000000322 00000 n  对象x的偏移地址为322
	// …………

	// 读取xref的offset
	if _, err = me.f.Seek(me.xrefOffset, 0); err != nil {
		return
	}
	bf := bufio.NewReader(me.f)

	var b []byte
	if b, _, err = bf.ReadLine(); err != nil {
		return
	}
	if string(b) != "startxref" {
		return errors.New("读取startxref出错！")
	}

	if b, _, err = bf.ReadLine(); err != nil {
		return
	}
	if n, err = strconv.Atoi(string(b)); err != nil {
		return errors.New("startxref 下一行的offset地址转换int出错！" + string(b))
	}
	me.xrefOffset = int64(n)

	// 开始读取xref
	if _, err = me.f.Seek(me.xrefOffset, 0); err != nil {
		return
	}
	bf = bufio.NewReader(me.f)
	if b, _, err = bf.ReadLine(); err != nil {
		return
	}
	if string(b) != "xref" {
		return errors.New("读取xref出错！")
	}
	if b, _, err = bf.ReadLine(); err != nil {
		return
	}
	bSlice := bytes.Split(b, []byte(" "))
	if len(bSlice) != 2 {
		return errors.New("读取xref下面1行的对象[开始号]和[数量]出错！")
	}
	if me.objStartIndex, err = strconv.Atoi(string(bSlice[0])); err != nil {
		return errors.New("xref下面1行的对象[开始号]转换int出错！")
	}
	if me.objCount, err = strconv.Atoi(string(bSlice[1])); err != nil {
		return errors.New("xref下面1行的对象[数量]转换int出错！")
	}

	me.objsOffset = make([]int64, me.objCount)
	var k int = 0
	if b, _, err = bf.ReadLine(); err != nil {
		return
	}
	for string(b) != "trailer" {
		// 固定格式：0000003278 00000 n
		if n, err = strconv.Atoi(string(b[:10])); err != nil {
			return errors.New("xref下面对象[地址]转换int出错！" + string(b))
		}
		me.objsOffset[k] = int64(n)
		k++
		if b, _, err = bf.ReadLine(); err != nil {
			return
		}
	}

	return
}

// 读取Catalog（目录）
// 可以读取到pages count总页数，以及kids——每一个页面信息的obj index
func (me *PDF) getCatalog() (err error) {
	// RootR指向的obj
	var bRoot []byte
	if bRoot, err = me.readObjByte(me.rootObjIndex); err != nil {
		return
	}
	if !bytes.Contains(bRoot, []byte("<< /Type /Catalog /Pages ")) {
		return errors.New("读取的Root不包含[<< /Type /Catalog /Pages ]!\n" + string(bRoot))
	}

	//  读取Pages
	if me.pagesObjIndex, err = findR(bRoot, "Pages"); err != nil {
		return
	}
	var bPages []byte
	if bPages, err = me.readObjByte(me.pagesObjIndex); err != nil {
		return
	}
	if !bytes.Contains(bPages, []byte("<< /Type /Pages ")) {
		return errors.New("读取的Pages不包含[<< /Type /Pages  ]!\n" + string(bPages))
	}

	x := bytes.Index(bPages, []byte("/Count "))
	if x == -1 {
		return errors.New("没有找到[/Count ]")
	}
	b := bPages[x+len("/Count "):]
	x = bytes.Index(b, []byte(" "))
	b = b[:x]
	if me.Pages, err = strconv.Atoi(string(b)); err != nil {
		return errors.New("Count # # R转化int出错。")
	}

	return me.getKids(bPages)
}

func (me *PDF) getKids(bPages []byte) (err error) {
	n := bytes.Index(bPages, []byte("/Kids "))
	if n == -1 {
		errors.New("读取的Root不包含[Kids ]!" + string(bPages))
	}
	bPages = bPages[n+len("/Kids "):]
	// 读取kids
	// /Kids [4 0 R]说明页的对象为4
	// 如果有多个页面，就多个页面直接连续下去
	// 比如说/Kids [ 4 0 R 10 0 R ], 就说明该PDF的第一页的对象号是4,第二页的对象号是10
	var re *regexp.Regexp
	var expr string = `(\d+) \d+ R`
	if re, err = regexp.Compile(expr); err != nil {
		return
	}

	bb := re.FindAllSubmatch(bPages, -1)
	if len(bb) == 0 {
		return errors.New("没有找到[Kids后面的 # # R]")
	}

	me.kidsObj = make([]int, len(bb))
	for i := range me.kidsObj {
		if me.kidsObj[i], err = strconv.Atoi(string(bb[i][1])); err != nil {
			return
		}
	}
	return me.getPagesInfo()
}

// 读取页面信息：主要是Resources（记录pdf页面信息的类型），Contents（页面内容）
func (me *PDF) getPagesInfo() (err error) {
	var b []byte
	me.pagesInfo = make([]pageInfo, len(me.kidsObj))
	for i := range me.kidsObj {
		if b, err = me.readObjByte(me.kidsObj[i]); err != nil {
			return
		}
		// << /Type /Page /Parent 3 0 R /Resources 6 0 R /Contents 4 0 R /MediaBox [0 0 595.28 841.89] >>
		if !bytes.Contains(b, []byte("<< /Type /Page ")) {
			return errors.New("读取的Page不包含[<< /Type /Page ]!\n" + string(b))
		}

		// 读取Resources和Contents的对象index
		if me.pagesInfo[i].ResourcesObjIndex, err = findR(b, "Resources"); err != nil {
			return
		}

		if me.pagesInfo[i].ContentsObjIndex, err = findR(b, "Contents"); err != nil {
			return
		}

		// 读取Resources的Type
		if b, err = me.readObjByte(me.pagesInfo[i].ResourcesObjIndex); err != nil {
			return
		}
		// << /ProcSet [ /PDF /Text ] /ColorSpace << /Cs2 11 0 R /Cs1 7 0 R >> /Font
		// << /C1 8 0 R /TT1 9 0 R /TT2 10 0 R >> >>
		if bytes.Contains(b, []byte("[ /PDF /Text ]")) {
			me.pagesInfo[i].ResourcesType = RSC_Text
		} else if bytes.Contains(b, []byte("[ /PDF /Image ]")) {
			me.pagesInfo[i].ResourcesType = RSC_Image
		} else if bytes.Contains(b, []byte("/Image")) && bytes.Contains(b, []byte("/Text")) {
			me.pagesInfo[i].ResourcesType = RSC_Image | RSC_Text
		} else {
			print(string(b))
			return errors.New("读取ResourcesType出错。\n" + string(b))
		}
	}
	return
}

func (me *PDF) readObjByte(objIndex int) (ret []byte, err error) {
	if _, err = me.f.Seek(me.objsOffset[objIndex], 0); err != nil {
		return
	}
	bf := bufio.NewReader(me.f)
	var b []byte
	for string(b) != "endobj" {
		if b, _, err = bf.ReadLine(); err != nil {
			return
		}
		ret = append(ret, b...)
		ret = append(ret, 0x0A)
	}
	return
}

//简单流程
//trailer→ Root→ Catalog→ Pages→ Page→ Contents
