package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Req struct {
	ctx *gin.Context
}

func ConfigReq(ctx *gin.Context) *Req {
	return &Req{
		ctx: ctx,
	}
}

func (r *Req) ErrorResponse(statuscode int, err error) {
	r.ctx.JSON(statuscode, gin.H{
		"sucess": "false",
		"error":  err,
	})
}

func (r *Req) OkRespose(data interface{}) {
	r.ctx.JSON(http.StatusOK, gin.H{
		"sucess": "true",
		"data":   data,
	})
}
