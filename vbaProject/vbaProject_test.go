package vbaProject

import (
	"io/ioutil"
	"testing"
)

//func Test_GetModuleCode(t *testing.T) {
//	vp := New()
//	if b, err := ioutil.ReadFile("07vbaProject.bin"); err != nil {
//		fmt.Println(err)
//	} else {
//		if err := vp.Parse(b); err != nil {
//			fmt.Println(err)
//		} else {
//			if str, err := vp.GetModuleCode("模块1"); err != nil {
//				t.Error(err)
//			} else {
//				t.Log(str)
//			}
//		}
//	}
//}

//func Test_HideModule(t *testing.T) {
//	vp := New()
//	if b, err := ioutil.ReadFile("03.xls"); err != nil {
//		t.Error(err)
//	} else {
//		if err := vp.Parse(b); err != nil {
//			t.Error(err)
//		} else {
//			if err := vp.HideModule("模块1"); err != nil {
//				t.Error(err)
//			} else {
//				// 改写文件
//				if err := ioutil.WriteFile("03.xls", vp.cf.SrcByte, 0666); err != nil {
//					t.Error(err)
//				} else {
//					t.Log("隐藏成功")
//				}
//			}
//		}
//	}
//}

func Test_UnHideModule(t *testing.T) {
	vp := New()
	if b, err := ioutil.ReadFile("03.xls"); err != nil {
		t.Error(err)
	} else {
		if err := vp.Parse(b); err != nil {
			t.Error(err)
		} else {
			if err := vp.UnHideModule("模块1"); err != nil {
				t.Error(err)
			} else {
				// 改写文件
				if err := ioutil.WriteFile("03.xls", vp.cf.SrcByte, 0666); err != nil {
					t.Error(err)
				} else {
					t.Log("取消隐藏成功")
				}
			}
		}
	}
}

//func Test_UnProtectProject(t *testing.T) {
//	vp := New()
//	if b, err := ioutil.ReadFile("03.xls"); err != nil {
//		t.Error(err)
//	} else {
//		if err := vp.Parse(b); err != nil {
//			t.Error(err)
//		} else {
//			if err := vp.UnProtectProject(); err != nil {
//				t.Error(err)
//			} else {
//				// 改写文件
//				if err := ioutil.WriteFile("03.xls", vp.cf.SrcByte, 0666); err != nil {
//					t.Error(err)
//				} else {
//					t.Log("破解成功")
//				}
//			}
//		}
//	}
//}
