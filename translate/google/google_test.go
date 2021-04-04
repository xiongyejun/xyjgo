package google

import (
	"testing"
)

func Test_func(t *testing.T) {
	if g, err := New(); err != nil {
		t.Error(err)
	} else {
		if ret, err := g.Translate(`misuse of aggregate`); err != nil {
			// ReplaceAll+returns+a+copy+of+the+slice+s+with+all++non-overlapping+instances+of+old+replaced+by+new.++If+old+is+empty%2C+it+matches+at+the+beginning+of+the+slice++and+after+each+UTF-8+sequence%2C+yielding+up+to+k%2B1+replacements++for+a+k-rune+slice.
			t.Error(err)
		} else {
			t.Log(ret)
		}
	}
}
