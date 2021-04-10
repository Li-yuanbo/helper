package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrResp struct {
	ResStatus *Res `json:"res_status"`
}

//返回请求信息
func SucResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

//返回错误信息
func ErrResponse(c *gin.Context, code int, errCode int64, msg string) {
	c.JSON(code, gin.H{
		"data": ErrResp{ResStatus: NewRes(errCode, msg)},
	})
}

func NewRes(code int64, msg string) *Res {
	return &Res{
		Code: code,
		Msg:  msg,
	}
}
