package main

import (
	"fmt"
	"math/rand"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sort"
	"sync"
	"syscall"
	"time"
)

var quit chan struct{}
var wg sync.WaitGroup

type vals struct {
	val1 []int
	val2 []int
}

func main() {
	quit = make(chan struct{})
	wg.Add(2)

	numChan := make(chan vals)
	go pump(numChan)
	go union(numChan)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	fmt.Println("Stopping...")
	close(quit)
	wg.Wait()
	fmt.Println("Done.")
}

func pump(numChan chan<- vals) {
	ticker := time.NewTicker(1000 * time.Millisecond)

	defer func() {
		ticker.Stop()
		wg.Done()
	}()

	for {
		select {
		case <-quit:
			return
		case <-ticker.C:
			vals1 := make([]int, 0, 1000)
			for i := 0; i < 1000; i++ {
				val := rand.Intn(1e6)
				vals1 = append(vals1, val)
			}
			vals2 := make([]int, 0, 1000)
			for i := 0; i < 1000; i++ {
				val := rand.Intn(1e6)
				vals2 = append(vals1, val)
			}
			select {
			case numChan <- vals{vals1, vals2}:
			case <-quit:
				return
			}
		}
	}
}

func union(numChan <-chan vals) {
	for {
		select {
		case v := <-numChan:
			Union(v.val1, v.val2)
		case <-quit:
			wg.Done()
			return
		}
	}
}

func Union(a, b []int) []int {
	maxlen := max(len(a), len(b))
	result := make([]int, 0, maxlen)

	m := make(map[int]bool)
	for _, val := range a {
		if !m[val] {
			result = append(result, val)
			m[val] = true
		}
	}
	for _, val := range b {
		if !m[val] {
			result = append(result, val)
			m[val] = true
		}
	}

	sort.Ints(result)

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
