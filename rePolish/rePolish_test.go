package rePolish

import (
	"testing"
)

func Test_func(t *testing.T) {
	var str []string = []string{
		"5",
		"+",
		"2", /*
			"*",
			"(",
			"3",
			"*",
			"(",
			"3",
			"-",
			"1",
			"*",
			"2",
			"+",
			"1",
			")",
			")",*/
	}
	t.Log(str)
	str, _ = RePolish(str)
	t.Log(str)

	t.Log(Calc(str))

}
