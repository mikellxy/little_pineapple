# little_pineapple(Snake game implemented in golang)
贪吃蛇golang实现  
Snake game implemented in golang  

数据结构：链表&数组  
Data structures used: linked list&array  

使用方式: go run main.go  
Start by: go run main.go  

NewSnake方法  
  winLen：蛇皮多上时游戏获胜  
  autoDir: 初始化蛇头移动方向  
  gmWidth: 游戏地图宽度  
  gmHeight: 游戏地图高度  
  ms: 间隔多少毫秒移动一次  
NewSnake Method:  
  winLen：snake grows to this length, you win the game  
  autoDir: initial moving direction  
  gmWidth: width of the game map  
  gmHeight: height of the game map  
  ms: move periodically by this milliseconds    
```
/*
little_pineapple.snake.NewSnake
*/
func NewSnake(winLen int, autoDir string, gmWidth, gmHeight, ms int) *List {
... ...
}
```

在游戏地图上, "4"代表蛇头，"1"代表蛇身，"2"代表小菠萝。蛇皮走位吃到小菠萝后变长。  
On the game map, "4" stands for snake's head, "1" stands for snake's head and "2" stands for little pineapple. 
Snake grows longer after eat a little pipeapple.  

启动脚本后，在标准输入，输入w/a/s/d + enter, 表示下一次移动时往蛇头往上/左/下/右移动, 如下:  
Type in w/a/s/d + enter in the cmd line where you start the script, which means snake's head will move up/left/down/right  
next time, like following:  

```
game starts
      0  0  0  0  0  0  0  0  0  0  0
      0  0  0  0  0  0  0  0  0  0  0
      0  0  2  0  0  0  0  0  0  0  0
      0  0  0  0  0  0  0  0  0  0  0
      0  0  0  0  0  0  0  0  0  0  0
      0  0  0  0  0  4  0  0  0  0  0
      0  0  0  0  0  0  0  0  0  0  0
      0  0  0  0  0  0  0  0  0  0  0
      0  0  0  0  0  0  0  0  0  0  0
      0  0  0  0  0  0  0  0  0  0  0
      0  0  0  0  0  0  0  0  0  0  0

```

```
game over!
      0  0  0  0  0  0  0  0  0  0  0
      4  1  1  0  0  0  0  0  0  0  0
      0  0  0  0  0  0  0  0  0  0  0
      0  0  0  0  0  0  0  0  0  0  0
      0  0  0  0  0  0  0  0  0  0  0
      2  0  0  0  0  0  0  0  0  0  0
      0  0  0  0  0  0  0  0  0  0  0
      0  0  0  0  0  0  0  0  0  0  0
      0  0  0  0  0  0  0  0  0  0  0
      0  0  0  0  0  0  0  0  0  0  0
      0  0  0  0  0  0  0  0  0  0  0

      strike edge error, Game Over!

```
