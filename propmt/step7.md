请继续编写：service/service.go: 服务逻辑文件，包含服务的核心逻辑，如发送邮件和短信进行提醒等。

以下是一个简单的service.go文件示例，包含了服务的核心逻辑：

go
Copy code
package service

import (
	"fmt"
	"github.com/aaronpeng2020/gonotify/platform"
	"github.com/aaronpeng2020/gonotify/util"
	"strings"
)

// Service 结构体表示一个服务
type Service struct {
	platforms []platform.Platform // 平台列表
	emailUtil *util.EmailUtil     // 邮件工具
	smsUtil   *util.SmsUtil       // 短信工具
	balance   map[string]float64  // 平台余额信息
	warn      float64             // 预警值
}

// NewService 函数创建一个服务实例
func NewService(platforms []platform.Platform, emailUtil *util.EmailUtil, smsUtil *util.SmsUtil, warn float64) *Service {
	return &Service{
		platforms: platforms,
		emailUtil: emailUtil,
		smsUtil:   smsUtil,
		balance:   make(map[string]float64),
		warn:      warn,
	}
}

// CheckBalance 方法检查平台余额并更新余额信息
func (s *Service) CheckBalance() {
	for _, p := range s.platforms {
		balance, err := p.Balance()
		if err != nil {
			fmt.Printf("Failed to get balance from platform %s: %v\n", p.Type(), err)
			continue
		}
		fmt.Printf("Got balance %.2f from platform %s\n", balance, p.Type())
		s.UpdateBalance(p.Type(), balance)
	}
}

// UpdateBalance 方法更新平台余额信息
func (s *Service) UpdateBalance(platformType string, balance float64) {
	s.balance[platformType] = balance
}

// GetBalance 方法获取指定平台的余额信息
func (s *Service) GetBalance(platformType string) float64 {
	return s.balance[platformType]
}

// SendNotification 方法根据平台余额信息发送提醒
func (s *Service) SendNotification() {
	var content strings.Builder

	for _, p := range s.platforms {
		balance := s.balance[p.Type()]
		if balance < s.warn {
			content.WriteString(fmt.Sprintf("Platform %s balance is %.2f, less than warn value %.2f.\n", p.Type(), balance, s.warn))
		}
	}

	if content.Len() > 0 {
		err := s.emailUtil.Send("GoNotify - Balance Warning", content.String())
		if err != nil {
			fmt.Printf("Failed to send email: %v\n", err)
		} else {
			fmt.Println("Email sent.")
		}

		err = s.smsUtil.Send("GoNotify", content.String())
		if err != nil {
			fmt.Printf("Failed to send SMS: %v\n", err)
		} else {
			fmt.Println("SMS sent.")
		}
	} else {
		fmt.Println("No need to send notification.")
	}
}
该文件定义了一个Service结构体，包含了服务的核心逻辑。具体来说，该结构体包含了平台列表、邮件工具、短信工具、平台余额信息和预警值等信息。CheckBalance方法用于检查平台余额并更新余额信息，UpdateBalance方法用于更新平台余额信息，GetBalance方法用于获取指定平台的余额信息，SendNotification方法用于根据