package main

import (
	"fmt"
	"time"
)

type StatefulWorker struct {
	ch chan int
}

func (w *StatefulWorker) Start() {
	go func(){
		var count int

		for value := range w.ch {
			count += value
			fmt.Println("Actual value:", count)
		}

		fmt.Println("Exit")
	}()
}

func (w *StatefulWorker) Send(count int) {
	w.ch <- count
}

func (w *StatefulWorker) Stop() {
	close(w.ch)
}

func main(){
	sw := &StatefulWorker{
		ch: make(chan int),
	}
	sw.Start()

	for i := range 5 {
		sw.Send(i)
		time.Sleep(200 * time.Millisecond)
	}

	sw.Stop()
	time.Sleep(time.Second)
}