package pdf

import (
	"errors"
	"regexp"
)

//-----------------------------------
// 3 0 obj <<

// /Type /Catalog

// /Pages 1 0 R
// >> endobj
//-----------------------------------

// 从obj中找到 /str # # R，返回x
func findR(b []byte, str string) (obj string, err error) {
	var re *regexp.Regexp
	var expr string = `/` + str + ` (\d{1,}) \d{1,} R`
	if re, err = regexp.Compile(expr); err != nil {
		return
	}
	bb := re.FindSubmatch(b)
	if bb == nil {
		return "", errors.New("没有找到[/" + str + " # # R]")
	}

	return string(bb[1]), nil
}
