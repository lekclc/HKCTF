package user

import (
	cfg "ctf/config"
	Db "ctf/database"
	"ctf/logic"

	"github.com/gin-gonic/gin"
)

func Login(r *gin.Context) {
	db := cfg.DB

	name := r.PostForm("name")
	passwd := r.PostForm("password")
	passwd = logic.Passwd_hash(passwd)
	var admin Db.Admin
	res := db.Where("name = ? AND passwd = ?", name, passwd).First(&admin)
	if res.Error == nil {
		token := logic.Jwt_get(name, true)
		logic.Res_msg(r, 200, 2, "admin login success", token)
		return
	}
	var user Db.User
	res = db.Where("name = ? AND passwd = ?", name, passwd).First(&user)
	if res.Error != nil {
		logic.Res_msg(r, 200, 0, "login failed")
	} else {
		token := logic.Jwt_get(name, false)
		logic.Res_msg(r, 200, 1, "user login success", token)
	}
}
