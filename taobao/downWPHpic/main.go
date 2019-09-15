package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

const (
	品牌名称 int = iota
	商品名称
	产地
	材质
	是否有防盗扣
	洗涤说明
	配件备注
	商品编号

	productInfoCount
)

// 下载图片
// https://detail.vip.com/610530-80623383.html

func main() {
	var i int = 80623383

	for {
		fmt.Print("请输入8位编号：")
		if _, err := fmt.Scanln(&i); err != nil {
			fmt.Println(err)
			return
		}

		tmp := len(strconv.Itoa(i))
		if tmp == 8 {
			break
		} else {
			fmt.Printf("输入的编号%d是%d位\n\n", i, tmp)
		}
	}

	// fmt.Println(i)
	// fmt.Scanln(&i)
	// return

	strUrl := "https://detail.vip.com/610530-"

	for {
		if b, err := httpGet(strUrl + strconv.Itoa(i) + ".html"); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%d……\n", i)
			getResult(b)
		}

		i++
	}
}

func getResult(b []byte) {
	var re *regexp.Regexp
	var err error
	var expr string = `(?s)<table class="dc-table fst">.*?</table>`
	if re, err = regexp.Compile(expr); err != nil {
		fmt.Println(err)
		return
	}
	bb := re.FindSubmatch(b)
	if len(bb) == 0 {
		fmt.Println("err:没有读取到商品信息。\n")
		return
	}

	expr = `(?s)<th class="dc-table-tit">(.*?)</th>.*?<td>(.*?)</td>`
	if re, err = regexp.Compile(expr); err != nil {
		fmt.Println(err)
		return
	}
	bbb := re.FindAllSubmatch(bb[0], -1)
	if len(bbb) != productInfoCount {
		fmt.Printf("err:读取到的商品信息没有%d个。\n\n", productInfoCount)
		return
	}
	if string(bbb[商品编号][1]) != "商品编号：" {
		fmt.Printf("err:读取到的商品信息第%d个不是[商品编号]。\n\n", productInfoCount)
		return
	}

	// 建文件夹
	dirName := string(bbb[商品编号][2])
	if err = os.Mkdir(dirName, 0666); err != nil {
		fmt.Println(err, "\n")
		return
	}

	// 写入商品信息
	var f *os.File
	if f, err = os.OpenFile(dirName+`\商品信息.txt`, os.O_CREATE|os.O_APPEND, 0666); err != nil {
		fmt.Println(err, "\n")
		return
	}
	defer f.Close()
	for i := range bbb {
		if _, err = f.Write(bbb[i][1]); err != nil {
			fmt.Println(err, "\n")
			return
		}
		if _, err = f.Write(bbb[i][2]); err != nil {
			fmt.Println(err, "\n")
			return
		}
		if _, err = f.Write([]byte{0x0d, 0x0a}); err != nil {
			fmt.Println(err, "\n")
			return
		}
	}

	/*
	   <table class="dc-table fst">
	       <tbody>
	           <tr>
	                               <th class="dc-table-tit">品牌名称：</th>
	                   <td>RYMA</td>
	                                   <th class="dc-table-tit">商品名称：</th>
	                   <td>绿色秋冬保暖羊毛百搭简洁休闲裤</td>
	               </tr><tr>                    <th class="dc-table-tit">产地：</th>
	                   <td>中国</td>
	                                   <th class="dc-table-tit">材质：</th>
	                   <td>面料：49.5%羊毛 28.8%粘纤 21.7%聚酯纤维 里料：84.4%聚酯纤维 15.6%（含微量其他纤维）（装饰性除外）</td>
	               </tr><tr>                    <th class="dc-table-tit">是否有防盗扣：</th>
	                   <td>否</td>
	                                   <th class="dc-table-tit">洗涤说明：</th>
	                   <td>30度以下缓和水洗，不可漂白，悬挂晾干，中温熨烫，不可干洗，不宜氯漂</td>
	               </tr><tr>                    <th class="dc-table-tit">配件/备注：</th>
	                   <td>暂无配件</td>
	                                   <th class="dc-table-tit">商品编号：</th>
	                   <td>RT41300244</td>
	               </tr>        </tbody>
	   </table>
	*/

	// 下载图片
	// <a href="//a.vpimg3.com/upload/merchandise/vis/136/20021226-3/4/1438782-1.jpg" class="J-mer-bigImgZoom">
	// <a href="//a.vpimg3.com/upload/merchandise/pdc/216/745/7817480819492745216/0/RYMA-FI15400954-1.jpg" class="J-mer-bigImgZoom">
	expr = `<a href="//a\.vpimg3\.com/upload/merchandise/(.*?)\.jpg" class="J-mer-bigImgZoom">`
	if re, err = regexp.Compile(expr); err != nil {
		fmt.Println(err)
		return
	}
	bbb = re.FindAllSubmatch(b, -1)
	if len(bbb) == 0 {
		fmt.Printf("err: 没有读取到图片。\n\n")
		return
	}

	for i := range bbb {
		bUrl := bbb[i][1]
		index := bytes.LastIndex(bUrl, []byte("/"))
		if index != -1 {
			picName := string(bUrl[index+1:]) + ".jpg"
			if b, err = httpGet(`https://a.vpimg3.com/upload/merchandise/` + string(bUrl) + ".jpg"); err != nil {
				fmt.Println(err, "\n")
				return
			}
			// 保存图片
			ioutil.WriteFile(dirName+`\`+picName, b, 0666)
		}
	}

}

// func httpPost(url string, strPost string) (ret []byte, err error) {
// 	var resp *http.Response
// 	if resp, err = http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(strPost)); err != nil {
// 		return
// 	}
// 	defer resp.Body.Close()
// 	if ret, err = ioutil.ReadAll(resp.Body); err != nil {
// 		return
// 	}
// 	return
// }

func httpGet(url string) (ret []byte, err error) {
	var resp *http.Response
	if resp, err = http.Get(url); err != nil {
		return
	}
	defer resp.Body.Close()
	if ret, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	return
}
