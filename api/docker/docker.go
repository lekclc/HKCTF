package dcli

import (
	cfg "ctf/config"

	"github.com/docker/docker/client"
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
