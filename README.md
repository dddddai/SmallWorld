# SmallWorld
游戏王“小世界现象”检索路径查询

查询卡组（.ydk文件）中所有的检索路径

基于 https://ygocdb.com

## Start


#### Windows
点击[此处](https://hub.fastgit.org/dddddai/SmallWorld/releases/latest/download/SmallWorld.exe)下载

#### Build from source
```
git clone git@github.com:dddddai/SmallWorld.git
cd SmallWorld
go build -ldflags="-s -w" -o SmallWorld main.go
```

## Example
```
请输入卡组（.ydk文件）路径（若输入*则自动查找当前目录）：
*
请输入想要检索的卡（若输入*则输出卡组中所有检索路径）：
G
Monsters:
{id:57143342 name:彼岸的恶鬼 齐里亚托 atk:1600 def:1200
 level:3 race:恶魔 attribute:暗}

{id:30227494 name:念力循轨人 atk:1600 def:600
 level:3 race:念动力 attribute:地}

{id:81035362 name:魔犀族战士 atk:1400 def:900
 level:3 race:恶魔 attribute:地}

{id:73213494 name:彼岸的恶鬼 卡尔卡布里纳 atk:1400 def:0
 level:3 race:恶魔 attribute:暗}

{id:62957424 name:彼岸的恶鬼 利比科克 atk:1300 def:700
 level:3 race:恶魔 attribute:暗}

{id:99745551 name:未界域的槌子蛇 atk:1300 def:0
 level:3 race:爬虫类 attribute:暗}

{id:88544390 name:幻影骑士团 污痕胫甲 atk:1200 def:600
 level:3 race:战士 attribute:暗}

{id:47728740 name:彼岸的恶鬼 阿利基诺 atk:1200 def:0
 level:3 race:恶魔 attribute:暗}

{id:36553319 name:彼岸的恶鬼 法尔法雷洛 atk:1000 def:1900
 level:3 race:恶魔 attribute:暗}

{id:20758643 name:彼岸的恶鬼 格拉菲亚卡内 atk:1000 def:1500
 level:3 race:恶魔 attribute:暗}

{id:10802915 name:由魔界到现世的死亡导游 atk:1000 def:600
 level:3 race:恶魔 attribute:暗}

{id:63821877 name:幻影骑士团 破手套 atk:1000 def:500
 level:3 race:战士 attribute:暗}

{id:84764038 name:彼岸的恶鬼 斯卡尔米利奥内 atk:800 def:2000
 level:3 race:恶魔 attribute:暗}

{id:90432163 name:幻影骑士团 沾尘袍 atk:800 def:1000
 level:3 race:战士 attribute:暗}

{id:25538345 name:幻影骑士团 破洞鳞甲 atk:600 def:1600
 level:3 race:战士 attribute:暗}

{id:43694650 name:未界域的鹿角兔 atk:500 def:2000
 level:3 race:兽 attribute:暗}

{id:36426778 name:幻影骑士团 无声靴 atk:200 def:1200
 level:3 race:战士 attribute:暗}

{id:14558127 name:灰流丽 atk:0 def:1800
 level:3 race:不死 attribute:炎}

{id:23434538 name:增殖的G atk:500 def:200
 level:2 race:昆虫 attribute:地}

{id:76218313 name:破坏剑-龙破坏之剑 atk:400 def:300
 level:1 race:龙 attribute:暗}

Routes:
+---------------------------+----------------+---------+
|展示手牌                   |展示卡组        |检索     |
+---------------------------+----------------+---------+
|念力循轨人                 |未界域的鹿角兔  |增殖的G  |
|魔犀族战士                 |未界域的鹿角兔  |增殖的G  |
|彼岸的恶鬼 卡尔卡布里纳    |念力循轨人      |增殖的G  |
|彼岸的恶鬼 利比科克        |念力循轨人      |增殖的G  |
|未界域的槌子蛇             |念力循轨人      |增殖的G  |
|未界域的槌子蛇             |魔犀族战士      |增殖的G  |
|幻影骑士团 污痕胫甲        |魔犀族战士      |增殖的G  |
|彼岸的恶鬼 阿利基诺        |念力循轨人      |增殖的G  |
|彼岸的恶鬼 法尔法雷洛      |念力循轨人      |增殖的G  |
|彼岸的恶鬼 格拉菲亚卡内    |念力循轨人      |增殖的G  |
|幻影骑士团 破手套          |念力循轨人      |增殖的G  |
|幻影骑士团 破手套          |魔犀族战士      |增殖的G  |
|彼岸的恶鬼 斯卡尔米利奥内  |念力循轨人      |增殖的G  |
|幻影骑士团 沾尘袍          |念力循轨人      |增殖的G  |
|幻影骑士团 沾尘袍          |魔犀族战士      |增殖的G  |
|幻影骑士团 破洞鳞甲        |念力循轨人      |增殖的G  |
|幻影骑士团 破洞鳞甲        |魔犀族战士      |增殖的G  |
|未界域的鹿角兔             |念力循轨人      |增殖的G  |
|未界域的鹿角兔             |魔犀族战士      |增殖的G  |
|幻影骑士团 无声靴          |念力循轨人      |增殖的G  |
|幻影骑士团 无声靴          |魔犀族战士      |增殖的G  |
|灰流丽                     |念力循轨人      |增殖的G  |
|灰流丽                     |魔犀族战士      |增殖的G  |
|灰流丽                     |未界域的鹿角兔  |增殖的G  |
|增殖的G                    |念力循轨人      |增殖的G  |
|增殖的G                    |魔犀族战士      |增殖的G  |
|增殖的G                    |未界域的鹿角兔  |增殖的G  |
|破坏剑-龙破坏之剑          |未界域的鹿角兔  |增殖的G  |
+---------------------------+----------------+---------+
按Ctrl+C退出
```