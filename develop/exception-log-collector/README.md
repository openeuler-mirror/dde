# README

## Python 脚本 scripts/log_electra_service.py

<!-- prettier-ignore -->
!!! NOTE 说明
    该脚本使用 Flask 框架创建了一个简单的 HTTP 服务。
    加载了预训练的 LogELECTRA 模型，用于日志的异常检测。
    提供 /analyze 接口，接受 POST 请求，输入日志内容，返回预测结果。

运行服务：

```bash
python scripts/log_electra_service.py
```

## 运行步骤

### 配置 Python 环境

- 安装 Python 3.6 或以上版本。
- 安装所需库：

  ```bash
  pip install flask transformers torch
  ```

- 如果下载模型速度较慢，可以考虑使用国内镜像源。

### 启动 Python AI 服务

```bash
python scripts/log_electra_service.py
```

### 编译并运行 Go 程序

- 在项目根目录下，执行：

  ```bash
  go mod tidy
  go build -o exception_log_collector
  ```

- 运行程序：

  ```bash
  ./exception_log_collector -config configs/config.yaml
  ```

## 注意事项

1. **权限问题：** 确保运行程序的用户对日志文件具有读取权限。
2. **配置文件：** 根据实际情况修改 `configs/config.yaml`，配置要监控的日志文件路径和告警方式。
3. **Python 服务：** 在运行 Go 程序之前，确保 Python 服务已经启动，并且配置文件中的 `ai_service_url` 地址正确。
4. **邮件告警：** 如果使用邮件告警，需在配置中添加 SMTP 服务器的认证信息。
5. **Kafka 告警器：** 如果需要使用 Kafka 告警器，需要在配置中添加 Kafka 的 brokers 和 topic 信息，并确保 Kafka 服务正常运行。
6. **模型加载时间：** 第一次运行 Python 服务时，模型加载可能需要一些时间，请耐心等待。

## 代码说明

### 日志采集器

- 使用 `fsnotify` 库监听日志文件的变化，实现实时日志采集。
- 当日志文件有新的内容写入时，读取新增的日志行并发送到 `logChannel`。

### 日志分析器

- **基础分析器：**

  - 使用正则表达式匹配常见的错误级别（`ERROR`、`WARN`、`FATAL`）。
  - 如果匹配成功，认为该日志行是异常日志，发送到 `alertChannel`。

- **AI 分析器：**
  - 调用 Python 服务，对日志行进行深度学习模型的分析。
  - 如果模型预测结果为异常日志（`label == 1`），则发送到 `alertChannel`。

### 告警器

- **Webhook 告警器：**

  - 发送 HTTP POST 请求，将异常日志发送到指定的 Webhook URL。

- **Email 告警器：**

  - 使用 `gomail` 库，通过 SMTP 协议发送邮件，将异常日志发送到指定的邮箱。

- **Kafka 告警器：**
  - 使用 `sarama` 库，将异常日志发送到指定的 Kafka topic。

### 主程序

- 加载配置，初始化各个模块。
- 启动日志采集器、分析器和告警器。
- 使用通道 (`channel`) 实现模块间的通信。
- 使用 `select {}` 防止主函数退出。

## 改进建议

1. **日志采集器改进：**

   - 处理日志文件滚动的情况，当日志文件被轮转时，重新打开文件句柄。
   - 考虑使用多线程或异步 IO，提高日志读取的效率。

2. **AI 分析器优化：**

   - 将 Python 服务部署为高可用的微服务，支持负载均衡。
   - 考虑模型的优化，加快预测速度，减少响应时间。

3. **配置管理：**

   - 使用更灵活的配置管理方式，例如支持通过环境变量或命令行参数配置。

4. **日志处理：**

   - 增加对不同日志格式的支持，例如 JSON 格式的日志解析。

5. **错误处理：**

   - 增加错误重试机制，对于告警发送失败的情况，进行重试或持久化。

6. **监控和日志：**
   - 为程序自身添加日志和监控，方便排查问题和性能调优。
