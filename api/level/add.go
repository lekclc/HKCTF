package lcli

import (
	"archive/zip"
	cfg "ctf/config"
	Db "ctf/database"
	"ctf/logic"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Level_Add(r *gin.Context) {
	file, header, err := r.Request.FormFile("file")
	if err != nil {
		logic.Res_msg(r, 400, 0, "上传文件错误")
		return
	}
	defer file.Close()
	file_name := header.Filename
	filename := logic.Passwd_hash(file_name+time.Now().GoString()) + ".zip"
	//imageName := r.PostForm("name")
	imageName := "test"
	f, err := os.Create("tmp/" + filename)
	if err != nil {
		logic.Res_msg(r, 500, 0, "文件保存错误")
		return
	}
	defer f.Close()
	io.Copy(f, file)
	err = unzip("tmp/"+filename, "upload/")
	if err != nil {
		logic.Res_msg(r, 500, 0, "解压文件错误")
		return
	}
	os.Remove("tmp/" + filename)

	dir := "upload/" + file_name[:len(file_name)-4]
	fmt.Println(dir)

	cmd := exec.Command("docker", "build", "-t", imageName, dir)
	fmt.Println(cmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("构建 Docker 镜像失败:", err)
		fmt.Println("命令输出:", string(output))
		os.RemoveAll(dir)
		return
	}

	fmt.Println("Docker 镜像构建成功!")
	fmt.Println("命令输出:", string(output))

	idx := strings.LastIndex(string(output), "writing image sha256:")

	if idx != -1 {
		idx += len("writing image sha256:")
		sha256 := string(output)[idx : idx+64]
		db := cfg.DB
		db.Create(&Db.Level{Name: imageName, Score: 100, Mode: 1})
		var level Db.Level
		res := db.Where("name = ?", imageName).First(&level)
		if res.Error != nil {
			logic.Res_msg(r, 500, 0, "数据库错误")
			os.RemoveAll(dir)
			return
		}
		db.Create(&Db.Images{ImageID: sha256, Name: imageName, ID: level.ID, Port: 0})
	}

	logic.Res_msg(r, 200, 1, "ok")
	os.RemoveAll(dir)

}

func unzip(src string, dest string) error {
	// 打开 zip 文件
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	// 遍历 zip 文件中的每个文件
	for _, file := range r.File {
		// 拼接目标路径
		filePath := filepath.Join(dest, file.Name)

		// 打印正在解压的文件名
		fmt.Println("解压: ", filePath)

		// 检查文件路径合法性
		if !strings.HasPrefix(filePath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("非法文件路径: %s", filePath)
		}

		// 如果是目录，创建目录
		if file.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		// 如果是文件，先创建目录
		os.MkdirAll(filepath.Dir(filePath), os.ModePerm)

		// 创建文件
		outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer outFile.Close()

		// 打开 zip 文件中的数据流
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		// 将 zip 文件内容写入目标文件
		_, err = io.Copy(outFile, rc)
		if err != nil {
			return err
		}
	}

	return nil
}
