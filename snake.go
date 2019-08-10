package little_pineapple

import (
	"fmt"
	"github.com/go-openapi/errors"
	"os"
	"os/exec"
	"time"
)

const (
	DIRLEFT  = "left"
	DIRRIGHT = "right"
	DIRUP    = "up"
	DIRDOWN  = "down"
)

var (
	moveMap     map[string]func(*snakeNode)
	xBounds     [2]int
	yBounds     [2]int
	isInRange   func(int, int, string) (error, [2]int)
	errGameOver error
	errNewHead  error
	playGround  [][]int
	thisSnake   *snake
)

func init() {
	// x轴y轴的边界
	xBounds = [2]int{0, 9}
	yBounds = [2]int{0, 9}
	errGameOver = errors.New(4001, "Game Over!")
	errNewHead = errors.New(4002, "New Head!")
	isInRange = func(x, y int, dir string) (error, [2]int) {
		gameOver := false
		if dir == DIRLEFT {
			if x <= xBounds[0] {
				gameOver = true
			} else {
				x -= 1
			}
		} else if dir == DIRRIGHT && x >= xBounds[1] {
			if x >= xBounds[1] {
				gameOver = true
			} else {
				x += 1
			}
		} else if dir == DIRUP && y <= yBounds[0] {
			if y <= yBounds[0] {
				gameOver = true
			} else {
				y -= 1
			}
		} else if dir == DIRDOWN && y >= yBounds[1] {
			if y >= yBounds[1] {
				gameOver = true
			} else {
				y += 1
			}
		}

		if gameOver {
			return errGameOver, [2]int{x, y}
		}
		if playGround[y][x] == 2 {
			return errNewHead, [2]int{x, y}
		}
		return nil, [2]int{x, y}
	}
	moveMap = make(map[string]func(*snakeNode))
	moveMap[DIRLEFT] = func(sn *snakeNode) {
		sn.Coor[0] -= 1
	}
	moveMap[DIRRIGHT] = func(sn *snakeNode) {
		sn.Coor[0] += 1
	}
	moveMap[DIRUP] = func(sn *snakeNode) {
		sn.Coor[1] -= 1
	}
	moveMap[DIRDOWN] = func(sn *snakeNode) {
		sn.Coor[1] += 1
	}
	playGround = make([][]int, yBounds[1]+1, yBounds[1]+1)
	for i := 0; i <= yBounds[1]; i++ {
		playGround[i] = make([]int, xBounds[1]+1, xBounds[1]+1)
	}
}

type snakeNode struct {
	Next *snakeNode
	Coor [2]int
}

type snake struct {
	Head *snakeNode
	Len  int
	Dir  string
}

func (s *snake) Move(dir string) {
	s.HeadMove(dir)
}

func (s *snake) HeadMove(dir string) {
	// 头部移动
	prevCoor := s.Head.Coor
	// update playground
	playGround[prevCoor[1]][prevCoor[0]] = 0
	err, nextCoor := isInRange(prevCoor[0], prevCoor[1], dir)
	if err == errNewHead {
		return s.AddNewHead
	}
	if err != nil {
		panic(err)
	}
	moveFunc, ok := moveMap[dir]
	if !ok {
		panic("move func not defined")
	}
	moveFunc(s.Head)
	// update playground
	playGround[s.Head.Coor[1]][s.Head.Coor[0]] = 1

	s.BodyMove(prevCoor)
}

func (s *snake) BodyMove(prevCoor [2]int) {
	currentNode := s.Head.Next
	for currentNode != nil {
		tmpCoor := currentNode.Coor
		// update playground
		playGround[tmpCoor[1]][tmpCoor[0]] = 0
		currentNode.Coor = prevCoor
		prevCoor = tmpCoor
		// update playground
		playGround[currentNode.Coor[1]][currentNode.Coor[0]] = 1
	}
}

func (s *snake) Tiking() {
	return
}

func initGame() *snake {
	xMid := (xBounds[0] + xBounds[1]) / 2
	yMid := (yBounds[0] + yBounds[1]) / 2
	playGround[yMid][xMid] = 1
	playGround[yMid][xMid+2] = 2
	head := &snakeNode{
		Coor: [2]int{xMid, yMid},
	}
	return &snake{
		Head: head,
		Len:  1,
		Dir:  DIRLEFT,
	}
}

func freshlayground() {
	clearStdOut()
	for _, line := range playGround {
		for i, item := range line {
			if i == len(line)-1 {
				fmt.Printf("%d\n", item)
			} else {
				fmt.Printf("%d ", item)
			}
		}
	}
}

func clearStdOut() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	thisSnake = initGame()
	freshlayground()
	for {
		time.Sleep(1 * time.Second)
		thisSnake.Move(DIRRIGHT)
		freshlayground()
	}
}
