// local area network
package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

var strSep string
var uploadPath string
var sharedPath string

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func init() {
	strSep = string(os.PathSeparator)
	uploadPath = os.Getenv("USERPROFILE") + strSep + "LVNupload" + strSep
	sharedPath = os.Getenv("USERPROFILE") + strSep + "Documents" + strSep

	mkDir(uploadPath)
	mkDir(sharedPath)
}

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
	fmt.Println("请访问下面的链接:")
	showip()
	http.HandleFunc("/", uploadFileHandler)
	http.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir(sharedPath))))
	http.ListenAndServe(":8080", nil)
}
func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>局域网</title>
</head>
<body> 
    <h1>局域网</h1>
    <br>
    <br>
    <form action="UploadFile.ashx" method="post" enctype="multipart/form-data">
    <input type="file" name="fileUpload" />
    <input type="submit" name="上传文件">
    </form>
        <br>
    <br>
        <br>
    <br>
    <a href="/file">文件下载</a>
</body>
</html>
        `)
	if r.Method == "POST" {
		file, handler, err := r.FormFile("fileUpload") //name的字段
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		newFile, err := os.Create(uploadPath + handler.Filename)
		check(err)
		defer newFile.Close()

		const FILE_BYTES int = 1024
		var n int = FILE_BYTES
		for n == FILE_BYTES {
			b := make([]byte, FILE_BYTES)
			n, _ = file.Read(b)
			newFile.Write(b)
		}

		fmt.Println("upload successfully:" + uploadPath + handler.Filename)
		w.Write([]byte("SUCCESS"))
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
