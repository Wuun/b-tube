package main

import (
	"btube/conf"
	"btube/server"
)

func main() {
	//load config from env.
	conf.Init()
	server := server.Router()
	server.Run(conf.GlobalConf.WebAddr)
}
