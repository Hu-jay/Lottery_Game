package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Hu-jay/Lottery_Game/app/service"
	"github.com/Hu-jay/Lottery_Game/app/util"
)

func RegisterRoutes(r *gin.Engine, svc *service.GameService) {
	r.GET("/players/:user", func(c *gin.Context) {
		u, err := svc.GetBalance(c.Param("user"))
		util.Wrap(c, u, err)
	})
	r.GET("/players/:user/:amount", func(c *gin.Context) {
		amt, err := strconv.Atoi(c.Param("amount"))
		if err != nil {
			util.Wrap(c, nil, err)
			return
		}
		u, e2 := svc.Bet(c.Param("user"), amt)
		util.Wrap(c, u, e2)
	})
	r.GET("/prize", func(c *gin.Context) {
		p := svc.GetPrize()
		util.Wrap(c, p, nil)
	})
	r.GET("/players", func(c *gin.Context) {
		bs := svc.GetBets()
		if len(bs) == 0 {
			util.Wrap(c, nil, fmt.Errorf("目前沒有任何記錄"))
			return
		}
		util.Wrap(c, bs, nil)
	})
}
