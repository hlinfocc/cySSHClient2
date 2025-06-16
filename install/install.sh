#!/bin/bash

os=$(uname -s | tr '[:upper:]' '[:lower:]')


dqwz=$(dirname `readlink -f $0`)
echo "Install CySSHClient V2 Begin"

linuxInstall(){
    if [ -d "/usr/local/CySSHClient/CySSHClient.py" ];then
        mv /usr/local/CySSHClient /usr/local/CySSHClient_v1
    fi
    # 检测并创建文件夹
    if [ ! -d "/usr/local/CySSHClient" ];then
      mkdir -p /usr/local/CySSHClient
    fi

    echo "cp -rf ${dqwz}/* /usr/local/CySSHClient/"
    \mv -f ${dqwz}/* /usr/local/CySSHClient/

    # 复制V1.0数据库文件到新版本
    if [ -f "/usr/local/CySSHClient_v1/cyssh.db" ];then
        echo "copy /usr/local/CySSHClient_v1/cyssh.db to /usr/local/CySSHClient/"
        \cp /usr/local/CySSHClient_v1/cyssh.db /usr/local/CySSHClient/
    fi

    # 删除V1版本
    if [ -d "/usr/local/CySSHClient_v1" ];then
        echo "delete /usr/local/CySSHClient_v1"
        rm -rf /usr/local/CySSHClient_v1
        rm -f /usr/bin/cyssh
    fi

    rm -f /usr/local/CySSHClient/install.sh

    chown -R root:root /usr/local/CySSHClient/cyssh
    chown -R root:root /usr/local/CySSHClient/cyscp
    chown -R root:root /usr/local/CySSHClient/csc-server

    chmod +x /usr/local/CySSHClient/cyssh
    chmod +x /usr/local/CySSHClient/cyscp
    chmod +x /usr/local/CySSHClient/csc-server

    ln -sf /usr/local/CySSHClient/cyssh /usr/local/bin/
    ln -sf /usr/local/CySSHClient/cyscp /usr/local/bin/

    echo "初始化数据......"
    cd /usr/local/CySSHClient/
    ./csc-server -init

    if [ ! -f "/etc/systemd/system/cyssh.service" ];then
      \cp -f /usr/local/CySSHClient/cyssh.service /etc/systemd/system/
      systemctl enable cyssh
      systemctl start cyssh
      systemctl status cyssh --no-pager
    else
      systemctl restart cyssh
      systemctl status cyssh --no-pager
    fi
    echo "安装完毕!"
}


# 判断操作系统类型
case "$os" in
  linux*)
    if [ $(id -u) != "0" ]; then
        echo "Error: You must be root to run this script, please use root to install"
        exit 1
    fi
    linuxInstall
    ;;
  darwin*)
    echo "当前系统是 macOS (Darwin)，请手动自行安装"
    ;;
  freebsd*)
    echo "当前系统是 FreeBSD，请手动自行安装"
    ;;
  *)
    echo "未知操作系统: ${os}，请手动自行安装"
    exit 1
    ;;
esac
