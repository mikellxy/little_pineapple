package ws_server

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
)

const (
	CodeRefreshMap = 2001
)

type RefreshMapResp struct {
	Code int
	Map [][]uint32
}

func NewServerMap(conn *websocket.Conn) *WsMap {
	return &WsMap{
		Conn:            conn,
		ColorHead:       4,
		ColorBody:       1,
		ColorPineapple:  2,
		ColorBackground: 0,
	}
}

type WsMap struct {
	Conn            *websocket.Conn
	ColorHead       uint32 // 4
	ColorBody       uint32 // 1
	ColorPineapple  uint32 // 2
	ColorBackground uint32 // 0
	ArrMap          [][]uint32
}

func (wm *WsMap) Init(w,h int, color uint32) error {
	wm.ArrMap = make([][]uint32, h, h)
	for i := range wm.ArrMap {
		wm.ArrMap[i] = make([]uint32, w, w)
	}
	wm.FillRect(0,0, wm.GetColorBackground(), true)
	return nil
}

func (wm *WsMap) FillRect(X, Y int, color uint32, fillAll bool) error {
	if fillAll {
		for i, line := range wm.ArrMap {
			for j := range line {
				wm.ArrMap[i][j] = color
			}
		}
		return nil
	}
	wm.ArrMap[Y][X] = color
	return nil
}

func (wm *WsMap) Refresh() error {
	resp := &RefreshMapResp{
		Code: CodeRefreshMap,
		Map: wm.ArrMap,
	}
	jResp, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	fmt.Printf(string(jResp))
	if err := websocket.Message.Send(wm.Conn, string(jResp)); err != nil {
		return err
	}
	return nil
}

func (wm *WsMap) Close() {
	wm.Conn.Close()
}

func (wm *WsMap) CatchInput(ch chan string) {
	var direction string
	for {
		if err := websocket.Message.Receive(wm.Conn, &direction); err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("receive direction:", direction)
		wm.Refresh()
		//ch <- direction
	}
}

func (wm *WsMap) GetColorHead() uint32 {
	return wm.ColorHead
}

func (wm *WsMap) GetColorBody() uint32 {
	return wm.ColorBody
}

func (wm *WsMap) GetColorPineapple() uint32 {
	return wm.ColorPineapple
}

func (wm *WsMap) GetColorBackground() uint32 {
	return wm.ColorBackground
}