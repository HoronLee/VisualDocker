package web

import (
	"VDController/config"
	"VDController/logger"
	"VDController/web/routes"
	"fmt"

	"sync"
)

var (

	// 互斥锁，保证线程安全
	mutex sync.Mutex
)

func CheckStatus() {
	if config.ConfigData.WebEnable {
		fmt.Println("✅在 http://" + config.ConfigData.ListeningAddr + " 上启动 Web 服务")
		go StartWeb()
	} else {
		fmt.Println("❎不启动 Web 服务")
	}
}

func StartWeb() {
	mutex.Lock()
	defer mutex.Unlock()

	logger.GlobalLogger.Log(logger.INFO, "Launching the Web Application")
	listeningAddr := config.ConfigData.ListeningAddr

	router := routes.SetupRouter()
	if router == nil {
		logger.GlobalLogger.Log(logger.ERROR, "Failed to setup router")
		return
	}

	if err := router.Run(listeningAddr); err != nil {
		logger.GlobalLogger.Log(logger.ERROR, "Failed to create listening port")
		panic("ListenAndServe: " + err.Error())
	} else {
		msg := "Listening and serving HTTP on " + listeningAddr
		logger.GlobalLogger.Log(logger.INFO, msg)
	}
}
