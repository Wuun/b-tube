package main

import (
	"btube/conf"
	cron "btube/cronjob"
	"btube/server"
)

func main() {
	//load config from env.
	conf.Init()
	//
	cron.Init()
	s := server.Router()
	panic(s.Run(conf.GlobalConf.WebAddr))
}
