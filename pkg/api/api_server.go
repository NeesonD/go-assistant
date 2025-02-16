package api

import (
    "fmt"
    "log"
    "net/http"
)

func StartAPIServer(port int) {
    http.HandleFunc("/voice", handleVoiceCommand)
    http.HandleFunc("/health", handleHealthCheck)
    
    go func() {
        if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
            log.Fatalf("API server failed: %v", err)
        }
    }()
}

func handleVoiceCommand(w http.ResponseWriter, r *http.Request) {
    // 语音指令处理流水线
    audio := r.Body
    text := speechToText(audio)
    intent := analyzeIntent(text)
    
    eventBus <- Event{
        Type:    "user.command",
        Payload: intent,
    }
    
    // 返回TTS响应
    response := generateResponse(intent)
    w.Write(response)
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
    // 健康检查逻辑
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}