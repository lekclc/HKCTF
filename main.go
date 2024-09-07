package main

import (
	cfg "ctf/config"
	myinit "ctf/init"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()
	s = myinit.Init(s)
	host_url := fmt.Sprintf(":%d", cfg.Get("g.port"))
	s.Run(host_url)
}
