package util

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Hu-jay/Lottery_Game/app/models"
)

func Wrap(c *gin.Context, data interface{}, err error) {

	r := models.Ret{Status: "ok", Msg: "", Data: []struct{}{}}
	if data != nil {
		r.Data = data
	}
	if err != nil {
		r.Status = "failed"
		r.Msg = err.Error()
	}
	c.JSON(http.StatusOK, r)
}
