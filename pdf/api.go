package pdf

import (
	"bytes"
	"compress/zlib"
	"errors"
	"io"
	"os"
	"regexp"

	"strconv"

	"github.com/xiongyejun/xyjgo/fileHeader"
)

func Parse(f *os.File) (p *PDF, err error) {
	p = new(PDF)
	p.f = f
	var fInfo os.FileInfo
	if fInfo, err = p.f.Stat(); err != nil {
		return
	}
	p.fSize = fInfo.Size()

	p.Header = make([]byte, 8)
	if _, err = p.f.Read(p.Header); err != nil {
		return
	}

	if !fileHeader.IsPDF(p.Header) {
		return nil, errors.New("不是PDF文件")
	}

	if err = p.getTrailer(); err != nil {
		return
	}

	if err = p.getCatalog(); err != nil {
		return
	}
	return
}

func (me *PDF) GetPageByte(pageIndex int) (ret []byte, err error) {
	var b []byte
	if b, err = me.readObjByte(me.pagesInfo[pageIndex].ContentsObjIndex); err != nil {
		return
	}
	iStart := bytes.Index(b, []byte("stream"))
	if iStart == -1 {
		err = errors.New("没有找到字符串[stream]。")
	}
	iEnd := bytes.Index(b, []byte("endstream"))
	if iEnd == -1 {
		err = errors.New("没有找到字符串[endstream]。")
	}
	ret = b[iStart+len("stream")+1 : iEnd] // +1有1个换行符0x0A
	// 是否压缩了
	if bytes.Contains(b, []byte("/Filter /FlateDecode")) {
		buf := bytes.NewReader(ret)
		var r io.ReadCloser
		if r, err = zlib.NewReader(buf); err != nil {
			return
		}
		defer r.Close()
		var out bytes.Buffer
		if _, err = io.Copy(&out, r); err != nil {
			return
		}

		ret = out.Bytes()
	}

	return
}

// 解析页面数据
func (me *PDF) ParsePageByte(pageByte []byte) (ret []byte, err error) {
	// 由()包含起来的一个字串,中间可以使用转义符"/".
	//  (abc) 表示abc
	//  (a//) 表示a/
	// 转义字符见下表

	// 由<>包含起来的一个16进制串,两位表示一个字符,不足两位用0补齐
	// <Aabb> 表示AA和BB两个字符
	// <AAB> 表示AA和B0两个字符

	var re *regexp.Regexp
	//
	var expr string = `Tf.*?(<[0-9a-zA-Z]+?>).*?Tj|Tf.*?(\(.+?\)).*?Tj`
	if re, err = regexp.Compile(expr); err != nil {
		return
	}
	bb := re.FindAllSubmatch(pageByte, -1)
	if len(bb) == 0 {
		err = errors.New("没有找到“Tf(xx)Tj”或者“Tf<xx>Tj”")
		return
	}

	var b []byte
	for i := range bb {
		print("bb=" + string(bb[i][1]) + "  bb1=" + string(bb[i][2]) + "\n")
		if len(bb[i][1]) != 0 { // '<'
			b = bb[i][1][1 : len(bb[i][1])-1]
			if len(b)%2 == 1 {
				b = append(b, '0')
			}
			for j := 0; j < len(b); j += 2 {
				strHex := "0x" + string(b[j:j+2])
				var n int64
				if n, err = strconv.ParseInt(strHex, 0, 64); err != nil {
					return
				}
				ret = append(ret, byte(n))
			}
		} else {
			b = bb[i][2][1 : len(bb[i][2])-1]
			if b[0] == '!' {
				b[0] = 0x0A
			}
			ret = append(ret, b...)
		}
	}
	return
}

// 转义字符	含义
// /n	换行
// /r	回车
// /t	水平制表符
// /b	退格
// /f	换页（Form feed (FF)）
// /(	左括号
// /)	右括号
// //	反斜杠
// /ddd	八进制形式的字符
