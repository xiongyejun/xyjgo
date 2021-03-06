package compoundFile

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/xiongyejun/xyjgo/colorPrint"
)

// 获取cfStream的byte
func (me *CompoundFile) GetStream(streamPath string) (b []byte, err error) {
	var arr []string = strings.Split(streamPath, `\`)
	var storagePath string
	if len(arr) > 1 {
		storagePath = strings.Join(arr[:len(arr)-1], `\`)
	}

	var s *Storage
	if s, err = me.getStorage(storagePath); err != nil {
		return
	}

	if index, ok := s.streamDic[arr[len(arr)-1]]; !ok {
		return nil, errors.New("不存在的Stream名称。" + arr[len(arr)-1])
	} else {
		// 读取stream是按照512的大小读取的，但最后1个可能没有512
		return s.Streams[index].stream.Bytes()[:s.streamSize[index]], nil
	}
}

// 获取cfStream
func (me *CompoundFile) getStreamItem(streamPath string) (ret *cfStream, err error) {
	var arr []string = strings.Split(streamPath, `\`)
	var storagePath string
	if len(arr) > 1 {
		storagePath = strings.Join(arr[:len(arr)-1], `\`)
	}

	var s *Storage
	if s, err = me.getStorage(storagePath); err != nil {
		return
	}

	if index, ok := s.streamDic[arr[len(arr)-1]]; !ok {
		return nil, errors.New("不存在的Stream名称:" + arr[len(arr)-1])
	} else {
		// 读取stream是按照512的大小读取的，但最后1个可能没有512
		return s.Streams[index], nil
	}
}

// 获取Storage
func (me *CompoundFile) getStorage(storagePath string) (ret *Storage, err error) {
	var s *Storage = me.Root
	if storagePath == "" {
		return s, nil
	}
	var arr []string = strings.Split(storagePath, `\`)

	for i := 0; i < len(arr); i++ {
		if index, ok := s.storageDic[arr[i]]; !ok {
			return nil, errors.New("不存在的路径:" + arr[i])
		} else {
			s = s.Storages[index]
		}
	}

	return s, nil
}

func (me *CompoundFile) PrintOut() {
	me.printOut(me.Root, "")
}

func (me *CompoundFile) printOut(s *Storage, strPre string) {
	colorPrint.SetColor(colorPrint.White, colorPrint.DarkYellow)
	fmt.Printf("%s%s [Storage]\r\n", strPre, getPrintString(s.dir.name))
	colorPrint.ReSetColor()

	for i := range s.Storages {
		//		fmt.Printf("%s%s\r\n", strPre+"--", s.Storages[i].dir.name)
		me.printOut(s.Storages[i], strPre+"----")
	}

	for i := range s.Streams {
		fmt.Printf("%s%s [Stream]\r\n", strPre+"----", getPrintString(s.Streams[i].name))
	}
}

func getPrintString(str string) string {
	var strRet string

	for _, r := range []rune(str) {
		if !unicode.IsPrint(r) {
			strRet = strRet + "[" + strconv.Itoa(int(r)) + "]"
		} else {
			strRet += string(r)
		}
	}

	return strRet
}

// 释放
func (me *CompoundFile) Release(bOffset int, strExt string) (err error) {
	if err = me.Root.release("", bOffset, strExt); err != nil {
		return
	}
	return nil
}

// bOffset	字节的偏移
// strExt	保存流的时候添加的后缀
func (me *Storage) release(path string, bOffset int, strExt string) (err error) {
	path += me.dir.name
	path += `\`

	if err = os.Mkdir(path, os.ModePerm); err != nil {
		return
	}

	for i := range me.Streams {
		if me.Streams[i].stream != nil {
			if err = ioutil.WriteFile(path+getPrintString(me.Streams[i].name)+strExt, me.Streams[i].stream.Bytes()[bOffset:], os.ModePerm); err != nil {
				return
			}
		}
	}

	for i := range me.Storages {
		if err = me.Storages[i].release(path, bOffset, strExt); err != nil {
			return
		}
	}

	return nil
}
