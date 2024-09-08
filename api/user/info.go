package user

import (
	"ctf/logic"

	"github.com/gin-gonic/gin"
)

func Info(r *gin.Context) {
	name, auth, id := logic.Jwt_Info(r)
	if !auth {
		logic.Res_msg(r, 200, 1, "user login success", gin.H{
			"username": name,
			"userId":   id,
		})
	} else if auth {
		logic.Res_msg(r, 200, 2, "admin login success", gin.H{
			"username": name,
			"userId":   id,
		})
	}
}
