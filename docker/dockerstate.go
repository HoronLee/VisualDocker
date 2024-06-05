package docker

import (
	"VDController/logger"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/docker/docker/client"
)

var EnvInfo = Info{}

type Info struct {
	DockerVersion  string `json:"dockerVersion"`
	DockerCVersion string `json:"dockerComposeVersion"`
}

func Checkstatus() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println("Failed to create Docker client:", err)
		os.Exit(1)
	}
	ifok, status := dockerChecks(cli)
	if !ifok {
		fmt.Println(status)
		os.Exit(1)
	}
	initDocker()
}

func dockerChecks(cli *client.Client) (ifok bool, status string) {
	ctx := context.Background()
	// 检查 Docker 是否在运行
	_, err := cli.Info(ctx)
	if err != nil {
		ifok, status = false, fmt.Sprint(err)
		logger.GlobalLogger.Log(logger.ERROR, status)
		return ifok, status
	}
	// 检查 Docker 版本
	sVersion, err := cli.ServerVersion(ctx)
	if err != nil {
		ifok, status = false, "无法获取 Docker 版本。"
		logger.GlobalLogger.Log(logger.ERROR, status)
		return ifok, status
	} else {
		EnvInfo.DockerVersion = string(sVersion.Version)
		dstatus := "Docker 版本:" + sVersion.Version
		// 检查 Docker Compose 版本
		dockerCompV, err := exec.Command("docker", "compose", "version").Output()
		if err != nil {
			ifok, status = false, "无法获取 Docker Compose 版本，将无法使用Docker Compose功能，\n"+"请参考 https://docs.docker.com/compose/install/ 安装 Docker Compose。"
			logger.GlobalLogger.Log(logger.WARNING, status)
			return ifok, status
		} else {
			versionIndex := strings.Index(string(dockerCompV), "version")
			if versionIndex != -1 {
				versionStr := strings.TrimSpace(string(dockerCompV)[versionIndex+len("version v"):])
				ifok, status = true, dstatus+", "+"Docker Compose 版本:"+versionStr
				EnvInfo.DockerCVersion = versionStr
			} else {
				ifok, status = true, dstatus+"\n"+"无法获取 Docker Compose 版本"
			}
		}
		logger.GlobalLogger.Log(logger.INFO, status)
		return ifok, status
	}
}
