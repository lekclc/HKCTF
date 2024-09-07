package crtl

import (
	"ctf/api/admin"
	dcli "ctf/api/docker"
	lcli "ctf/api/level"
	"ctf/api/tool"
	"ctf/api/user"

	"github.com/gin-gonic/gin"
)

func Router_init(r *gin.Engine) *gin.Engine {

	r.GET("/ping", tool.Ping)

	r.POST("/login", user.Login)
	r.POST("/register", user.Register)

	v1 := r.Group("/v1") //, mid.AuthJwt())
	v1.POST("/info", user.Info)

	v2 := r.Group("/v2") //, mid.AuthAdmin())
	v2.POST("/add", admin.Add)

	level := v2.Group("/level")
	level.POST("/add", lcli.Level_Add)

	docker := v2.Group("/docker")
	docker.POST("/run", dcli.Run)

	return r
}
