﻿﻿﻿﻿#  终端免密提权测试

## 前置条件

1. 存在一个用户能提权到root。

   （操作时应谨慎，避免造成sudo功能不可用）

<br>



## 用例步骤

| 编号 | 步骤                                          | 预期                 |
| ---- | --------------------------------------------- | ------------------- |
| 1    | 普通用户执行："sudo su" | 提权到root   |
| 2    |visudo  |打开sudoers编辑界面|
| 3    |为测试用户配置免密提权，添加一行参数：`testuser     ALL=(ALL:ALL) NOPASSWD:ALL`   |添加成功 |
| 4    | 步骤3添加配置后，依次执行：Ctrl+O、Enter、Ctrl+X | 退出sudoers编辑界面 |
| 5    | 退出当前登陆环境，切换到testuser登陆 | 成功登陆testuser |
| 6 | 以testuser打开终端执行："sudo su" | 不需要输入密码直接提权到root |

<br>





