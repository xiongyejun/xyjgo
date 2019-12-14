// https://ai.baidu.com/docs#/OCR-API-AccurateBasic/top
package baidu

type BaiDu struct {
	URL          string
	access_token string
}

func (me *BaiDu) OCR(picPath string) (ret string, err error) {
	// 参数	是否必选	类型	可选值范围	说明
	// image	true	string	-	图像数据，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/jpeg/png/bmp格式
	// recognize_granularity	false	string	big、small	是否定位单字符位置，big：不定位单字符位置，默认值；small：定位单字符位置
	// detect_direction	false	string	true、false	是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括:
	// 												- true：检测朝向；
	// 												- false：不检测朝向。
	// vertexes_location	false	string	true、false	是否返回文字外接多边形顶点位置，不支持单字位置。默认为false
	// probability	false	string	true、false	是否返回识别结果中每一行的置信度

	return "baidu"
}

func New() (ret *BaiDu) {
	ret = new(BaiDu)
	ret.URL = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic"

	if ret.access_token, err = ret.get_access_token(); err != nil {
		return
	}
}

// access_token	通过API Key和Secret Key获取的access_token
func (me *BaiDu) get_access_token() (ret string, err error) {
	/*
		   获取Access Token
		   请求URL数据格式

		   向授权服务地址https://aip.baidubce.com/oauth/2.0/token发送请求（推荐使用POST），并在URL中带上以下参数：

		   grant_type： 必须参数，固定为client_credentials；
		   client_id： 必须参数，应用的API Key；
		   client_secret： 必须参数，应用的Secret Key；

			服务器返回的JSON文本参数如下：

			access_token： 要获取的Access Token；
			expires_in： Access Token的有效期(秒为单位，一般为1个月)；
			其他参数忽略，暂时不用;

	*/

	/*
	   若请求错误，服务器将返回的JSON文本包含以下参数：

	   error： 错误码；关于错误码的详细信息请参考下方鉴权认证错误码。
	   error_description： 错误描述信息，帮助理解和解决发生的错误。
	*/

	return nil
}
