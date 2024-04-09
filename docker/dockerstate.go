package docker

import (
	"VDController/logger"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var tLogger = logger.NewLogger(logger.INFO)

func CheckState() {
	ifok, state := dockerChecks()
	if ifok {
		fmt.Println(state)
	} else if strings.Contains(state, "Docker") {
		fmt.Println(state)
		os.Exit(1)
	} else if strings.Contains(state, "无法获取 Docker") {
		fmt.Println(state)
		os.Exit(1)
	} else if strings.Contains(state, "无法获取 Docker Compose") {
		fmt.Println(state)
	} else {
		fmt.Println(state)
		os.Exit(1)
	}
}

// 环境信息全局变量以及结构
var eInfo = envInfo{}

type envInfo struct {
	DockerVersion   string
	DcomposeVersion string
}

func dockerChecks() (ifok bool, state string) {
	// 检查 Docker 是否在运行
	_, err := exec.Command("docker", "info").Output()
	if err != nil {
		ifok, state = false, "Docker 不在运行，请先启动 Docker! \n"+"请参考 https://docs.docker.com/engine/install/ 安装 Docker。"
		tLogger.Log(logger.ERROR, state)
		return ifok, state
	} else {
		// 检查 Docker 版本
		dockerV, err := exec.Command("docker", "version", "--format", "{{.Server.Version}}").Output()
		if err != nil {
			ifok, state = false, "无法获取 Docker 版本。"
			tLogger.Log(logger.ERROR, state)
			return ifok, state
		} else {
			eInfo.DockerVersion = string(dockerV)
			dstate := "Docker 版本:" + strings.TrimSpace(string(dockerV))
			// 检查 Docker Compose 版本
			dockerCompV, err := exec.Command("docker-compose", "version").Output()
			if err != nil {
				ifok, state = false, "无法获取 Docker Compose 版本，将无法使用Docker Compose功能，\n"+"请参考 https://docs.docker.com/compose/install/ 安装 Docker Compose。"
				tLogger.Log(logger.WARNING, state)
				return ifok, state
			} else {
				versionIndex := strings.Index(string(dockerCompV), "version ")
				if versionIndex != -1 {
					versionStr := strings.TrimSpace(string(dockerCompV)[versionIndex+len("version "):])
					ifok, state = true, dstate+", "+"Docker Compose 版本:"+versionStr
					eInfo.DcomposeVersion = string(versionStr)
				} else {
					ifok, state = true, dstate+"\n"+"无法获取 Docker Compose 版本"
				}
			}
			tLogger.Log(logger.INFO, state)
			return ifok, state
		}
	}
}

func GetEnvInfo() envInfo {
	return eInfo
}
