package user

import (
	cfg "ctf/config"
	Db "ctf/database"
	"ctf/logic"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Context) {
	name := r.PostForm("name")
	passwd := r.PostForm("password")
	db := cfg.DB
	var user Db.User
	res := db.Where("name = ?", name).First(&user)
	if res.Error == nil {
		logic.Res_msg(r, 200, 0, "user exists")
		return
	} else {
		passwd = logic.Passwd_hash(passwd)
		db.Create(&Db.User{Name: name, Passwd: passwd, ContNum: 0})
		logic.Res_msg(r, 200, 1, "register success")
	}
}
