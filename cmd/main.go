package main

import (
	"log"
	"market/api"
	"market/config"
	"market/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {

	var cfg = config.Load()

	pgStorage, err := postgres.NewConnectionPostgres(&cfg)
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	api.SetUpApi(r, &cfg, pgStorage)

	log.Println("Listening:", cfg.ServiceHost+cfg.ServiceHTTPPort, "...")
	log.Println("Swagger: http://"+cfg.ServiceHost+cfg.ServiceHTTPPort+"/swagger/index.html", "...")
	if err := r.Run(cfg.ServiceHost + cfg.ServiceHTTPPort); err != nil {
		panic("Listent and service panic:" + err.Error())
	}
}
