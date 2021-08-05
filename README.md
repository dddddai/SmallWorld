# SmallWorld
游戏王“小世界现象”检索路径查询

查询卡组（.ydk文件）中所有的检索路径

基于 https://ygocdb.com

## Build


#### Windows
点击[此处](https://hub.fastgit.org/dddddai/SmallWorld/releases/latest/download/SmallWorld.exe)下载

#### Build from source
```
git clone git@github.com:dddddai/SmallWorld.git
cd SmallWorld
go build -ldflags="-s -w" -o SmallWorld main.go
```