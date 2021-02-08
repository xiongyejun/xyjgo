package oleXML

import (
	"errors"
	"strconv"

	"github.com/axgle/mahonia"
	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

type XML struct {
	unknown *ole.IUnknown

	DOMDocument *ole.IDispatch // iexplore.exe

	unknownXMLSchemaCache *ole.IUnknown
	XMLSchemaCache        *ole.IDispatch
}

func (me *XML) InitXML() (err error) {
	err = ole.CoInitialize(0)
	// TODO 多次调用ole.CoInitialize会出错：函数不正确。
	// 忽略它
	err = nil

	me.unknown, err = oleutil.CreateObject("MSXML2.DOMDocument.6.0")
	if err != nil {
		err = errors.New("oleutil.CreateObject\n" + err.Error())
		return
	}
	if me.DOMDocument, err = me.unknown.QueryInterface(ole.IID_IDispatch); err != nil {
		err = errors.New("me.unknown.QueryInterface\n" + err.Error())
		return
	}

	me.unknownXMLSchemaCache, err = oleutil.CreateObject("MSXML2.XMLSchemaCache.6.0")
	if err != nil {
		err = errors.New("oleutil.CreateObject\n" + err.Error())
		return
	}
	if me.XMLSchemaCache, err = me.unknownXMLSchemaCache.QueryInterface(ole.IID_IDispatch); err != nil {
		err = errors.New("me.unknownXMLSchemaCache.QueryInterface\n" + err.Error())
		return
	}

	if _, err = oleutil.PutProperty(me.DOMDocument, "Schemas", me.XMLSchemaCache); err != nil {
		err = errors.New("oleutil.PutProperty(Schemas)\n" + err.Error())
		return
	}

	return nil
}

func (me *XML) UnInit() {
	me.XMLSchemaCache.Release()
	me.DOMDocument.Release()

	me.unknown.Release()
	me.unknownXMLSchemaCache.Release()
	ole.CoUninitialize()
}

func (me *XML) Validate(bXMLUFT8 []byte, XSDFileName, SchemaURL string) (err error) {
	if bXMLUFT8[0] == 0xef && bXMLUFT8[1] == 0xbb && bXMLUFT8[2] == 0xbf {
		bXMLUFT8 = bXMLUFT8[3:]
	}
	encoder := mahonia.NewEncoder("gbk")
	bXMLUFT8 = []byte(encoder.ConvertString(string(bXMLUFT8)))

	if _, err = oleutil.CallMethod(me.XMLSchemaCache, "Add", SchemaURL, XSDFileName); err != nil {
		err = errors.New("oleutil.CallMethod(Add)\n" + err.Error())
		return
	}
	if _, err = oleutil.PutProperty(me.DOMDocument, "validateOnParse", true); err != nil {
		err = errors.New("oleutil.PutProperty(validateOnParse)\n" + err.Error())
		return
	}
	if _, err = oleutil.PutProperty(me.DOMDocument, "async", false); err != nil {
		err = errors.New("oleutil.PutProperty(async)\n" + err.Error())
		return
	}
	var v *ole.VARIANT
	if v, err = oleutil.CallMethod(me.DOMDocument, "LoadXML", string(bXMLUFT8)); err != nil {
		err = errors.New("oleutil.CallMethod(LoadXML)\n" + err.Error())
		return
	} else if !v.Value().(bool) {
		if v, err = oleutil.GetProperty(me.DOMDocument, "parseError"); err != nil {
			err = errors.New("oleutil.GetProperty(parseError)\n" + err.Error())
			return
		}
		tmp := v.ToIDispatch()
		defer tmp.Release()
		if v, err = oleutil.GetProperty(tmp, "reason"); err != nil {
			err = errors.New("oleutil.GetProperty(reason)\n" + err.Error())
			return
		}
		sTmp := v.ToString()
		if v, err = oleutil.GetProperty(tmp, "Line"); err != nil {
			err = errors.New("oleutil.GetProperty(Line)\n" + err.Error())
			return
		}
		err = errors.New(sTmp + "\nLine :" + strconv.Itoa(int(v.Val)))

		return
	}

	return
}

func New() *XML {
	return new(XML)
}
