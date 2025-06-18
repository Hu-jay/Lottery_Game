package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Wrap(c *gin.Context, data interface{}, err error) {
	type ret struct {
		Status string      `json:"status"`
		Msg    string      `json:"msg"`
		Data   interface{} `json:"data"`
	}
	r := ret{Status: "ok", Msg: "", Data: []struct{}{}}
	if data != nil {
		r.Data = data
	}
	if err != nil {
		r.Status = "failed"
		r.Msg = err.Error()
	}
	c.JSON(http.StatusOK, r)
}
