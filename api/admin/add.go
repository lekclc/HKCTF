package admin

import (
	cfg "ctf/config"
	Db "ctf/database"
	"ctf/logic"

	"github.com/gin-gonic/gin"
)

func Add(r *gin.Context) {
	name := r.PostForm("name")
	db := cfg.DB
	var admin Db.Admin
	var user Db.User
	res := db.Where("name = ?", name).First(&admin)
	if res.Error == nil {
		logic.Res_msg(r, 200, 0, "admin exists")
		return
	} else {
		res = db.Where("name = ?", name).First(&user)
		if res.Error != nil {
			logic.Res_msg(r, 200, 0, "user not exists")
			return
		}
		db.Create(&Db.Admin{Name: name, Passwd: admin.Passwd, UserID: user.UserID})
		logic.Res_msg(r, 200, 1, "add admin success")
	}

}
