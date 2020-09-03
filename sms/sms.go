// 短信服务
// 抓取https://www.materialtools.com的信息

package sms

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

type SMS struct {
	Page  int
	Infos []*Info
}

type Info struct {
	// 86中国、
	EM    string
	Phone string
	// 信息的url地址
	SMSContent string
}

type Msg struct {
	NO      string
	From    string
	Content string
	Time    string
}

// 获取短信
func GetMsg(SMSAddrNum string) (ret []*Msg, err error) {
	var b []byte
	if b, err = getSrc("https://www.materialtools.com/SMSContent/" + SMSAddrNum); err != nil {
		return
	}

	// if b, err = ioutil.ReadFile("sms.txt"); err != nil {
	// 	return
	// }

	expr := `(?s)<tr>.*?<td>(\d{1,2})</td>.*?<td>(.*?)</td>.*?<td>(.*?)</td>.*?<td>.*?<time>(.*?)</time>.*?</td>.*?</tr>`

	var reg *regexp.Regexp
	if reg, err = regexp.Compile(expr); err != nil {
		return
	}

	bbb := reg.FindAllSubmatch(b, -1)
	ret = make([]*Msg, len(bbb))
	for i := range bbb {
		ret[i] = new(Msg)
		ret[i].NO = string(bbb[i][1])
		ret[i].From = string(bbb[i][2])
		ret[i].Content = string(bbb[i][3])
		ret[i].Time = string(bbb[i][4])
	}

	return
}

// 获取所有的电话信息
func GetInfos() (ret []*SMS, err error) {
	var b []byte

	ret = make([]*SMS, 13)

	for i := 1; i < 14; i++ {
		ret[i-1] = new(SMS)
		ret[i-1].Page = i

		//网抓会出现 正在检测用户环境！！！！！！！！
		if b, err = ioutil.ReadFile(strconv.Itoa(i) + ".html"); err != nil {
			return
		}

		if ret[i-1].Infos, err = getInfo(b); err != nil {
			return
		}
	}

	return
}

// 从网页数据中提取电话信息
func getInfo(b []byte) (ret []*Info, err error) {
	expr := `(?s)<small><em>\+(\d{2,3})</em></small>.*?<h3>(\d{8,11})</h3>.*?href=".*?/SMSContent/(\d{1,10})"`

	var reg *regexp.Regexp
	if reg, err = regexp.Compile(expr); err != nil {
		return
	}

	bbb := reg.FindAllSubmatch(b, -1)
	ret = make([]*Info, len(bbb))
	for i := range bbb {
		ret[i] = new(Info)
		ret[i].EM = string(bbb[i][1])
		ret[i].Phone = string(bbb[i][2])
		ret[i].SMSContent = string(bbb[i][3])
	}
	return
}

func GetScrHtml(savepath string) (err error) {
	var b []byte
	for i := 12; i < 13; i++ {
		if b, err = getSrc("https://www.materialtools.com/?page=" + strconv.Itoa(i)); err != nil {
			return
		}

		if err = ioutil.WriteFile(savepath+strconv.Itoa(i)+".html", b, 0666); err != nil {
			return
		}
	}

	return
}

func getSrc(url string) (b []byte, err error) {
	var resp *http.Response
	if resp, err = http.Get(url); err != nil {
		return
	}
	defer resp.Body.Close()

	if b, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	return
}

func SaveToJson(fileName string, infos []*SMS) (err error) {
	var b []byte
	if b, err = json.MarshalIndent(&infos, "", "\t"); err != nil {
		return
	}

	if err = ioutil.WriteFile(fileName, b, 0666); err != nil {
		return
	}
	return
}
