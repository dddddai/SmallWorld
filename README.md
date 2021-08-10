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
go build -o SmallWorld.exe -ldflags="-s -w"
```

## Example
```
请选择要查询的卡组文件：
选择example_deck.ydk

要检索的卡(输入*表示检索所有怪兽)：
G

Monsters:

彼岸的恶鬼 利比科克

彼岸的恶鬼 齐里亚托

魔犀族战士

幻影骑士团 沾尘袍

彼岸的恶鬼 卡尔卡布里纳

彼岸的恶鬼 斯卡尔米利奥内

未界域的鹿角兔

幻影骑士团 破洞鳞甲

破坏剑-龙破坏之剑

幻影骑士团 污痕胫甲

未界域的槌子蛇

幻影骑士团 无声靴

灰流丽

彼岸的恶鬼 格拉菲亚卡内

彼岸的恶鬼 法尔法雷洛

由魔界到现世的死亡导游

彼岸的恶鬼 阿利基诺

幻影骑士团 破手套

念力循轨人

怒炎坏兽 多哥兰

增殖的G


Routes:
+---------------------------+----------------+---------+
|展示手牌                   |展示卡组        |检索     |
+---------------------------+----------------+---------+
|彼岸的恶鬼 利比科克        |念力循轨人      |增殖的G  |
|魔犀族战士                 |未界域的鹿角兔  |增殖的G  |
|幻影骑士团 沾尘袍          |魔犀族战士      |增殖的G  |
|幻影骑士团 沾尘袍          |念力循轨人      |增殖的G  |
|彼岸的恶鬼 卡尔卡布里纳    |念力循轨人      |增殖的G  |
|彼岸的恶鬼 斯卡尔米利奥内  |念力循轨人      |增殖的G  |
|未界域的鹿角兔             |魔犀族战士      |增殖的G  |
|未界域的鹿角兔             |念力循轨人      |增殖的G  |
|幻影骑士团 破洞鳞甲        |魔犀族战士      |增殖的G  |
|幻影骑士团 破洞鳞甲        |念力循轨人      |增殖的G  |
|破坏剑-龙破坏之剑          |未界域的鹿角兔  |增殖的G  |
|幻影骑士团 污痕胫甲        |魔犀族战士      |增殖的G  |
|未界域的槌子蛇             |魔犀族战士      |增殖的G  |
|未界域的槌子蛇             |念力循轨人      |增殖的G  |
|幻影骑士团 无声靴          |魔犀族战士      |增殖的G  |
|幻影骑士团 无声靴          |念力循轨人      |增殖的G  |
|灰流丽                     |魔犀族战士      |增殖的G  |
|灰流丽                     |未界域的鹿角兔  |增殖的G  |
|灰流丽                     |念力循轨人      |增殖的G  |
|彼岸的恶鬼 格拉菲亚卡内    |念力循轨人      |增殖的G  |
|彼岸的恶鬼 法尔法雷洛      |念力循轨人      |增殖的G  |
|彼岸的恶鬼 阿利基诺        |念力循轨人      |增殖的G  |
|幻影骑士团 破手套          |魔犀族战士      |增殖的G  |
|幻影骑士团 破手套          |念力循轨人      |增殖的G  |
|念力循轨人                 |未界域的鹿角兔  |增殖的G  |
|增殖的G                    |魔犀族战士      |增殖的G  |
|增殖的G                    |未界域的鹿角兔  |增殖的G  |
|增殖的G                    |念力循轨人      |增殖的G  |
+---------------------------+----------------+---------+

查询完毕，感谢使用^_^

按Ctrl+C退出
```