# [English](#EN) | [Chinese](#CN)
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
+ restart    : restart AriaNg                            +
+ install    : install AriaNg least release              +
+ reinstall  : delate old AriaNg release and install new +
+--------------------------------------------------------+
```
```bash
./AriaNg-service install
./AriaNg-service start
```
如果您没有`aria2c.conf`，则可以使用 `aria2c.conf.example` 创建软连接到 `$HOME/.aria2c/aria2c.conf`。
示例脚本：
```bash
./use-dotfile.sh
```
另外BT需要更新 `bt-tracker`, 可以使用以下脚本（ 当配置文件使用的是上面的时 ）:
```bash
./update-tracker.sh
```
***************
# EN
## AriaNg-service
A Mini program to start, stop install the "mayswind/AriaNg" latest release
### install
Rely on `jq wget unzip go lsof`
Solve dependencies, take archlinux as an example
```bash
sudo pacman -S jq wget unzip go lsof
```
Obtain the source code and compile it into an executable binary
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
+ restart    : restart AriaNg                            +
+ install    : install AriaNg least release              +
+ reinstall  : delate old AriaNg release and install new +
+--------------------------------------------------------+
```
```bash
./AriaNg-service install
./AriaNg-service start
```
If you don't have `aria2c.conf`, you can use `aria2c.conf.example` to create a soft connection to `$HOME/.aria2c/aria2c.conf`.
Example script:
```bash
./use-dotfile.sh
```
In addition, BT needs to update `bt-tracker`, you can use the following script (when the configuration file is using the above):
```bash
./update-tracker.sh
```

# Thank for mayswind/AriaNg

