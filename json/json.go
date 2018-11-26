// 将1个json字符输出成1个结构体

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

var structCount int
var arrCount int
var arrStruct []string

func main() {
	if len(os.Args) == 1 {
		fmt.Println("json <jsonstring.txt>")
		return
	}

	var f *os.File
	var err error
	if f, err = os.Open(os.Args[1]); err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	var b []byte
	if b, err = ioutil.ReadAll(f); err != nil {
		fmt.Println(err)
		return
	}
	// 处理windows txt文件前面3个字节可能是0xEF, 0xBB, 0xBF的情况
	if bytes.HasPrefix(b, []byte{0xEF, 0xBB, 0xBF}) {
		b = b[3:]
	}

	if err = jsonUnmarshal(b); err != nil {
		fmt.Println(err)
		return
	}

	for i := range arrStruct {
		fmt.Println(arrStruct[i])
	}

}

func jsonUnmarshal(b []byte) (err error) {
	fmt.Printf("%s\r\n\r\n", string(b))

	d := make(map[string]interface{})
	if err := json.Unmarshal(b, &d); err != nil {
		return err
	} else {
		printOut(d)
	}
	return nil
}

// 根据json的值的类型，返回go结构体需要的字符串
func getType(v interface{}) string {
	vType := reflect.TypeOf(v).String()

	switch vType {
	case "map[string]interface {}":
		// 新的1个结构体
		structCount++
		printOut(v.(map[string]interface{}))
		return fmt.Sprintf("s%d", structCount)
	case "[]interface {}":
		return fmt.Sprint(printOutArr(v.([]interface{}), "[]"))

	default:
		return vType
	}
}

// 得到1个“map[string]interface{}”类型的结构体字符str，存放到arrStruct
func printOut(d map[string]interface{}) {
	var str string
	// 结构体的开头 type s0 struct {
	str = fmt.Sprintf("type s%d struct {\r\n", structCount)

	for k, v := range d {
		// 首字母必须大写才能解析
		b := []byte(k)
		var ucaseK string = k
		if b[0] >= 'a' && b[0] <= 'z' {
			b[0] = b[0] + 'A' - 'a'
			ucaseK = string(b)
		}
		// 结构体的属性名称
		str += fmt.Sprintf("\t%s ", ucaseK)
		// 获取结构体的属性的类型
		str += getType(v)
		// 为结构体添加tag
		str += fmt.Sprintln("\t`json:" + `"` + k + "\"`")
	}
	// 最后的"}"
	str += "}\r\n"
	// 完成1个结构体字符的构造，存放到arrStruct
	arrStruct = append(arrStruct, str)
}

// strPre	存在数组中的数组情况，所以可能有多个“[]”
func printOutArr(v []interface{}, strPre string) (ret string) {
	if len(v) == 0 {
		return " string"
	} else {
		return strPre + getType(v[0])
	}
}
