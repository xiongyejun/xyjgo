package pdf

import (
	"errors"
	"os"
	"regexp"
)

//-----------------------------------
// 3 0 obj <<

// /Type /Catalog

// /Pages 1 0 R
// >> endobj
//-----------------------------------

// 从obj中找到 /str # # R，返回x
func findR(b []byte, str string) (obj string, err error) {
	var re *regexp.Regexp
	var expr string = `/` + str + ` (\d+) \d+ R`
	if re, err = regexp.Compile(expr); err != nil {
		return
	}
	bb := re.FindSubmatch(b)
	if len(bb) == 0 {
		return "", errors.New("没有找到[/" + str + " # # R]\n" + string(b))
	}

	return string(bb[1]), nil
}

// 从文件后面读取nBytes个字节，如果nBytes=-1，nBytes=512
func readFileEndByte(f *os.File, nBytes int64) (ret []byte, err error) {
	if nBytes == -1 {
		nBytes = 512
	}

	var fInfo os.FileInfo
	if fInfo, err = f.Stat(); err != nil {
		return
	}

	var nOffset int64
	if fInfo.Size() < nBytes {
		nBytes = fInfo.Size()
		nOffset = 0
	} else {
		nOffset = fInfo.Size() - nBytes
	}

	ret = make([]byte, nBytes)
	if _, err = f.ReadAt(ret, nOffset); err != nil {
		return
	}

	return
}
