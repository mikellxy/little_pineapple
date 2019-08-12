package main

import (
	snake2 "github.com/mikellxy/little_pineapple/snake"
)

func main() {
	snake := snake2.NewSnake(8, snake2.DIRLEFT, 16, 12, 1000)
	snake.Start()

}
