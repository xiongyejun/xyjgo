// 对外的一些功能

package vbaProject

import (
	"errors"

	"github.com/axgle/mahonia"
)

// 获取某个模块的代码
func (me *VBAProject) GetModuleCode(moduleName string) (ret string, err error) {
	if index, ok := me.dic[moduleName]; ok {
		decoder := mahonia.NewDecoder("gbk")

		ret = decoder.ConvertString(string(me.Module[index].Code))
		return
	} else {
		return "", errors.New("不存在的模块名称：" + moduleName)
	}
}
