package lcli

import (
	cfg "ctf/config"
	Db "ctf/database"
	"ctf/logic"
	"fmt"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func Cont_del(r *gin.Context) {
	_, _, user_id := logic.Jwt_Info(r)
	if user_id == 0 {
		return
	}
	level_id := r.PostForm("level_id")
	db := cfg.DB
	var cont Db.Container
	err_ := db.Where("User_Id = ? AND Level_Id = ?", user_id, level_id).First(&cont)
	if err_.Error != nil {
		logic.Res_msg(r, 404, 0, "容器不存在")
		return
	}
	Cont_id := cont.ContainerID

	exec_str := fmt.Sprintf("docker rm -f %s", Cont_id)
	cmd := exec.Command("sh", "-c", exec_str)
	out, _ := cmd.CombinedOutput()

	if string(out) == cont.ContainerID {
		db.Delete(&cont)
		var user Db.User
		db.Where("id = ?", user_id).First(&user)
		user.ContNum = user.ContNum - 1
		db.Model(&user).Update("Score", user.ContNum)
		logic.Res_msg(r, 200, 1, "ok")
		return
	}
	logic.Res_msg(r, 500, 0, "删除容器失败")

}
func Level_del(r *gin.Context) {

}
