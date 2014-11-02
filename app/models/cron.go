// +----------------------------------------------------------------------
// | GoCMS 0.1
// +----------------------------------------------------------------------
// | Copyright (c) 2013-2014 http://www.6574.com.cn All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: zzdboy <zzdboy1616@163.com>
// +----------------------------------------------------------------------

package models

//定时执行任务
import "time"
import "github.com/revel/revel"

//每分钟执行一次
type EveryMinute struct {
}

func (c EveryMinute) Run() {
	//task := new(Task)

	//抓取cnBeta
	//task.GrabCnbeta()
}

//每五分钟执行一次
type FiveMinutes struct {
}

func (c FiveMinutes) Run() {
	revel.WARN.Println("Cron Time:" + time.Now().Format("2006-01-02 15:04:04"))
}

//每三十分钟执行
type ThirtyMinutes struct {
}

func (c ThirtyMinutes) Run() {
	revel.WARN.Println("Cron Time:" + time.Now().Format("2006-01-02 15:04:04"))
}

//每小时执行
type HourlyMinutes struct {
}

func (c HourlyMinutes) Run() {
	revel.WARN.Println("Cron Time:" + time.Now().Format("2006-01-02 15:04:04"))
}

//每天执行
//每天运行一次,午夜
type DailyMinutes struct {
}

func (c DailyMinutes) Run() {
	revel.WARN.Println("Cron Time:" + time.Now().Format("2006-01-02 15:04:04"))
}

//每周执行
//每周运行一次,周日午夜
type WeeklyMinutes struct {
}

func (c WeeklyMinutes) Run() {
	revel.WARN.Println("Cron Time:" + time.Now().Format("2006-01-02 15:04:04"))
}

//每月执行
//一个月运行一次,半夜,第一个月
type MonthlyMinutes struct {
}

func (c MonthlyMinutes) Run() {
	revel.WARN.Println("Cron Time:" + time.Now().Format("2006-01-02 15:04:04"))
}

//每年执行
//运行一年一次,1月1日午夜
type YearlyMinutes struct {
}

func (c YearlyMinutes) Run() {
	revel.WARN.Println("Cron Time:" + time.Now().Format("2006-01-02 15:04:04"))
}
