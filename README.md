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
https://gist.github.com/nerdalert/3d2b891d41e0fa8d688c              # dhcp 插件技术
```
