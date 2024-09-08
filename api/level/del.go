package lcli

import (
	cfg "ctf/config"
	Db "ctf/database"
	"ctf/logic"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/gin-gonic/gin"
)

func Cont_del(r *gin.Context) {
	_, _, user_id := logic.Jwt_Info(r)
	level_id := r.PostForm("level_id")
	db := cfg.DB
	var cont Db.Container
	err_ := db.Where("User_Id = ? AND Level_Id = ?", user_id, level_id).First(&cont)
	if err_.Error != nil {
		logic.Res_msg(r, 404, 0, "容器不存在")
		return
	}
	Cont_id := cont.ContainerID
	err := cfg.Docker.ContainerRemove(r.Request.Context(), Cont_id, container.RemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   true,
		Force:         true,
	})
	if err != nil {
		fmt.Println(err)
		logic.Res_msg(r, 500, 0, "删除容器失败")
		return
	}

	/*
		cmd := exec.Command("docker", "rm", "-f", Cont_id)
		fmt.Println(cmd)
		out, err := cmd.CombinedOutput()
		if err != nil {
			logic.Res_msg(r, 500, 0, "删除容器失败")
			return
		}
		fmt.Println(string(out))
		if string(out) == container.ContainerID {
			db.Delete(&container)
			logic.Res_msg(r, 200, 1, "ok")
		}
		logic.Res_msg(r, 500, 0, "删除容器失败")
	*/

}
func Level_del() {

}
