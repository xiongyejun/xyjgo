// txt格式转换
package main

import (
	"bufio"
	"bytes"
	"fmt"

	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const (
	_ = iota
	src序号
	src记账日
	src起息日
	src交易类型
	src凭证
	src摘要
	src金额_借
	src金额_贷
	src金额_余额
	src流水
	src备注
)

const (
	des账号 = iota
	des日期
	des柜员
	des交易套号
	des套内序号
	des起始日
	des交易金额
	des备注
)

// 源
// 每一组格式：
// 第1行：1  账号     366258341155         账户名称  ……（提取出账号）
// 接下来3行不需要
// 第5行： ────（横杆在txt里绘制的表格）
// 第6行： |序号……
// 第7行： |No.
// 第8行： ───
// 接下来是数据开始行，有序号1、2、3……20，最多20行，结束行号x
// x+1: ──────
// x+2:   借方合计  ……
// x+3:   Debit Total ……
// x+4:  1. 余额前面标注"-"…………
// 直到下1个第1行。
type txtSrc struct {
	fileName string
	f        *os.File
}

// 转换后
type txtDes struct {
	fileName string
	f        *os.File
}

type datas struct {
	ts *txtSrc
	td *txtDes

	strDir string
	files  []string
	cLine  chan []byte
}

var d *datas

func init() {
	d = new(datas)
	d.ts = new(txtSrc)
	d.td = new(txtDes)
	d.cLine = make(chan []byte, 100)

	d.files = make([]string, 0)

}

var bSEP []byte = []byte("|")

func main() {
	//	var err error

	//	for i := range d.files {

	//	}

}

func (me *datas) getResult(i int) (err error) {

	if d.ts.f, err = os.Open(me.files[i]); err != nil {
		return
	}
	defer d.ts.f.Close()

	os.Mkdir(me.strDir+"/conv/", os.ModeDir)
	if d.td.f, err = os.OpenFile("conv/"+me.files[i], os.O_CREATE, os.ModeAppend); err != nil {
		return
	}
	defer d.td.f.Close()
	// 标题
	d.td.f.Write([]byte{0x20, 0x20, 0x20, 0x20, 0x20, 0xD5, 0xCA, 0xBA, 0xC5, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xC8, 0xD5, 0xC6, 0xDA, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xB9, 0xF1, 0xD4, 0xB1, 0x20, 0xBD, 0xBB, 0xD2, 0xD7, 0xCC, 0xD7, 0xBA, 0xC5, 0x20, 0xCC, 0xD7, 0xC4, 0xDA, 0xD0, 0xF2, 0xBA, 0xC5, 0x20, 0xC6, 0xF0, 0xCA, 0xBC, 0xC8, 0xD5, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xBD, 0xBB, 0xD2, 0xD7, 0xBD, 0xF0, 0xB6, 0xEE, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xB1, 0xB8, 0xD7, 0xA2, 0x20, 0x20, 0x20, 0x20, 0x20})

	// 正则
	var reg *regexp.Regexp
	//								序号        记账日  起息日    Type vou ……     余额
	if reg, err = regexp.Compile(` \| .{2,2} \|\d{6,6}\|\d{6,6}\|.*?\|.*?\|.*?\|.*?\|.*?\|.*?\|.*?\|`); err != nil {
		return
	}
	var regzh *regexp.Regexp // 账号
	// 1  账号     366258341155         账户名称
	//	var tmpb []byte = []byte{0x31, 0x20, 0x20, 0xd5, 0xcb, 0xba, 0xc5, 0x20, 0x20, 0x20, 0x20, 0x20}
	//	tmpb = append(tmpb, []byte()...)
	//	tmpb = append(tmpb, []byte{0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xd5, 0xcb, 0xbb, 0xa7, 0xc3, 0xfb, 0xb3, 0xc6}...)

	if regzh, err = regexp.Compile(`1  .*?     (\d{12,12})         `); err != nil {
		return
	}

	bf := bufio.NewReader(me.ts.f)
	var b []byte
	if b, err = ioutil.ReadAll(bf); err != nil {
		return
	}
	ret := reg.FindAll(b, -1)
	retzh := regzh.FindAllSubmatch(b, -1)

	for i := range ret {
		//		d.td.f.Write([]byte{0x0d, 0x0a})
		//		d.td.f.Write(ret[i])
		arrSrc := bytes.Split(ret[i], bSEP)
		var arrDes [][]byte = make([][]byte, des备注+1)
		// 转换
		arrDes[des账号] = retzh[0][1]
		arrDes[des日期] = arrSrc[src记账日]
		arrDes[des起始日] = arrSrc[src起息日]
		arrDes[des备注] = arrSrc[src摘要]
		// 判断金额是在借方还是贷方（看一下最后1个字符是不是空白）
		if arrSrc[src金额_借][len(arrSrc[src金额_借])-1] != 0x20 {
			arrDes[des交易金额] = arrSrc[src金额_借]
		} else {
			arrDes[des交易金额] = arrSrc[src金额_贷]
		}
		if err = d.td.writeIn(arrDes); err != nil {
			return
		}
	}
	return

	//	bf := bufio.NewReader(me.ts.f)
	//	var arrDes [][]byte = make([][]byte, des备注+1)
	//	var arrSrc [][]byte
	//	var bWrite bool = false
	//	for {
	//		var b []byte
	//		b, _, err = bf.ReadLine()
	//		if err == io.EOF {
	//			break
	//		}

	//		if b[0] == 0x20 && b[1] == 0x7c && b[2] == 0x20 { // " | "
	//			fmt.Printf("%s\r\n", b)
	//			if b[3] == 0x20 {
	//				// 还是属于上面一行的内容，一般是摘要列写不下
	//				arrSrc = bytes.Split(b[:len(b)-1], bSEP)
	//				arrDes[des备注] = append(arrDes[des备注], arrSrc[src摘要]...)
	//			} else {
	//				if bWrite {
	//					bWrite = false
	//					if err = d.td.writeIn(arrDes); err != nil {
	//						return
	//					}
	//				}
	//				arrSrc = bytes.Split(b, bSEP)
	//				// 转换
	//				arrDes[des账号] = []byte("xxxxx")
	//				arrDes[des日期] = arrSrc[src记账日]
	//				arrDes[des起始日] = arrSrc[src起息日]
	//				arrDes[des备注] = arrSrc[src摘要]
	//				// 判断金额是在借方还是贷方（看一下最后1个字符是不是空白）
	//				if arrSrc[src金额_借][len(arrSrc[src金额_借])-1] != 0x20 {
	//					arrDes[des交易金额] = arrSrc[src金额_借]
	//				} else {
	//					arrDes[des交易金额] = arrSrc[src金额_贷]
	//				}
	//				bWrite = true
	//			}
	//		} else {
	//			if bWrite {
	//				bWrite = false
	//				if err = d.td.writeIn(arrDes); err != nil {
	//					return
	//				}
	//			}

	//		}

	//	}

	fmt.Println("getResult")
	return nil
}

func (me *txtDes) writeIn(b [][]byte) (err error) {
	b[des账号] = append(b[des账号], []byte("  ")...)
	b[des账号] = append([]byte{0x0d, 0x0a}, b[des账号]...)
	b[des日期] = append(b[des日期], []byte("                           ")...)
	b[des日期] = append([]byte("20"), b[des日期]...)

	// 交易金额是右对齐
	var nSpace int = len("          11621.35") - len(b[des交易金额])
	if nSpace > 0 {
		b[des起始日] = append(b[des起始日], []byte(strings.Repeat(" ", nSpace))...)
	}
	b[des起始日] = append([]byte("20"), b[des起始日]...)

	b[des交易金额] = append(b[des交易金额], []byte("   ")...)

	if _, err = me.f.Write(bytes.Join(b, nil)); err != nil {
		return
	}
	b = nil

	return
}

func (me *datas) getFiels() (err error) {
	entrys, err := ioutil.ReadDir(me.strDir)
	if err != nil {
		return
	}

	for _, entry := range entrys {
		if !entry.IsDir() {
			if strings.ToLower(entry.Name()[len(entry.Name())-4:]) == ".txt" {
				d.files = append(d.files, entry.Name())
			}
		}
	}

	return nil
}
