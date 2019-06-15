// local area network
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
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
		if r.URL.String() == "/inputvalue.ashx" {
			fmt.Printf("%#v\r\n", r.PostForm)
		} else {
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
				fmt.Printf("IP:%s 子网掩码:%s\n", ipnet.IP.String()+":8080", ipnet.Mask)
			}

			if ipnet.IP.String() == "192.168.1.5" {
				tmp := Table(ipnet)
				for i := range tmp {
					fmt.Println(i, tmp[i])
				}
			}
		}
	}
}

// https://studygolang.com/articles/11517
// 算出内网IP范围
type IP uint32

func Table(ipnet *net.IPNet) (ret []net.IP) {
	ip := ipnet.IP.To4()
	fmt.Println("本机ip:", ip)
	var min, max IP
	for i := 0; i < 4; i++ {
		b := IP(ip[i] & ipnet.Mask[i])
		min += b << ((3 - uint(i)) * 8)
	}
	one, _ := ipnet.Mask.Size()
	max = min | IP(math.Pow(2, float64(32-one))-1)
	fmt.Printf("内网IP范围：%d --- %d\n", min, max)
	// max 是广播地址
	// i & 0x0000 00ff ==0  是为段位0的ip，根据RFC的规定，忽略
	for i := min; i < max; i++ {
		if i&0x000000ff == 0 {
			continue
		}
		ret = append(ret, UInt32ToIP(i))
	}
	return
}

func UInt32ToIP(intIP IP) net.IP {
	var bytes [4]byte
	bytes[0] = byte(intIP & 0xFF)
	bytes[1] = byte((intIP >> 8) & 0xFF)
	bytes[2] = byte((intIP >> 16) & 0xFF)
	bytes[3] = byte((intIP >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}
