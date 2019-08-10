package main

import (
	snake2 "github.com/mikellxy/little_pineapple/snake"
)

func main() {
	snake := snake2.NewSnake(2, snake2.DIRLEFT, 11, 11, 1000)
	snake.Start()
}
