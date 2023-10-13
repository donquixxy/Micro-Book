package main

import (
	"log"
	"micro-book/config"
)

func main() {
	conf := config.InitConfig()

	log.Println("App is running on env :", conf.Environment)
	pg := config.InitPostgres(conf)

	_ = pg
}
