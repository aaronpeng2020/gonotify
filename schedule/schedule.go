package schedule

import (
	"fmt"

	"github.com/aaronpeng2020/gonotify/platform"
	"github.com/aaronpeng2020/gonotify/service"
	"github.com/robfig/cron/v3"
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
