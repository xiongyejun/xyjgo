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

// {
//    "Lng": "VBA32",
//    "FuncSql": "select \ncase Function.IsFunction\n\twhen 1 then 'Public Declare Function '\n\telse 'Public Declare Sub '\nend\n\n|| Function.Name || ' Lib \"' || Dll.Name || '.dll\"' || \n\t\t\t\ncase Function.Alias\n\twhen '' then ' '\n\telse ' Alias \"' || Function.Alias || '\" '\n\tend \n\n|| Function.FullName\nFrom Function,Dll where Function.DllID=Dll.ID",
//    "ByVal": "ByVal",
//    "ByRef": "ByRef",
//    "B1": "As Byte",
//    "B2": "As integer",
//    "B4": "As Long",
//    "StructSql": "select 'Type ' || Name || '\n' || FullName || '\nEnd Type' from Struct where ",
//    "ConstSql": "select 'Const ' || Name || ' = ' || Value from Constant where ",
//    "Hex": "\u0026H"
//   }
