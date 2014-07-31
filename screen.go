package aspirin

import (
	"fmt"
	"time"
)

type screen struct {
	width, height int
	cells []string
	fps int
}

func newScreen(width, height, fps int) *screen{
	s := new(screen)
	s.width  = width
	s.height = height
	s.fps    = fps
	return s
}

func (s *screen)initialize() {

}

func (s *screen)Flush() {

}

func (s *screen)Start() {
	go func() {
		for {
			s.Print()
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

func (s *screen)Stop(){

}

func (s *screen)Print() {
	fmt.Println("Hello")

}
