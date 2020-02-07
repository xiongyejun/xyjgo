package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("dirJson.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var di dirInfos = dirInfos{}
	if err = json.Unmarshal(b, &di); err != nil {
		fmt.Println(err)
		return
	}
	di.downXS()
	for i := range di.DirInfos {
		<-ch
		i++
	}
	return
}
