# QR_code_wechat
## 微信获取带参数二维码
### 1.导入包
#### git get xxx
### 2.创建QRCode对象
#### QRCode := NewQRCode("access_token")
### 3.生成对应用户的含参二维码，获取二维码media_id
#### media_id = QRCode.GenerateQRCode("user_id")