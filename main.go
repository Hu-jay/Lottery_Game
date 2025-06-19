package main

import (
	"github.com/Hu-jay/Lottery_Game/app/config"
	"github.com/Hu-jay/Lottery_Game/app/controller"
	"github.com/Hu-jay/Lottery_Game/app/repository"
	"github.com/Hu-jay/Lottery_Game/app/service"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB("root:wenga369@tcp(127.0.0.1:3306)/sys?charset=utf8mb4&parseTime=True&loc=Local")
	rc := config.NewClient()
	repo := repository.NewRedisRepo(rc)
	mRepo := repository.NewMySQLRepo(config.DB)
	svc := service.NewGameService(repo, mRepo)

	go svc.GameServer()

	router := gin.Default()
	controller.RegisterRoutes(router, svc)
	router.Run(":8080")
}
