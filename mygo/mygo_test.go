package main

import (
	"fmt"
	"regexp"
	"testing"
)

func Test_func(t *testing.T) {
	var fi *FileInfo
	var err error
	if fi, err = scanDir(`C:\Users\Administrator\Desktop\srcHtml\bcgm本草纲目\`, "srcHtml", 0); err != nil {
		t.Error(err)
		return
	}

	t.Log(fi.HasFile, fi.IsDir)
	for i := range fi.Subs {
		t.Logf("%d %s\n", i, fi.Subs[i].Name)
		t.Log(fi.Subs[i].HasFile)
		//		for j := range fi.Subs[i].Subs {
		//			t.Logf("%d %s\n", j, fi.Subs[i].Subs[j].Name)
		//		}
	}
}

func testRegexp() (err error) {
	var exp string = `(?i)(落.*?霞.*?l.*?u.*?o.*?.*?c.*?o.*?m)`
	var str string = `</p><p> 🌵 落+霞-小+說 L U ox i a - c o m +</p>`
	var reg *regexp.Regexp
	if reg, err = regexp.Compile(exp); err != nil {
		return
	}
	str = reg.ReplaceAllString(str, "")
	fmt.Println(str)
	return
}
