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
	fmt.Printf("\033[2J");
	fmt.Printf("\033[0;0H");
}

func (s *screen)Start() {
	go func() {
		num := 0
		for {
			s.Flush()
			// s.Print()
			s.Print2(num)
			time.Sleep(time.Duration(1000/s.fps) * time.Millisecond)
			num += 1
		}
	}()
}

func (s *screen)Stop(){

}

func (s *screen)Print() {
	fmt.Println("Hello")
}

func (s *screen)Print2(num int) {
	fmt.Printf("\033[%d;%dH", num % s.height, (num * 2) % s.width);
	fmt.Println("Hello")
}
