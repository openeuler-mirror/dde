# 标准用户权限

| 编号 | 步骤                                                         | 预期                                                         |
| ---- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| 1    | 打开控制中心-帐户，在帐户列表点击“+”，新建一个标准用户Standard | 新建标准用户Standard成功                                     |
| 2    | 切换用户为标准用户Standard                                   | 切换Standard用户成功                                         |
| 3    | 使用ctrl+alt+t打开终端，输入sudo passwd root                 | 标准用户没有sudo权限，命令执行失败，提示告警信息【Standard 不在 sudoers 文件中。此事将被报告】 |
