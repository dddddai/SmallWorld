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

## Usage

- 将SmallWorld.exe与ydk文件放在同一目录下


- 或者使用`-p`参数指定ydk文件的路径，例如 `SmallWorld.exe -p path/to/mydeck.ydk`

运行即可打印卡组中的所有检索路径