package main

import (
    "go-assistant/pkg/core"
    "go-assistant/pkg/api"
    "log"
)

func main() {
    assistant := core.Assistant{
        Modules: make(map[string]core.Module),
        Users:   make(map[string]*core.UserContext),
    }

    // 启动API服务器
    go api.StartAPIServer(8080)

    // 启动主控模块
    assistant.Run()

    log.Println("Assistant is running...")
}