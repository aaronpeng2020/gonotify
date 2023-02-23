请继续编写schedule/schedule.go文件，实现功能：
定时任务调度文件，用于定期请求平台余额信息。
以下是一个简单的schedule.go文件示例，用于定时请求平台余额信息：

go
Copy code
package schedule

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/xxxxxx/GoNotify/platform"
	"github.com/xxxxxx/GoNotify/service"
)

// Scheduler 结构体表示一个定时任务调度器
type Scheduler struct {
	srv       *service.Service
	platforms []platform.Platform
}

// NewScheduler 函数创建一个定时任务调度器
func NewScheduler(srv *service.Service, platforms []platform.Platform) *Scheduler {
	return &Scheduler{
		srv:       srv,
		platforms: platforms,
	}
}

// Start 方法启动定时任务调度器
func (s *Scheduler) Start() {
	c := cron.New()
	spec := "0 */10 * * * *"
	_, err := c.AddFunc(spec, s.checkBalance)
	if err != nil {
		panic(err)
	}
	c.Start()
}

// checkBalance 方法检查平台余额
func (s *Scheduler) checkBalance() {
	for _, p := range s.platforms {
		balance, err := p.Balance()
		if err != nil {
			fmt.Printf("Failed to get balance from platform %s: %v\n", p.Type(), err)
			continue
		}
		fmt.Printf("Got balance %.2f from platform %s\n", balance, p.Type())
		s.srv.UpdateBalance(p.Type(), balance)
	}
}
该文件定义了一个Scheduler结构体，用于定时请求平台余额信息。NewScheduler函数创建了一个调度器实例，Start方法启动了定时任务调度器。checkBalance方法遍历平台列表，依次请求每个平台的余额信息，并调用Service的UpdateBalance方法更新余额信息。具体来说，checkBalance方法调用平台的Balance方法获取余额信息，然后打印日志并调用Service的UpdateBalance方法更新余额信息。最后，该方法会在控制台输出获取到的余额信息。