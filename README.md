[![GitHub release](https://img.shields.io/github/v/tag/hlinfocc/cySSHClient2.svg?label=%E6%9C%80%E6%96%B0%E7%89%88%E6%9C%AC)](https://github.com/hlinfocc/cySSHClient2/releases)
[![GitHub release](https://img.shields.io/badge/%E7%AB%8B%E5%8D%B3%E4%B8%8B%E8%BD%BD-cf2727)](https://github.com/hlinfocc/cySSHClient2/releases)

# 介绍
基于Linux系统终端ssh命令的ssh客户端，方便在Linux系统使用ssh命令登录主机，不用记很多主机地址,直接一个命令就可以登录远程主机，省去输入主机地址的麻烦！web端可以管理主机信息，SSH密钥，以及云主机的过期信息（即将过期的主机将会通过钉钉机器人发送通知）等

## 安装

下载系统版本相关的包，如一般电脑下载amd64.tar.gz 结尾的包即可.

解压：

```
tar -zxvf cyssh-client_2.1.0_linux_amd64.tar.gz
```
进入解压目录：

```
cd cyssh-client_2.1.0_linux_amd64
```

执行安装脚本：

```
/bin/bash ./install.sh
```

## 钉钉机器人webhook地址配置

编辑配置文件/etc/cysshClient.yml（没有此文件就新建）钉钉机器人webhook地址

创建钉钉机器人时候，需要添加关键词：主机监控提醒

编辑/etc/cysshClient.yml

```yml
webhook: https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxxx

```

## 使用须知

* 安装后可以使用`cyssh`和`cyscp`命令执行相应操作。

* 了解参数运行`cyssh -?` 、`cyscp -?`

* web端默认端口为31918，如果启动时候被占用则会自动+1，默认账号为admin，默认密码为123456


# 许可证
MIT License 