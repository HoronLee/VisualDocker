package docker

import (
	"VDController/logger"
	"context"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

var (
	dockerClient *client.Client
)

func initDocker() {
	var err error
	dockerClient, err = newClient()
	if err != nil {
		logger.GlobalLogger.Log(logger.ERROR, "Docker Client creation failed")
	} else {
		logger.GlobalLogger.Log(logger.INFO, "Docker Client was successfully created")
	}
}

// NewClient 创建一个包含 Docker 客户端的新实例
func newClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	return cli, nil
}

// Close 关闭 Docker 客户端连接
func Close() error {
	return dockerClient.Close()
}

// Dockerls 获取当前容器
func Dockerls() ([]types.Container, error) {
	containers, err := dockerClient.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		logger.GlobalLogger.Log(logger.ERROR, "Failed to get containers")
	} else {
		logger.GlobalLogger.Log(logger.INFO, "Success to get containers")
	}
	return containers, err
}

// DockerlsByImg 通过镜像名获得容器
func DockerlsByImg(imgName string) (map[string]interface{}, error) {
	containers, err := Dockerls()
	if err != nil {
		return nil, err
	}

	output := make(map[string]interface{})
	for _, ctr := range containers {
		if strings.Contains(ctr.Image, imgName) {
			output[ctr.Image] = ctr
		}
	}

	if len(output) == 0 {
		output["WARN"] = "No Container matches this condition."
	}
	return output, nil
}
