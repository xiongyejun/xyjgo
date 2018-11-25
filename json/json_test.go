package main

import (
	"fmt"
	"testing"
)

func Test_print(t *testing.T) {
	jsonUnmarshal([]byte(`{"type":"ZH_CN2EN","errorCode":0,"elapsedTime":0,"translateResult":[[{"src":"你好","tgt":"hello"}]]}`))

	for i := range arrStruct {
		fmt.Println(arrStruct[i])
	}
}
