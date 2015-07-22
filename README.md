Docker external ceph rbd graph driver
=====================================

# How to compile
```bash
go build -v
```
# How to use

## install ceph cluster
TODO:

## run docker rbd driver daemon

```bash
docker-rbd-driver -D
```

## run docker daemon

```bash
# docker -d -D -s rbd
...
```
## pull images

```bash
# docker pull centos:latest
Pulling repository centos
7322fbe74aa5: Download complete 
f1b10cd84249: Download complete 
c852f6d61e65: Download complete 
Status: Downloaded newer image for centos:latest
```

## list rbd image

```bash
# rbd list
docker_image_7322fbe74aa5632b33a400959867c8ac4290e9c5112877a7754be70cfe5d66e9
docker_image_base_image
docker_image_c852f6d61e65cddf1e8af1f6cd7db78543bfb83cdcd36845541cf6d9dfef20a0
docker_image_f1b10cd842498c23d206ee0cbeaa9de8d2ae09ff3c7af2723a9e337a6965d639
```
## run container

```bash
# docker run -it --rm centos:latest /bin/bash
[root@290238155b54 /]#
```

```bash
# rbd list
docker_image_290238155b547852916b732e38bc4494375e1ed2837272e2940dfccc62691f6c
docker_image_290238155b547852916b732e38bc4494375e1ed2837272e2940dfccc62691f6c-init
docker_image_7322fbe74aa5632b33a400959867c8ac4290e9c5112877a7754be70cfe5d66e9
docker_image_base_image
docker_image_c852f6d61e65cddf1e8af1f6cd7db78543bfb83cdcd36845541cf6d9dfef20a0
docker_image_f1b10cd842498c23d206ee0cbeaa9de8d2ae09ff3c7af2723a9e337a6965d639
```
