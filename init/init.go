package myinit

import (
	dcli "ctf/api/docker"
	cfg "ctf/config"
	"ctf/crtl"
	db "ctf/database"

	"github.com/gin-gonic/gin"
)

func Init(s *gin.Engine) *gin.Engine {
	cfg.Init()
	crtl.Router_init(s)
	db.Init()
	dcli.Init()

	return s
}
