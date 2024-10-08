package user

import (
	cfg "ctf/config"
	Db "ctf/database"
	"ctf/logic"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(r *gin.Context) {
	db := cfg.DB

	passwd := r.PostForm("password")
	name := r.PostForm("name")
	fmt.Println(name, passwd)
	passwd = logic.Passwd_hash(passwd)
	var admin Db.Admin

	res := db.Where("name = ? AND passwd = ?", name, passwd).First(&admin)
	if res.Error == nil {
		token := logic.Jwt_get(name, true, int(admin.UserID))
		logic.Res_msg(r, 200, 2, "admin login success", gin.H{"token": token})
		return
	}
	var user Db.User
	res = db.Where("name = ? AND passwd = ?", name, passwd).First(&user)
	if res.Error != nil {
		logic.Res_msg(r, 200, 0, "login failed")
	} else {
		token := logic.Jwt_get(name, false, int(user.UserID))
		logic.Res_msg(r, 200, 1, "user login success", gin.H{"token": token})
	}
}
