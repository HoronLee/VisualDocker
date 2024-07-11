package controller

import (
	"GoToKube/docker"
	"GoToKube/kubernetes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"DockerV":        docker.EnvInfo.DockerVersion,
		"DockerComposeV": docker.EnvInfo.DockerCVersion,
		"KubeVersion":    kubernetes.EnvInfo.KubeVersion,
	})
}
