package analyzer

import (
    "bytes"
    "encoding/json"
    "log"
    "net/http"
)

type AIAnalyzer struct {
    output     chan string
    serviceURL string
}

func NewAIAnalyzer(output chan string, serviceURL string) *AIAnalyzer {
    return &AIAnalyzer{
        output:     output,
        serviceURL: serviceURL,
    }
}

func (aa *AIAnalyzer) Analyze(logLine string) {
    // 构造请求数据
    requestData := map[string]string{
        "log": logLine,
    }
    jsonData, err := json.Marshal(requestData)
    if err != nil {
        log.Printf("Error marshaling request data: %v", err)
        return
    }

    // 发送 POST 请求
    resp, err := http.Post(aa.serviceURL, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        log.Printf("Error calling AI service: %v", err)
        return
    }
    defer resp.Body.Close()

    // 解析响应
    var result struct {
        Label      int     `json:"label"`
        Confidence float64 `json:"confidence"`
    }
    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        log.Printf("Error parsing AI service response: %v", err)
        return
    }

    // 根据结果判断是否发送告警
    if result.Label == 1 {
        aa.output <- logLine
    }
}
