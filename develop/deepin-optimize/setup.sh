#!/bin/bash

# 设置桌面所需后台服务为自动启动。
# 这样后续黑屏等待时间会减少
systemctl enable dbus-com.deepin.dde.lockservice.service
systemctl enable systemd-hostnamed.service 
