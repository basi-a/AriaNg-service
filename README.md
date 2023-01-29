** English | [Chinese](#CN) **
# EN
## AriaNg-service
A Mini program to start, stop install the "mayswind/AriaNg" latest release
### install
Rely on `jq wget unzip go lsof`
Resolve dependencies example Archlinux
```bash
sudo pacman -S jq wget unzip go lsof
```
Get source code and build to binary file
```bash
git clone https://github.com/basi-a/AriaNg-service.git
cd AriaNg-service
go mod tidy && go build AriaNg-service.go
```

### How to Use
```bash
./AriaNg-service help
+--------------------------------------------------------+
+ command  ->  AriaNg-service <option>                   +
+--------------------------------------------------------+
+ help       : show help                                 +
+ start      : start AriaNg                              +
+ stop       : stop AriaNg                               +
+ install    : install AriaNg least release              +
+ reinstall  : delate old AriaNg release and install new +
+--------------------------------------------------------+
```

# CN
## AriaNg-service
一个小程序来开启、关闭和安装 `mayswind/AriaNg` 的最新release
## 安装
依赖 `jq wget unzip go lsof`
解决依赖，以archlinux为例
```bash
sudo pacman -S jq wget unzip go lsof
```
获取源码并编译成可执行二进制文件
```bash
git clone https://github.com/basi-a/AriaNg-service.git
cd AriaNg-service
go mod tidy && go build AriaNg-service.go
```
### 如何食用
```bash
./AriaNg-service help
+--------------------------------------------------------+
+ command  ->  AriaNg-service <option>                   +
+--------------------------------------------------------+
+ help       : show help                                 +
+ start      : start AriaNg                              +
+ stop       : stop AriaNg                               +
+ install    : install AriaNg least release              +
+ reinstall  : delate old AriaNg release and install new +
+--------------------------------------------------------+
```
