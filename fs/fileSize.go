// file size
// 读取某个文件夹下所有文件的大小

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	cp "github.com/xiongyejun/xyjgo/colorPrint"
)

type FileInfo struct {
	IsDir      bool
	Size       int64
	Path, Name string
	Level      int

	Subs []*FileInfo
}

var strPathSeparator = string(os.PathSeparator)
var fi *FileInfo = new(FileInfo)
var bDir bool = false

func main() {
	if len(os.Args) == 1 {
		fmt.Println("fileSize <dir> [d-打印文件夹;10-打印前10个文件夹（每个文件夹前10）]")
		return
	}

	path := os.Args[1]
	var err error
	if path, err = filepath.Abs(path); err != nil {
		fmt.Println(err)
		return
	}
	index := strings.LastIndex(path, strPathSeparator)
	name := path[index+1:]
	path = path[:index] + strPathSeparator

	if fi, err = scanDir(path, name, 0); err != nil {
		fmt.Println(err)
		return
	}

	fi.Sort()

	if len(os.Args) == 3 {
		switch os.Args[2] {
		case "10":
			fi.Print10()
		case "d":
			bDir = true
			fi.Print()
		default:
			fmt.Println("未知的第3参数。")
		}

	} else {
		fi.Print()
	}
}

func scanDir(path, name string, iLevel int) (ret *FileInfo, err error) {
	var entrys []os.FileInfo
	if entrys, err = ioutil.ReadDir(path + name); err != nil {
		return
	}

	ret = &FileInfo{
		IsDir: true,
		Size:  0,
		Path:  path,
		Name:  name,
		Level: iLevel,
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
			ret.Subs = append(ret.Subs, &FileInfo{
				IsDir: false,
				Size:  entrys[i].Size(),
				Path:  path + name + strPathSeparator,
				Name:  entrys[i].Name(),
				Level: iLevel + 1,
			})
		}
	}

	return
}

var iDir int = 0

func (me *FileInfo) Print10() {
	me.printDir()

	var iFile int = 0
	for i := range me.Subs {
		if !me.Subs[i].IsDir {
			me.Subs[i].printFile()
			iFile++
			if iFile == 10 {
				break
			}
		}
	}

	for i := range me.Subs {
		if me.Subs[i].IsDir {
			me.Subs[i].Print10()
			iDir++
			if iDir == 10 {
				os.Exit(0)
			}
		}
	}

}

func (me *FileInfo) Print() {
	me.printDir()

	if !bDir {
		for i := range me.Subs {
			if !me.Subs[i].IsDir {
				me.Subs[i].printFile()
			}
		}
	}

	for i := range me.Subs {
		if me.Subs[i].IsDir {
			me.Subs[i].Print()
		}
	}

}

func (me *FileInfo) printDir() {
	cp.SetColor(cp.Black, cp.Yellow)
	me.print()
	cp.ReSetColor()
	fmt.Println()
}
func (me *FileInfo) printFile() {
	me.print()
	fmt.Println()
}

func (me *FileInfo) print() {
	fmt.Printf("%10s |%2d  %s", strSize(me.Size), me.Level, strings.Repeat("  ", me.Level)+me.Name)
}

func (me *FileInfo) Sort() {
	for i := range me.Subs {
		if me.Subs[i].IsDir {
			me.Subs[i].Sort()
		}
	}
	sort.Sort(me)
}
func (me *FileInfo) Less(i, j int) bool {
	if me.Subs[i].IsDir == me.Subs[j].IsDir {
		return me.Subs[i].Size > me.Subs[j].Size
	}
	return false
}
func (me *FileInfo) Swap(i, j int) {
	me.Subs[i], me.Subs[j] = me.Subs[j], me.Subs[i]
}
func (me *FileInfo) Len() int {
	return len(me.Subs)
}

type size struct {
	value int64
	unit  string
}

var ssize []size = []size{
	size{1, "B"},
	size{1024, "K"},
	size{1024 * 1024, "M"},
	size{1024 * 1024 * 1024, "G"},
}

func strSize(size int64) (ret string) {
	var index int
	for index < len(ssize) && ssize[index].value <= size {
		index++
	}
	if index > 0 {
		index--
	}

	f := float64(size) / float64(ssize[index].value)
	ret = strconv.FormatFloat(f, 'f', 2, 64) + ssize[index].unit

	return
}
