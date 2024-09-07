package user

import (
	cfg "ctf/config"
	Db "ctf/database"
	"ctf/logic"

	"github.com/gin-gonic/gin"
)

func Info(r *gin.Context) {
	name, auth := logic.Jwt_Info(r)
	db := cfg.DB
	if !auth {
		var user Db.User
		res := db.Where("name = ?", name).First(&user)
		if res.Error != nil {
			logic.Res_msg(r, 200, 0, "user not exists")
			return
		}
		r.JSON(200, gin.H{
			"code": 1,
			"msg":  "success",
			"data": gin.H{
				"username": user.Name,
				"userId":   user.UserID,
			},
		})
	} else if auth {
		var admin Db.Admin
		res := db.Where("name = ?", name).First(&admin)
		if res.Error != nil {
			logic.Res_msg(r, 200, 0, "admin not exists")
			return
		}
		r.JSON(200, gin.H{
			"code": 1,
			"msg":  "success",
			"data": gin.H{
				"username": admin.Name,
				"adminId":  admin.AdminID,
			},
		})
	}
}
