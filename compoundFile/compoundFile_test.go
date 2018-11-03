package compoundFile

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_ModifyStream(t *testing.T) {
	if b, err := ioutil.ReadFile("03.xls"); err != nil {
		t.Error(err)
		return
	} else {
		if cf, err := NewCompoundFile(b); err != nil {
			t.Error(err)
			return
		} else {
			if err := cf.Parse(); err != nil {
				t.Error(err)
				return
			} else {
				if b, err := cf.GetStream(`_VBA_PROJECT_CUR\PROJECT`); err != nil {
					t.Error(err)
					return
				} else {
					b = bytes.Replace(b, []byte("Document="), []byte("123456789"), 99)

					if err := cf.ModifyStream(`_VBA_PROJECT_CUR\PROJECT`, b); err != nil {
						t.Error(err)
						return
					}
				}
			}
		}
	}
}

//func Test_PrintOut(t *testing.T) {
//	if b, err := ioutil.ReadFile("03.xls"); err != nil {
//		t.Error(err)
//		return
//	} else {
//		if cf, err := NewCompoundFile(b); err != nil {
//			t.Error(err)
//			return
//		} else {
//			if err := cf.Parse(); err != nil {
//				t.Error(err)
//				return
//			} else {
//				cf.PrintOut()
//			}

//		}
//	}
//}

//func Test_GetStream(t *testing.T) {
//	if b, err := ioutil.ReadFile("03.xls"); err != nil {
//		t.Error(err)
//		return
//	} else {
//		if cf, err := NewCompoundFile(b); err != nil {
//			t.Error(err)
//			return
//		} else {
//			if err := cf.Parse(); err != nil {
//				t.Error(err)
//				return
//			} else {
//				if b, err := cf.GetStream(`_VBA_PROJECT_CUR\PROJECT`); err != nil {
//					t.Error(err)
//					return
//				} else {
//					t.Logf("%s\r\n", b)
//				}
//			}

//		}
//	}
//}
