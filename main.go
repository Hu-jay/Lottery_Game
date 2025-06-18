package main

import (
	"github.com/Hu-jay/Lottery_Game/app/config"
	"github.com/Hu-jay/Lottery_Game/app/controller"
	"github.com/Hu-jay/Lottery_Game/app/repository"
	"github.com/Hu-jay/Lottery_Game/app/service"
	"github.com/gin-gonic/gin"
)

func main() {
	rc := config.NewClient()
	repo := repository.NewRedisRepo(rc)
	svc := service.NewGameService(repo)

	go svc.GameServer()

	router := gin.Default()
	controller.RegisterRoutes(router, svc)
	router.Run(":8080")
}
