package QR_code_wechat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type QRCode struct {
	AccessToken string // 微信access_token
}

// GenerateQRCode 生成二维码，传入用户id，返回二维码的图片media_id
func (q *QRCode) GenerateQRCode(userId string) (string, error) {
	// 请求微信接口，生成二维码
	url := "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" + q.AccessToken
	jsonStr := []byte(`{"expire_seconds": 604800, "action_name": "QR_STR_SCENE", "action_info": {"scene": {"scene_str": "` + userId + `"}}}`)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 解析微信返回的json
	var result map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	// 获取ticket
	ticket := result["ticket"].(string)
	// 请求微信接口，获取二维码图片
	url = "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + ticket
	resp, err = http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 上传图片到微信服务器
	url = "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=" + q.AccessToken + "&type=image"
	resp, err = http.Post(url, "image/jpeg", resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 解析微信返回的json
	body, err = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	// 获取media_id
	media_id := result["media_id"].(string)
	return media_id, nil
}

func NewQRCode(accessToken string) *QRCode {
	return &QRCode{
		AccessToken: accessToken,
	}
}
