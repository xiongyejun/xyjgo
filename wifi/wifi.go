package wifi

import (
	"errors"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

var (
	platform = runtime.GOOS
)

// 获取连接的wifi SSID
func GetSSID() (ssid string, err error) {
	if platform == "darwin" {
		path := "/System/Library/PrivateFrameworks/Apple80211." +
			"framework/Versions/Current/Resources/airport"

		if _, err = os.Stat(path); err != nil {
			return "", err
		}

		var b []byte
		if b, err = exec.Command(path, "-I").CombinedOutput(); err != nil {
			return "", err
		}

		str := string(b)
		regex, _ := regexp.Compile(" SSID: (.*)")
		match := regex.FindStringSubmatch(str)

		if len(match) < 2 {
			return "", errors.New("SSID: not found")
		}

		return match[1], nil
	}

	if platform == "windows" {
		var b []byte
		if b, err = exec.Command("cmd", "/C",
			"CHCP 65001 && netsh wlan show interfaces | findstr SSID").CombinedOutput(); err != nil {
			return "", err
		}

		str := string(b)

		regex, _ := regexp.Compile(` SSID *?: (.*)`)
		match := regex.FindStringSubmatch(str)
		if len(match) < 2 {
			return "", errors.New("SSID: Not found")
		}

		return strings.TrimSpace(match[1]), nil
	}

	return "", errors.New("SSID: not found")
}

// 根据ssid获取wifi密码
func GetPsw(ssid string) (psw string, err error) {
	if ssid == "" {
		return "", errors.New("SSID: not found")
	}

	if platform == "darwin" {
		cmd := exec.Command("security", "find-generic-password", "-l", ssid,
			"-D", "AirPort network password", "-w")

		var b []byte
		if b, err = cmd.CombinedOutput(); err != nil {
			return "", errors.New("permission denied")
		}
		return strings.TrimSpace(string(b)), nil
	}

	if platform == "windows" {
		var b []byte

		if b, err = exec.Command("cmd", "/C", `CHCP 65001 && netsh wlan show profile name="`+ssid+
			`" key=clear | findstr Key`).CombinedOutput(); err != nil {
			return "", err
		}

		str := string(b)
		regex, _ := regexp.Compile(` Content\s*:\s?(.*)`)
		match := regex.FindStringSubmatch(str)

		if len(match) < 2 {
			return "", errors.New("password: Not found")
		}

		return strings.TrimSpace(match[1]), nil
	}

	return "", errors.New("password: Not found")
}

//  返回二维码需要的字符串
func QRCodeFormat(account, psd string) (ret string) {
	return "WIFI:T:WPA;S:" + account + ";P:" + psd + ";;"
}
