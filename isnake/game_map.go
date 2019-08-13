package isnake

type IGameMap interface {
	Init(int, int, uint32) error
	FillRect(int, int, uint32, bool) error
	Refresh() error
	Close()
	CatchInput(chan string)
}