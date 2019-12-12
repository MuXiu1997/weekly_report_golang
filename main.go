package main

import (
	"weekly_report_golang/router"
	"weekly_report_golang/scheduler"
)

func main() {
	scheduler.Run()
	r := router.Router
	if err := r.Run("0.0.0.0:80"); err != nil {
		panic(err)
	}
}
