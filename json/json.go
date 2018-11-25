// 将1个json字符输出成1个结构体

package main

import (
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

	if err = jsonUnmarshal(b); err != nil {
		fmt.Println(err)
	}

	for i := range arrStruct {
		fmt.Println(arrStruct[i])
	}

}

func jsonUnmarshal(b []byte) (err error) {
	fmt.Printf("%s\r\n", string(b))

	d := make(map[string]interface{})
	if err := json.Unmarshal(b, &d); err != nil {
		return err
	} else {
		printOut(d)
	}
	return nil
}

func printOut(d map[string]interface{}) {
	var str string
	str = fmt.Sprintf("type s%d struct {\r\n", structCount)

	for k, v := range d {
		// 首字母必须大写才能解析
		b := []byte(k)
		if b[0] >= 'a' && b[0] <= 'z' {
			b[0] = b[0] + 'A' - 'a'
			k = string(b)
		}
		str += fmt.Sprint("\t", k)
		vType := reflect.TypeOf(v).String()
		//		fmt.Println(vType)

		switch vType {
		case "map[string]interface{}":
			// 新的1个结构体
			structCount++
			printOut(v.(map[string]interface{}))
			str += fmt.Sprintf(" s%d", structCount)
		case "[]interface {}":
			str += fmt.Sprint("\t", printOutArr(v.([]interface{}), "[]"))

		default:
			str += fmt.Sprint(" ", vType)
		}
		//		switch v.(type) {
		//		case string:
		//			fmt.Println(" ", v)
		//		case int:
		//			fmt.Println(" ", v)
		//		case float64:
		//			fmt.Println(" ", v)
		//		case map[string]interface{}:
		//			printOut(v.(map[string]interface{}))
		//		case []interface{}:
		//			fmt.Println(" ", v)

		//			vv := v.([]interface{})
		//			printOutArr(vv)
		//		}

		str += fmt.Sprintln("\t`json:" + `"` + k + "\"`")
	}

	str += "}\r\n"

	arrStruct = append(arrStruct, str)
}

// strPre	存在数组中的数组情况，所以可能有多个“[]”
func printOutArr(v []interface{}, strPre string) (ret string) {
	for i := range v {
		switch v[i].(type) {
		case []interface{}:
			// 数组中的数组
			strPre += "[]"

			vv := v[i].([]interface{})
			return printOutArr(vv, strPre)
		case map[string]interface{}:
			// 新的1个结构体
			structCount++
			printOut(v[i].(map[string]interface{}))
			return fmt.Sprintf("%ss%d", strPre, structCount)
		}
	}
	return
}
