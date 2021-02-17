package main

import (
	"log"
	"sync"
)

func main() {
	m := map[int]int{}
	mux := &sync.RWMutex{}

	go writeLoop(m, mux)
	go readLoop(m, mux)
	go readLoop(m, mux)
	go readLoop(m, mux)
	go readLoop(m, mux)

	block := make(chan struct{})
	<-block
}

func writeLoop(m map[int]int, mux *sync.RWMutex) {
	for {
		log.Println("================== WRITING ==================")

		for i := 0; i < 10; i++ {
			mux.Lock()
			m[i] = i
			mux.Unlock()
		}
	}
}

func readLoop(m map[int]int, mux *sync.RWMutex) {
	for {
		mux.RLock()
		for k, v := range m {
			log.Println(k, "-", v)
		}
		mux.RUnlock()
	}
}
