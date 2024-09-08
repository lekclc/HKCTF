package dcli

import (
	cfg "ctf/config"
	"ctf/logic"

	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

func Init() {
	remoteDockerURL := "unix:///var/run/docker.sock"
	apiClient, err := client.NewClientWithOpts(
		client.WithHost(remoteDockerURL),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		panic(err)
	}
	//defer apiClient.Close()
	cfg.Docker = apiClient
}

func Run(r *gin.Context) {
	logic.Res_msg(r, 200, 1, "ok")
	Test()
}

func Test() {

}
