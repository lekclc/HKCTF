package lcli

import (
	cfg "ctf/config"
	Db "ctf/database"
	"ctf/logic"
	"fmt"
	"os/exec"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Level_Start(r *gin.Context) {
	// 获取参数
	db := cfg.DB
	level_id := r.PostForm("level_id")
	user_name, _, user_id := logic.Jwt_Info(r)
	var level Db.Level
	var image Db.Images
	db.Where("ID = ?", level_id).First(&level)
	if level.ID == 0 {
		logic.Res_msg(r, 404, 0, "level不存在")
		return
	}
	db.Where("ID = ?", level.ID).First(&image)
	if image.ID == 0 {
		logic.Res_msg(r, 404, 0, "image不存在")
		return
	}
	image_id := image.ImageID
	image_port := strconv.Itoa(int(image.Port))
	Level_port := uint(Get_port())
	port := strconv.Itoa(int(Level_port))
	flagstr := Get_Flag(user_name, uint(user_id))
	var container Db.Container
	db.Where("User_Id = ? AND Level_Id = ?", user_id, level_id).First(&container)
	if container.ID != 0 {
		logic.Res_msg(r, 500, 0, "容器已存在")
		return
	}

	cmd := exec.Command("docker", "run", "-d", "-e", "FLAG="+flagstr, "--name", image.Name+port, "-p", port+":"+image_port, image_id)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("创建容器失败:", err)
		fmt.Println("命令输出:", string(out))
		logic.Res_msg(r, 500, 0, "创建容器失败")
		return
	}
	Container_ID := string(out)
	flagstr = logic.Md5(flagstr)
	db.Create(&Db.Container{ContainerID: Container_ID, Port: Level_port, Flag: flagstr, UserID: uint(user_id), LevelId: level.ID})
	var user Db.User
	db.Where("User_id = ?", user_id).First(&user)
	db.Model(&user).Update("ContNum", user.ContNum+1)
	logic.Res_msg(r, 200, 1, "ok", gin.H{"port": Level_port})
}
