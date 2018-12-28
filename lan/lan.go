// local area network
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

type datas struct {
	strSep     string
	uploadPath string
	sharedPath string
	imagesPath string
	musicPath  string
	bHtml      []byte
}

var d *datas

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func init() {
	d = new(datas)
	// 获取上传路径、分享路径、图片路径、音乐路径
	d.strSep = string(os.PathSeparator)
	d.uploadPath = os.Getenv("USERPROFILE") + d.strSep + "LVNupload" + d.strSep
	d.sharedPath = os.Getenv("USERPROFILE") + d.strSep + "Documents" + d.strSep
	d.imagesPath = os.Getenv("USERPROFILE") + d.strSep + "images" + d.strSep
	d.musicPath = os.Getenv("USERPROFILE") + d.strSep + "Music" + d.strSep
	// 如果上传路径、分享路径不存在就新建
	mkDir(d.uploadPath)
	mkDir(d.sharedPath)
}

// 新建文件夹
func mkDir(dirName string) (err error) {
	if _, err = os.Stat(dirName); err != nil {
		if err = os.Mkdir(dirName, 0666); err != nil {
			fmt.Println(err)
			return
		}
	}
	return nil
}

func main() {
	var err error
	// 读取要写入的html
	if d.bHtml, err = ioutil.ReadFile(os.Getenv("GOPATH") + `\src\github.com\xiongyejun\xyjgo\lan\temp.html`); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("请访问下面的链接:")
	showip()
	http.HandleFunc("/", uploadFileHandler)
	http.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir(d.sharedPath))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(d.imagesPath))))
	http.Handle("/Music/", http.StripPrefix("/Music/", http.FileServer(http.Dir(d.musicPath))))
	http.ListenAndServe(":8080", nil)
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, string(d.bHtml))
	fmt.Println("r.RemoteAddr=", r.RemoteAddr)

	if r.Method == "POST" {
		file, handler, err := r.FormFile("fileUpload") //name的字段
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		// 保存文件
		newFile, err := os.Create(d.uploadPath + handler.Filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer newFile.Close()
		// 文件太大时候，避免卡死，固定FILE_BYTES来读取
		const FILE_BYTES int = 1024 * 1024
		var n int = FILE_BYTES
		for n == FILE_BYTES {
			b := make([]byte, FILE_BYTES)
			n, err = file.Read(b)
			if err != nil && err != io.EOF {
				fmt.Println(err)
				return
			}
			_, err = newFile.Write(b)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		fmt.Println("upload successfully:" + d.uploadPath + handler.Filename)
		w.Write([]byte("<br>Success成功上传"))
	}
}

func showip() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String() + ":8080")
			}
		}
	}
}
