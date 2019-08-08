// Portable Document Format的简称，意为“便携式文档格式”
// https://www.jianshu.com/p/51eb811ba935
package pdf

import (
	"errors"
	"io/ioutil"

	"github.com/xiongyejun/xyjgo/fileHeader"
)

type obj struct {
	strIndex string // obj # 0
	indexSrc int    // 在Src里的位置
	b        []byte // 数据
}
type PDF struct {
	Src    []byte
	RootR  string // trailer里Root对应的obj
	Header []byte // %PDF-1.0

	mObj map[string]int // 记录obj所在的objs的下标
	objs []obj

	CountPage int
}

func New(FilePath string) (p *PDF, err error) {
	p = new(PDF)
	if p.Src, err = ioutil.ReadFile(FilePath); err != nil {
		return
	}

	if !fileHeader.IsPDF(p.Src) {
		return nil, errors.New("不是PDF文件")
	}
	return
}

func (me *PDF) GetObjByte(strObjIndex string) (b []byte, err error) {
	var index int
	if index, err = me.GetObjIndex(strObjIndex); err != nil {
		return
	}
	b = me.objs[index].b
	return

}
func (me *PDF) GetObjIndex(strObjIndex string) (i int, err error) {
	if index, ok := me.mObj[strObjIndex]; !ok {
		return -1, errors.New("没有找到[" + strObjIndex + "]")
	} else {
		i = index
		return
	}
}
