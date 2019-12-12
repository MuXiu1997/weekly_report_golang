package scheduler

import (
	"github.com/robfig/cron/v3"
	"time"
	"weekly_report_golang/setting"
)

func Run() {
	tz, err := time.LoadLocation(setting.TimeZone)
	if err != nil {
		panic(err)
	}
	c := cron.New(cron.WithSeconds(), cron.WithLocation(tz))

	spec := "0 0 16 ? * thu"
	_, err = c.AddFunc(spec, job)
	if err != nil {
		panic(err)
	}
	c.Start()
}
