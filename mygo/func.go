package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// 判断所给路径文件/文件夹是否存在
func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func httpGet(url string) (ret []byte, err error) {
	var resp *http.Response
	if resp, err = http.Get(url); err != nil {
		return
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func getHtmlFromDivId(strHtml string, DivId string) (ret string, err error) {
	if DivId == "" {
		return
	}

	arr := strings.Split(strHtml, DivId)

	if len(arr) > 1 {
		ret = arr[1]
		arr = strings.Split(ret, `</div>`)
		ret = arr[0]
	} else {
		err = errors.New("不存在的divID")
	}
	return
}

func gbkToUtf8(b []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(b), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

type FileInfo struct {
	IsDir      bool
	Size       int64
	Path, Name string
	Level      int
	HasFile    bool

	Subs []*FileInfo
}

func scanDir(path, name string, iLevel int) (ret *FileInfo, err error) {
	var entrys []os.FileInfo
	if entrys, err = ioutil.ReadDir(path + name); err != nil {
		return
	}

	ret = &FileInfo{
		IsDir:   true,
		Size:    0,
		Path:    path,
		Name:    name,
		Level:   iLevel,
		HasFile: false,
	}

	for i := range entrys {
		if entrys[i].IsDir() {
			var tmp *FileInfo
			if tmp, err = scanDir(path+name+strPathSeparator, entrys[i].Name(), iLevel+1); err != nil {
				return
			}
			ret.Subs = append(ret.Subs, tmp)
			ret.Size += tmp.Size
		} else {
			ret.Size += entrys[i].Size()
			ret.HasFile = true
			ret.Subs = append(ret.Subs, &FileInfo{
				IsDir:   false,
				Size:    entrys[i].Size(),
				Path:    path + name + strPathSeparator,
				Name:    entrys[i].Name(),
				Level:   iLevel + 1,
				HasFile: false,
			})
		}
	}

	return
}
