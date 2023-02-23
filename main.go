package main

import (
	"fmt"

	"github.com/aaronpeng2020/gonotify/platform"
	"github.com/aaronpeng2020/gonotify/service"
	"github.com/aaronpeng2020/gonotify/util"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

func main() {
	// 读取配置文件
	viper.SetConfigName("config")
	viper.AddConfigPath("config/")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// 初始化平台信息
	platforms := make([]platform.Platform, 0)
	for _, p := range viper.GetStringMap("platforms") {
		pCfg := p.(map[string]interface{})
		pType := pCfg["type"].(string)
		pConfig := viper.GetStringMap(pType)
		p := platform.NewPlatform(pType, pConfig)
		platforms = append(platforms, p)
	}

	// 初始化邮件和短信工具
	emailUtil := util.NewEmailUtil(viper.GetStringMap("email"))
	smsUtil := util.NewSmsUtil(viper.GetStringMap("sms"))

	// 初始化服务
	srv := service.NewService(platforms, emailUtil, smsUtil, viper.GetFloat64("warn"))

	// 初始化定时任务
	c := cron.New()
	spec := "0 */10 * * * *"
	_, err = c.AddFunc(spec, func() {
		srv.CheckBalance()
	})
	if err != nil {
		panic(err)
	}
	c.Start()

	// 运行服务
	fmt.Println("GoNotify service started.")
	select {}
}
