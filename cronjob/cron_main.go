package cron

import (
	"fmt"
	"reflect"
	"runtime"
	"time"

	"github.com/robfig/cron"
)

//Init start a cron job.
func Init() {
	c := cron.New()
	c.AddFunc("0 /5 * * * *", func() {
		Run(Like)
	})

	c.AddFunc("0 0 0 * *  *", func() {
		Run(DeleteTodayView)
	})
	c.Start()
}

//Run run a job by format,recod it's start time 、 end time and func name.
func Run(job func() error) {
	from := time.Now().UnixNano()
	err := job()
	to := time.Now().UnixNano()
	jobName := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()
	if err != nil {
		fmt.Printf("%s error: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	} else {
		fmt.Printf("%s success: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	}
}
