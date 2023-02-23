package util

import (
	"fmt"
	"net/smtp"
)

// EmailConfig 结构体表示邮件配置信息
type EmailConfig struct {
	Host     string // SMTP服务器主机名
	Port     int    // SMTP服务器端口号
	Username string // SMTP服务器登录用户名
	Password string // SMTP服务器登录密码
	From     string // 发件人
	To       string // 收件人（多个收件人用逗号分隔）
}

// EmailUtil 结构体表示一个发送邮件的工具
type EmailUtil struct {
	config *EmailConfig // 邮件配置信息
}

// NewEmailUtil 函数创建一个邮件工具实例
func NewEmailUtil(config *EmailConfig) *EmailUtil {
	return &EmailUtil{
		config: config,
	}
}

// Send 方法发送邮件
func (u *EmailUtil) Send(subject string, content string) error {
	auth := smtp.PlainAuth("", u.config.Username, u.config.Password, u.config.Host)
	to := []string{}
	if len(u.config.To) > 0 {
		to = append(to, u.config.To)
	}

	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", u.config.To, subject, content))
	addr := fmt.Sprintf("%s:%d", u.config.Host, u.config.Port)
	err := smtp.SendMail(addr, auth, u.config.From, to, msg)
	if err != nil {
		return err
	}
	return nil
}
