package core

import (
    "context"
    "fmt"
    "log"
)

// 主控服务
type Assistant struct {
    modules map[string]Module
    users   map[string]*UserContext
}

// 运行主控服务
func (a *Assistant) Run() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // 启动所有模块
    for _, m := range a.modules {
        go m.Start(ctx)
    }

    // 主事件循环
    for {
        select {
        case event := <-eventBus:
            a.routeEvent(event)
        case <-ctx.Done():
            return
        }
    }
}

// 路由事件到相应模块
func (a *Assistant) routeEvent(evt Event) {
    for _, m := range a.modules {
        if err := m.HandleEvent(evt); err != nil {
            log.Printf("Error handling event: %v", err)
        }
    }
}

// 添加模块
func (a *Assistant) AddModule(name string, module Module) {
    a.modules[name] = module
}

// 添加用户上下文
func (a *Assistant) AddUserContext(userID string, context *UserContext) {
    a.users[userID] = context
}

// 初始化助手
func NewAssistant() *Assistant {
    return &Assistant{
        modules: make(map[string]Module),
        users:   make(map[string]*UserContext),
    }
}