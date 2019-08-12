# little_pineapple(Snake game implemented in golang)
贪吃蛇golang实现  
Snake game implemented in golang  
---  

数据结构：链表&数组  
Data structures used: linked list&array  
---  

使用方式: go run main.go  
Start by: go run main.go  
---  

v2版本更新基于sdl2模块实现的简单uid界面
In version2.0, update a simple user interface implemented using sdl2  
定义了一个游戏地图的接口, 如下  
Add a abstract interface for the game map, as following:  
```
type IGameMap interface {
	Init(int, int, uint32) error
	FillRect(int, int, uint32, bool) error
	Refresh() error
	Close()
	CatchInput(chan string)
}
```
* Init: 初始化地图. Initailize the game map.
* FillRect: 设置/改变组成地图的最小方格的颜色. Set/update the color of a rectangular on the game map.  
* Refresh: 如果地图上的元素发生改变，用于刷新地图. To refresh the map when items on the map changed.  
* Close: 关闭地图，释放资源. Close the game map, release resources.
* CatchInput: 获取用户输入，通过管道传递给snake模块. Catch user's input, send to the snake model through a channel.  

>游戏地图可以是实现了这个接口的UI(本地游戏), websocket/TCP服务器(在线游戏)  
>A game map can be a UI handler(play locally), a websocket/TCP server, etc. which implements this interface.   
---  

NewSnake方法  
+ winLen：蛇皮多上时游戏获胜  
+ autoDir: 初始化蛇头移动方向  
+ gmWidth: 游戏地图宽度  
+ gmHeight: 游戏地图高度  
+ ms: 间隔多少毫秒移动一次  

NewSnake Method:  
* winLen：snake grows to this length, you win the game  
* autoDir: initial moving direction  
* gmWidth: width of the game map  
* gmHeight: height of the game map  
* ms: move periodically by this milliseconds  
---  

```
/*
little_pineapple.snake.NewSnake
Define the size of game map, how long the snake grows to you can win the game, inital move direction and moving periodic here!
*/
func NewSnake(winLen int, autoDir string, gmWidth, gmHeight, ms int) *List {
... ...
}
```
---  

启动脚本后，将弹出游戏窗口，可以通过键盘上的方向键控制移动  
After the script runs, a game window will be shown. To control the movement of the snake by your keyborad.  

On the game map, "4" stands for snake's head, "1" stands for snake's head and "2" stands for little pineapple. 脚本
