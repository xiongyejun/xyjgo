// 文件的sha1 hash
package fileHash

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
)

type FileInfo struct {
	Name     string
	FullName string
	Hash     string
}

type DirInfo struct {
	Path      string
	FilesInfo []FileInfo
	// Hash对应的FilesInfo下标
	MHashIndex map[string]int
}

var strSep string = string(os.PathSeparator)

func New(DirPath string) (ret *DirInfo) {
	if DirPath[len(DirPath)-1:] != strSep {
		DirPath += strSep
	}

	ret = new(DirInfo)
	ret.Path = DirPath
	ret.MHashIndex = make(map[string]int)

	return
}

func (me *DirInfo) GetFilesInfo() (err error) {
	var entrys []os.FileInfo
	if entrys, err = ioutil.ReadDir(me.Path); err != nil {
		return
	}

	for i := range entrys {
		if !entrys[i].IsDir() {
			tmp := FileInfo{}
			tmp.Name = entrys[i].Name()
			tmp.FullName = me.Path + tmp.Name
			if err = tmp.GetHash(); err != nil {
				return
			}

			me.MHashIndex[tmp.Hash] = len(me.FilesInfo)
			me.FilesInfo = append(me.FilesInfo, tmp)
		}
	}

	return
}

func (me *FileInfo) GetHash() (err error) {
	var f *os.File
	if f, err = os.Open(me.FullName); err != nil {
		return
	}
	defer f.Close()

	h := sha1.New()
	FILE_BYTES := 1024 * 1024
	var n int = FILE_BYTES
	var b []byte = make([]byte, FILE_BYTES)
	for n == FILE_BYTES {
		n, err = f.Read(b)
		if err != nil && err != io.EOF {
			return
		}
		_, err = h.Write(b[:n])
		if err != nil {
			return
		}
	}

	me.Hash = hex.EncodeToString(h.Sum(nil))

	return
}
