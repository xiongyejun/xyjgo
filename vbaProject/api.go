// 对外的一些功能

package vbaProject

import (
	"bytes"
	"errors"
	"regexp"

	"github.com/axgle/mahonia"
)

// 获取某个模块的代码
func (me *VBAProject) GetModuleCode(moduleName string) (ret string, err error) {
	var index int
	if index, err = me.checkModuleExists(moduleName); err != nil {
		return
	}
	decoder := mahonia.NewDecoder("gbk")

	ret = decoder.ConvertString(string(me.Module[index].Code))
	return

}

// 获取模块在VBAProject.Module里的下标
func (me *VBAProject) GetModuleIndex(moduleName string) (ret int, err error) {
	return me.checkModuleExists(moduleName)
}

// 隐藏模块
func (me *VBAProject) HideModule(moduleName string) (err error) {
	var index int
	if index, err = me.checkModuleExists(moduleName); err != nil {
		return
	}
	if me.Module[index].Type == CLASS_MODULE {
		return errors.New("类模块不能隐藏")
	}

	var bProject []byte
	if bProject, err = me.cf.GetStream(me.VBA_PROJECT_CUR + `PROJECT`); err != nil {
		return
	}

	// 将Module=moduleName替换为空，再改写文件即可
	encoder := mahonia.NewEncoder("gbk")
	bReplace := []byte(encoder.ConvertString(`Module=` + moduleName))
	bReplace = append(bReplace, '\r', '\n')

	if bytes.Contains(bProject, bReplace) {
		// 改写文件
		newByte := bytes.Replace(bProject, bReplace, []byte(nil), -1)
		if err = me.cf.ModifyStream(me.VBA_PROJECT_CUR+`PROJECT`, newByte); err != nil {
			return
		}
	} else {
		return errors.New("PROJECT中没有找到：" + `Module=` + moduleName)
	}
	return
}

// 取消隐藏模块
func (me *VBAProject) UnHideModule(moduleName string) (err error) {
	var index int
	if index, err = me.checkModuleExists(moduleName); err != nil {
		return
	}
	if me.Module[index].Type == CLASS_MODULE {
		return errors.New("类模块是不能隐藏的")
	}

	// 读取PROJECT byte
	var bProject []byte
	if bProject, err = me.cf.GetStream(me.VBA_PROJECT_CUR + `PROJECT`); err != nil {
		return
	}

	// 判断是否是被隐藏了
	encoder := mahonia.NewEncoder("gbk")
	bModule := []byte(encoder.ConvertString(`Module=` + moduleName))
	bModule = append(bModule, '\r', '\n')
	if bytes.Contains(bProject, bModule) {
		return errors.New("模块并没有被隐藏。")
	}

	// HelpFile="" 在这个前面添加 Module=moduleNameODOA
	bOld := []byte(`HelpFile=""`)
	bNew := make([]byte, len(bModule)+len(bOld))
	copy(bNew[0:], bModule)
	copy(bNew[len(bModule):], bOld)

	newByte := bytes.Replace(bProject, bOld, bNew, 1)
	if err = me.cf.ModifyStream(me.VBA_PROJECT_CUR+`PROJECT`, newByte); err != nil {
		return
	}

	return
}

// 检查模块是否存在，返回模块所在me.Module的下标
func (me *VBAProject) checkModuleExists(moduleName string) (index int, err error) {
	var ok bool
	if index, ok = me.dic[moduleName]; ok {
		return
	} else {
		return -1, errors.New("不存在的模块名称：" + moduleName)
	}
}

// 清除vba工程密码
func (me *VBAProject) UnProtectProject() (err error) {
	var bProject []byte
	if bProject, err = me.cf.GetStream(me.VBA_PROJECT_CUR + `PROJECT`); err != nil {
		return
	}

	pattern := `CMG="[A-Z\d]+"\r\n|DPB="[A-Z\d]+"\r\n|GC="[A-Z\d]+"\r\n`

	var bMatch bool
	if bMatch, err = regexp.Match(pattern, bProject); !bMatch {
		return
	}

	var reg *regexp.Regexp
	if reg, err = regexp.Compile(pattern); err != nil {
		return
	}
	// 替换后的byte
	newByte := reg.ReplaceAll(bProject, []byte{})
	if err = me.cf.ModifyStream(me.VBA_PROJECT_CUR+`PROJECT`, newByte); err != nil {
		return
	}

	return
}
