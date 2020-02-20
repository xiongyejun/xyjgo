package main

import (
	"fmt"
	"regexp"
	"testing"
)

func Test_func(t *testing.T) {
	var err error
	err = testRegexp()

	if err != nil {
		t.Error(err)
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
