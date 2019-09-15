package main

import (
	"btube/conf"
	"btube/server"
)

func main() {
	//load config from env.
	conf.Init()
	server := server.Router()
	server.Run("127.0.0.1:8948")
}
