package snake

import (
	"errors"
	"github.com/mikellxy/little_pineapple/isnake"
	"math/rand"
	"sync"
	"time"
)

const (
	DIRLEFT          = "left"
	DIRRIGHT         = "right"
	DIRUP            = "up"
	DIRDOWN          = "down"
	COLORHEAD        = 0x087006
	COLORBOBY        = 0x10e10c
	COLORPINEAPPLE   = 0xd18429
	COLORBACKGROUND  = 0x0
	INTCODEHEAD      = 4
	INTCODEBODY      = 1
	INTCODPINEAPPLE  = 2
	INTCODEBACKGROUD = 0
)

var (
	errStrikeEdge error                             // Sanke strike edges of the game map, game over
	errInvalidDir error                             // invalid movement direction
	errSelfStrike error                             // snake's head strikes its body
	errSnakeGrows error                             // the snake grows, no need to move
	errWin        error                             // win the game
	nextXYFunc    map[string]func(*Node) (int, int) // next get next x/y coordinate points the Sanke moving to
	dirMap        map[string]string
	dirMapZh      map[string]string
)

func init() {
	errStrikeEdge = errors.New("strike edge error")
	errInvalidDir = errors.New("invalid movement direction error")
	errSnakeGrows = errors.New("snake got new head")
	errSelfStrike = errors.New("strike self")
	errWin = errors.New("you win")
	nextXYFunc = make(map[string]func(*Node) (int, int))
	nextXYFunc[DIRLEFT] = func(head *Node) (int, int) {
		return head.xCoor - 1, head.yCoor
	}
	nextXYFunc[DIRRIGHT] = func(head *Node) (int, int) {
		return head.xCoor + 1, head.yCoor
	}
	nextXYFunc[DIRUP] = func(head *Node) (int, int) {
		return head.xCoor, head.yCoor - 1
	}
	nextXYFunc[DIRDOWN] = func(head *Node) (int, int) {
		return head.xCoor, head.yCoor + 1
	}
	dirMap = make(map[string]string)
	dirMap["a"] = DIRLEFT
	dirMap["d"] = DIRRIGHT
	dirMap["w"] = DIRUP
	dirMap["s"] = DIRDOWN

	dirMapZh = make(map[string]string)
	dirMapZh[DIRLEFT] = "向左"
	dirMapZh[DIRRIGHT] = "向右"
	dirMapZh[DIRUP] = "向上"
	dirMapZh[DIRDOWN] = "向下"
}

// a Sanke is described by a linked list
type Node struct {
	Next  *Node
	xCoor int // x coordinate point
	yCoor int // y coordinate point
}

func NewSnake(winLen int, autoDir string, gmWidth, gmHeight, ms int, gm isnake.IGameMap) *List {
	l := &List{
		WinLen:       winLen,
		AutoDir:      autoDir,
		GameOver:     false,
		MilliSeconds: ms,
		InputChan:    make(chan string),
	}
	l.AddGameMap(gmWidth, gmHeight, gm)
	return l
}

// hold head, length, game map etc. of the Sanke(linked list)
// make the linked list easier to use
// WinLen: win the game if the snake has grown to this length
// AutoDir: the snake auto moves periodically in this direction if the player dose not input any direction
type List struct {
	Head         *Node
	Len          int
	WinLen       int
	GameMap      isnake.IGameMap
	GameMapArr   [][]int
	LeftLimit    int
	RightLimit   int
	UpLimit      int
	DownLimit    int
	AutoDir      string
	GameOver     bool
	MilliSeconds int
	InputChan    chan string
	sync.Mutex
}

func (l *List) AddGameMap(width, height int, gm isnake.IGameMap) {
	gma := make([][]int, height, height)
	for i := range gma {
		gma[i] = make([]int, width, width)
	}
	l.GameMapArr = gma

	gm.Init(width, height, 0x0)
	l.GameMap = gm

	l.RightLimit = width - 1
	l.DownLimit = height - 1
	headX := width / 2
	headY := height / 2
	l.Head = &Node{
		xCoor: headX,
		yCoor: headY,
	}
	l.Len = 1
	gma[headY][headX] = 4
	gm.FillRect(headX, headY, COLORHEAD, false)
	l.setPineApple()
}

func (l *List) setPineApple() {
	rand.Seed(time.Now().Unix())
	for {
		randX := rand.Intn(l.RightLimit)
		randY := rand.Intn(l.DownLimit)
		if l.GameMapArr[randY][randX] != 0 {
			continue
		}
		l.GameMapArr[randY][randX] = 2
		l.GameMap.FillRect(randX, randY, COLORPINEAPPLE, false)
		return
	}
}

func (l *List) Start() {
	l.showGameMap()

	go func() {
		for {
			dir := <- l.InputChan
			l.AutoDir = dir
		}
	}()

	// periodically auto move
	go func() {
		for {
			time.Sleep(time.Duration(l.MilliSeconds) * time.Millisecond)
			err := l.AutoMove()
			if err != nil {
				return
			}
			l.showGameMap()
		}

	} ()

	l.GameMap.CatchInput(l.InputChan)
}

func (l *List) AutoMove() error {
	err := l.Move(l.AutoDir)
	if err != nil {
		return err
	}
	return nil
}

func (l *List) Move(dir string) error {
	l.AutoDir = dir
	// firstly, get the next location of snake's head after move
	nextX, nextY, err := l.PrepareMove(dir)
	if err == errSnakeGrows {
		// snake get new head, no need to move other nodes
		l.Head = &Node{
			xCoor: nextX,
			yCoor: nextY,
			Next:  l.Head,
		}
		l.GameMapArr[l.Head.Next.yCoor][l.Head.Next.xCoor] = 1
		l.GameMap.FillRect(l.Head.Next.xCoor, l.Head.Next.yCoor, COLORBOBY, false)
		l.GameMapArr[nextY][nextX] = 4
		l.GameMap.FillRect(nextX, nextY, COLORHEAD, false)
		// set a new pineapple
		l.setPineApple()
		l.Len++
		if l.Len == l.WinLen {
			l.GameMap.Refresh()
			l.GameOver = true
			return errWin
		}
		return nil
	}
	if err == errInvalidDir {
		return nil
	}
	if err != nil {
		return err
	}
	// other nodes of snake move to the previous location of its "prev" node one by one
	l.moveOneByOne(nextX, nextY)
	return nil
}

// validate the movement direction
// if validation passed, return head or new head(in case of snake grows) after move
func (l *List) PrepareMove(dir string) (int, int, error) {
	// 头部移动
	f := nextXYFunc[dir]
	nextX, nextY := f(l.Head)
	if nextX < l.LeftLimit || nextX > l.RightLimit || nextY < l.UpLimit || nextY > l.DownLimit {
		// strike edge of game map
		return nextX, nextY, errStrikeEdge
	} else if l.Head.Next != nil && l.Head.Next.xCoor == nextX && l.Head.Next.yCoor == nextY {
		// in case of the snake's length is large than 1, it can never go back to the most last location
		return nextX, nextY, errInvalidDir
	} else if l.GameMapArr[nextY][nextX] == 1 {
		// except of invalid direction case, if the location moving to is part of the snake, means it strikes its body
		return nextX, nextY, errSelfStrike
	} else if l.GameMapArr[nextY][nextX] == 2 {
		// the snake grows
		return nextX, nextY, errSnakeGrows
	}
	return nextX, nextY, nil
}

func (l *List) moveOneByOne(nextX, nextY int) {
	current := l.Head
	for current != nil {
		tempX := current.xCoor
		tempY := current.yCoor
		current.xCoor = nextX
		current.yCoor = nextY
		if current == l.Head {
			l.GameMapArr[nextY][nextX] = 4
			l.GameMap.FillRect(nextX, nextY, COLORHEAD, false)
		} else {
			l.GameMapArr[nextY][nextX] = 1
			l.GameMap.FillRect(nextX, nextY, COLORBOBY, false)
		}
		l.GameMapArr[tempY][tempX] = 0
		l.GameMap.FillRect(tempX, tempY, COLORBACKGROUND, false)
		current = current.Next
		nextX = tempX
		nextY = tempY
	}
}

func (l *List) showGameMap() {
	l.GameMap.Refresh()
}
