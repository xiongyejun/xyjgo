package main

type Configs struct {
	Cfgs []*Config
}

// 配置文件
type Config struct {
	Lng        string // 语言名称
	FuncSelect string // 函数声明语句

	// 1个Byte，2个Byte，4个Byte的类型名称
	B1 string
	B2 string
	B4 string
}

// Public Declare Function SetWindowsHookEx Lib "User32" Alias "SetWindowsHookExA" (ByVal idHook As Long, ByVal lpfn As LongPtr, ByVal hmod As LongPtr, ByVal dwThreadId As Long) As LongPtr
func VBAConfig() (ret *Config) {
	ret = new(Config)
	ret.Lng = "VBA"
	ret.FuncSelect = `'Public Declare IsFunction Function.Name Lib Function.FullName`

	ret.B1 = "Byte"
	ret.B2 = "Integer"
	ret.B4 = "Long"

	return
}
