package logic

import (
	"github.com/gin-gonic/gin"
)

func Res_msg(r *gin.Context, statu int, code int, msg string, data ...string) {
	if len(data) > 0 {
		r.JSON(statu, gin.H{"code": code, "msg": msg, "data": data})
		return
	}
	r.JSON(statu, gin.H{"code": code, "msg": msg})
}
