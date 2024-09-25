package main

import (
    "flag"
    "log"

    "exception_log_collector/analyzer"
    "exception_log_collector/collector"
    "exception_log_collector/config"
    "exception_log_collector/notifier"
)

func main() {
    configPath := flag.String("config", "configs/config.yaml", "Path to configuration file")
    flag.Parse()

    // 加载配置
    if err := config.LoadConfig(*configPath); err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

    // 通信通道
    logChannel := make(chan string, 100)
    alertChannel := make(chan string, 100)

    // 初始化告警器
    notifiers := []notifier.Notifier{}
    for _, alertConfig := range config.GlobalConfig.Alerts {
        n := notifier.NewNotifier(alertConfig)
        if n != nil {
            notifiers = append(notifiers, n)
        }
    }

    // 启动日志采集器
    logCollector := collector.NewLogCollector(config.GlobalConfig.Paths, logChannel)
    logCollector.Start()

    // 启动分析器
    basicAnalyzer := analyzer.NewBasicAnalyzer(alertChannel)
    aiAnalyzer := analyzer.NewAIAnalyzer(alertChannel, config.GlobalConfig.AIServiceURL)

    go func() {
        for logLine := range logChannel {
            basicAnalyzer.Analyze(logLine)
            aiAnalyzer.Analyze(logLine)
        }
    }()

    // 启动告警器
    go func() {
        for alert := range alertChannel {
            for _, n := range notifiers {
                if err := n.Notify(alert); err != nil {
                    log.Printf("Failed to send alert: %v", err)
                }
            }
        }
    }()

    // 防止主函数退出
    select {}
}
