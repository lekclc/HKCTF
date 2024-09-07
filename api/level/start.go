package lcli

import (
	"ctf/logic"

	"github.com/gin-gonic/gin"
)

func Start(r *gin.Context) {
	logic.Res_msg(r, 200, 1, "ok")
}
