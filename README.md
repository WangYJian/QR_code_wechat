# QR_code_wechat
## 微信获取带参数二维码
### 1.导入包
#### go get github.com/WangYJian/QR_code_wechat
### 2.创建QRCode对象
#### QRCode, err := QR_code_wechat.NewQRCode("access_token")
### 3.生成对应用户的含参二维码，获取二维码media_id
#### media_id, err = QRCode.GenerateQRCode("user_id")