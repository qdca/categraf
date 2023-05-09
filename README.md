## Categraf


## 编译安装

```shell
# export GO111MODULE=on
# export GOPROXY=https://goproxy.cn
go build
```

## 二进制安装

```shell
tar zcvf categraf.tar.gz categraf conf
```


## 运行

```shell
# test mode: just print metrics to stdout
./categraf --test

# test system and mem plugins
./categraf --test --inputs system:mem

# print usage message
./categraf --help

# run
./categraf

# run with specified config directory
./categraf --configs /path/to/conf-directory

# only enable system and mem plugins
./categraf --inputs system:mem

# use nohup to start categraf
nohup ./categraf &> stdout.log &
```




## 增加UI编辑配置
![image](images/ui.jpg)

