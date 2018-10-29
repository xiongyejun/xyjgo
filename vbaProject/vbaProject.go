// vbaProject
package vbaProject

import (
	"errors"

	"github.com/xiongyejun/xyjgo/compoundFile"
	"github.com/xiongyejun/xyjgo/rleVBA"
	"github.com/xiongyejun/xyjgo/vbaDir"
)

type ModuleInfo struct {
	Type string // ".bas" ".cls"
	Name string
	Code []byte // vba代码是gbk编码，在go里要显示的话，需要转换到utf8
}

type VBAProject struct {
	cf *compoundFile.CompoundFile

	Module []*ModuleInfo
	dic    map[string]int // 记录模块名称的下标
}

func init() {

}
func New() *VBAProject {
	return new(VBAProject)
}

// 解析
func (me *VBAProject) Parse(b []byte) (err error) {
	if me.cf, err = compoundFile.NewCompoundFile(b); err != nil {
		return
	}
	if err = me.cf.Parse(); err != nil {
		return
	}

	if err = me.getModuleInfo(); err != nil {
		return
	}

	return
}

// 获取所有模块的信息
func (me *VBAProject) getModuleInfo() (err error) {
	var b []byte
	var strPre string = ""

	if b, err = me.cf.GetStream(`VBA\dir`); err != nil {
		// 03版本的是放在_VBA_PROJECT_CUR\目录之下
		if b, err = me.cf.GetStream(`_VBA_PROJECT_CUR\VBA\dir`); err != nil {
			return errors.New("没有找到dir流。")
		} else {
			strPre = `_VBA_PROJECT_CUR\`
		}
	}
	// 解压缩dir流
	rle := rleVBA.NewRLE(b)
	b = rle.UnCompress()
	// 分析dir流
	if mi, err1 := vbaDir.GetModuleInfo(b); err1 != nil {
		return err1
	} else {
		count := len(mi)

		me.Module = make([]*ModuleInfo, count)
		me.dic = make(map[string]int)
		for i := range mi {
			me.Module[i] = new(ModuleInfo)
			me.Module[i].Name = mi[i].Name
			me.dic[me.Module[i].Name] = i

			if mi[i].ModuleType == 33 {
				// 标准模块
				me.Module[i].Type = ".bas"
			} else {
				me.Module[i].Type = ".cls"
			}
			// 解压模块代码
			if b, err = me.cf.GetStream(strPre + `VBA\` + me.Module[i].Name); err != nil {
				return
			}
			rle = rleVBA.NewRLE(b[mi[i].TextOffset:])
			me.Module[i].Code = rle.UnCompress()
		}
	}

	return
}
