package crons

import (
	"log"

	"github.com/msterzhang/onelist/plugins/watch"
	"github.com/robfig/cron/v3"
)

var Cron *cron.Cron

func Run() {
	watch.WatchPath()
}

// 5分钟运行一次
func RunFiveM() {

}

// 6小时运行一次
func RunSixH() {
	watch.WatchPath()
}

// 凌晨两点运行
func DayWork() {
	watch.WatchPath()
}

// init cron jobs
func InitCron() {
	go Run()
	Cron = cron.New()
	_, err := Cron.AddFunc("@every 6h", RunSixH)
	if err != nil {
		log.Fatal("添加任务失败:" + err.Error())
	}
	_, err = Cron.AddFunc("@every 5m", RunFiveM)
	if err != nil {
		log.Fatal("添加任务失败:" + err.Error())
	}
	_, err = Cron.AddFunc("0 00 02 * * *", DayWork)
	if err != nil {
		log.Fatalf("添加任务失败:%s", err.Error())
	}
	Cron.Start()
}
