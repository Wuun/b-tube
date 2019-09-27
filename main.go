package main

import (
	"btube/conf"
	"btube/cron"
	"btube/server"
)

func main() {
	//load config from env.
	conf.Init()
	cron.Init()
	server := server.Router()
	server.Run(conf.GlobalConf.WebAddr)
}
