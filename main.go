package main

import (
	"fmt"
	"github.com/mikellxy/little_pineapple/ws_server"
	"golang.org/x/net/websocket"
	"net/http"
)

func Echo(ws *websocket.Conn) {
	mm := ws_server.NewServerMap(ws)
	mm.Init(11, 11, mm.GetColorBackground())
	mm.Refresh()
	mm.CatchInput(make(chan string))
}

func main() {
	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
