package tool

import (
	"ctf/logic"

	"github.com/gin-gonic/gin"
)

func Ping(r *gin.Context) {
	logic.Res_msg(r, 200, 1, "pong")
}
