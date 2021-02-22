package pool

import (
	"fmt"
	"testing"
	"time"
)

type task int

func (t task) Run() {
	fmt.Println("doing someting")
	time.Sleep(time.Second)
}

func TestPool_Start(t *testing.T) {
	p := New(3)
	for i := 0; i < 50; i++ {
		worker := <-p.Workers
		worker <- task(i)
	}
	p.Stop()
}
