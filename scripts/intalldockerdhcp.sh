#!/bin/bash
# 稍作修改执行这个脚本可以不用退出 ssh
IP=$(ip add|grep 192|awk -F' ' '{print $2}')
ip link add bridge0 type bridge
ip link set ens33 up
ip link set ens33 master bridge0
iptables -A FORWARD -i bridge0 -j ACCEPT
docker network create -d ghcr.io/devplayer0/docker-net-dhcp:release-linux-amd64 --ipam-driver null -o bridge=bridge0 dhcp-bridge
ip add add ${IP} dev bridge0
ip add del ${IP} dev ens33
ip link set bridge0 up