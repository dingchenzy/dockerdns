# dockerdns
多主机互通
    可以直接通过指路由的方式实现互通（网段要求要不同）

```bash
创建 docker 自定义网桥
    创建网桥
    brctl addbr bridge0

    添加物理网卡到网桥
    brctl addif bridge0 ens33&&ip addr del 10.0.0.11/24 dev bridge0

    删除物理网卡上的 ip
    ip addr del 10.0.0.11/24 dev bridge0&&

    停止网桥
    ip link set dev bridge0 down

    添加 ip 到网卡
    ip addr add 10.0.0.128/24 dev bridge0


--driver ./docker-1.11.0-dev 

docker network create -d macvlan \
  --ipam-driver=dhcp \
  -o parent=ens33 \
  --ipam-opt dhcp_interface=ens33 bridge0 
```


```bash
创建网桥
 1752  ip link add bridge0 type bridge
 1753  ip link set bridge0 up

添加物理网卡到网桥中
 1754  ip link set ens33 up
 1755  ip link set ens33 master bridge0

开启网桥同步
 1757  iptables -A FORWARD -i bridge0 -j ACCEPT

启动容器饼指定网络模式
 1759  docker network create -d ghcr.io/devplayer0/docker-net-dhcp:release-linux-amd64 --ipam-driver null -o bridge=bridge0  bridge0-dhcp

给网桥添加 ip
 1762  ip add add 10.0.0.128/24 dev bridge0

 1764  brctl show

启动带 dhcp-client 的容器
 1765  docker run -it --network bridge0-dhcp --name busybox02 --rm busybox sh
```

```http
https://gist.github.com/nerdalert/3d2b891d41e0fa8d688c              # docker dhcp 前代
https://github.com/devplayer0/docker-net-dhcp                       # docker dhcp 现代
```

# ipam 是个什么东西

在docker网络中，CNM(Container Network Management)模块通过IPAM(IP address management)驱动管理IP地址的分配。Libnetwork内含一个默认的IPAM驱动，同时它也允许动态地增加第三方IPAM驱动。在用户创建网络时可以指定libnetwork使用的IPAM驱动。本文档用于解释IPAM驱动需要遵守的API以及相关的HTTPS请求和响应消息体。

```text 
理解
也就是一个模块的软件用来控制当前主机上的容器上的 ip 地址的分配
```

# 2021-8-29

## 实现的内容

调研了技术，必须实现 dns 的自动识别 hosts 文件变化然后动态的就会重载配置
