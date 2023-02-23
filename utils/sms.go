package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// SmsConfig 结构体表示短信配置信息
type SmsConfig struct {
	URL      string `json:"url"`       // API请求地址
	AppID    string `json:"app_id"`    // 应用ID
	AppKey   string `json:"app_key"`   // 应用Key
	SignName string `json:"sign_name"` // 签名名称
	Template string `json:"template"`  // 模板ID
}

// SmsUtil 结构体表示一个发送短信的工具
type SmsUtil struct {
	config *SmsConfig // 短信配置信息
}

// NewSmsUtil 函数创建一个短信工具实例
func NewSmsUtil(config *SmsConfig) *SmsUtil {
	return &SmsUtil{
		config: config,
	}
}

// Send 方法发送短信
func (u *SmsUtil) Send(subject string, content string) error {
	type SmsRequest struct {
		PhoneNumbers  string                 `json:"phone_numbers"`
		SignName      string                 `json:"sign_name"`
		TemplateCode  string                 `json:"template_code"`
		TemplateParam map[string]interface{} `json:"template_param"`
	}
	req := &SmsRequest{
		PhoneNumbers: u.config.URL,
		SignName:     u.config.SignName,
		TemplateCode: u.config.Template,
		TemplateParam: map[string]interface{}{
			"subject": subject,
			"content": content,
			"time":    time.Now().Format("2006-01-02 15:04:05"),
		},
	}
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := http.Post(u.config.URL, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", resp.StatusCode)
	}
	return nil
}
