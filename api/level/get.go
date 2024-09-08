package lcli

import (
	"crypto/rand"
	"ctf/logic"
	"encoding/base64"
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

var Start_port int
var End_port int
var now_port int

func Get_port() int {
	now_port++
	//判断端口是否超出范围
	if now_port > End_port {
		now_port = Start_port
		return Get_port()
	}
	//判断端口是否被占用
	if PortInUse(now_port) {
		now_port++
		return Get_port()
	}

	return now_port
}

func PortInUse(port int) bool {
	checkStatement := fmt.Sprintf("lsof -i:%d ", port)
	output, _ := exec.Command("sh", "-c", checkStatement).CombinedOutput()
	return len(output) > 0
}

func Get_Flag(username string, userid uint) string {
	var flag_ string
	flag_ = username + " " + strconv.FormatUint(uint64(userid), 10) + " " + time.Now().Format("2006-01-02 15:04:05")
	prime_512_1, _ := rand.Prime(rand.Reader, 512)
	prime_512_2, _ := rand.Prime(rand.Reader, 512)
	flag_ = base64.StdEncoding.EncodeToString([]byte(flag_))
	flag_ = flag_ + prime_512_1.String()
	flag_ = base64.StdEncoding.EncodeToString([]byte(flag_))
	flag_ = prime_512_2.String() + flag_
	flag_ = base64.StdEncoding.EncodeToString([]byte(flag_))
	flag := "flag{" + logic.Md5(flag_) + "}"
	return flag
}
