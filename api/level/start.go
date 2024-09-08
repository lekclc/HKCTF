package lcli

import (
	cfg "ctf/config"
	Db "ctf/database"
	"ctf/logic"
	"fmt"
	"os/exec"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Level_Start(r *gin.Context) {
	// 获取参数
	db := cfg.DB
	level_id := r.PostForm("level_id")
	user_id__ := r.PostForm("user_id")
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
	user_id_, _ := strconv.Atoi(user_id__)
	user_id := uint(user_id_)
	Level_port := uint(Get_port())
	port := strconv.Itoa(int(Level_port))
	flagstr := Get_Flag(user_id__, user_id)
	cmd := exec.Command("docker", "run", "-d", "-e", "FLAG="+flagstr, "--name", user_id__+image.Name+strconv.Itoa(time.Now().Second()), "-p", port+":"+image_port, image_id)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("创建容器失败:", err)
		fmt.Println("命令输出:", string(out))
		logic.Res_msg(r, 500, 0, "创建容器失败")
		return
	}
	Container_ID := string(out)

	db.Create(&Db.Container{ContainerID: Container_ID, Port: Level_port, Flag: flagstr, UserID: user_id})
	logic.Res_msg(r, 200, 1, "ok", gin.H{"port": Level_port})
}
