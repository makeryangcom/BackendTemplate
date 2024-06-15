package Crontab

import (
	"github.com/gookit/color"
	"github.com/robfig/cron/v3"
	"log"
)

func Start() {
	cronTab := cron.New(cron.WithSeconds())
	_, _ = cronTab.AddFunc("*/10 * * * * *", func() {
		log.Println("[crontab]", color.Green.Text("10秒一次..."))
	})
	_, _ = cronTab.AddFunc("0 */1 * * * *", func() {
		log.Println("[crontab]", color.Green.Text("1分钟一次..."))
	})
	cronTab.Start()
}
