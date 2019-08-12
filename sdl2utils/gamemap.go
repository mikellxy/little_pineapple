package sdl2utils

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	RECTSIZE = 50
	KEYCODEUP = 82
	KEYCODEDOWN = 81
	KEYCODELEFT = 80
	KEYCODERIGHT = 79
	DIRLEFT   = "left"
	DIRRIGHT  = "right"
	DIRUP     = "up"
	DIRDOWN   = "down"
)

type GameMap struct {
	Window  *sdl.Window
	Surface *sdl.Surface
	WinID   uint32
	W       int
	H       int
}

func (gm *GameMap) Init(w, h int, color uint32) error {
	// w or h stands for the number rects per
	// w *RECTSIZE/ h * RECTSIZE is the final width, height of of the map in screen coordinates
	if gm.Window != nil {
		return nil
	}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}

	window, err := sdl.CreateWindow("Snake", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(w * RECTSIZE), int32(h * RECTSIZE), sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}
	gm.Window = window
	gm.WinID, _ = window.GetID()

	surface, err := window.GetSurface()
	if err != nil {
		return err
	}
	gm.Surface = surface
	gm.W = w
	gm.H = h
	err = gm.FillRect(0, 0, 0x0, true)
	if err != nil {
		return err
	}
	return gm.Refresh()
}

func (gm *GameMap) Refresh() error {
	return gm.Window.UpdateSurface()
}

func (gm *GameMap) FillRect(x, y int, color uint32, fillAll bool) error {
	if fillAll {
		return gm.Surface.FillRect(nil, color)
	}
	rect := sdl.Rect{int32(x*RECTSIZE), int32(y*RECTSIZE), RECTSIZE, RECTSIZE}
	return gm.Surface.FillRect(&rect, color)

}

func (gm *GameMap) Close() {
	gm.Window.Destroy()
	sdl.Quit()
}

func (gm *GameMap) CatchInput(ch chan string) {
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("Quit")
				running = false
				gm.Close()
				break
			case *sdl.KeyboardEvent:
				event := event.(*sdl.KeyboardEvent)
				keyCode := event.Keysym.Scancode
				if event.State == 0 {
					switch keyCode {
					case KEYCODEUP:
						ch <- DIRUP
					case KEYCODEDOWN:
						ch <- DIRDOWN
					case KEYCODELEFT:
						ch <- DIRLEFT
					case KEYCODERIGHT:
						ch <- DIRRIGHT
					default:
					}
				}
			}
		}
	}
}
