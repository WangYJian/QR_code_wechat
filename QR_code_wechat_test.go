package QR_code_wechat

import (
	"os"
	"testing"
)

func TestGenerateQRCode(t *testing.T) {
	qr, err := NewQRCode(os.Args[1])
	if err != nil {
		t.Errorf("NewQRCode() error = %v", err)
	}
	got, err := qr.GenerateQRCode("123456")
	if err != nil {
		t.Errorf("GenerateQRCode() error = %v", err)
	}
	if got == "" {
		t.Errorf("GenerateQRCode() got empty string")
	} else {
		t.Logf("GenerateQRCode() got %s", got)
	}
}
