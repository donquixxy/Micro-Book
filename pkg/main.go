package main

import (
	"log"
	"micro-book/config"
	"micro-book/server"
	"os"
	"os/signal"
)

func main() {
	conf := config.InitConfig()

	log.Println("App is running on env :", conf.Environment)
	pg := config.InitPostgres(conf)

	server := server.NewServer(conf.Environment)
	_ = pg

	server.ServerWrapper(pg)

	go func() {
		if err := server.Router.Run("localhost:8982"); err != nil {
			log.Fatalf("Gin router error %v", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
