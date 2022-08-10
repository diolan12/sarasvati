package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/fatih/color"
)

type StatusBar struct {
	mu      sync.Mutex
	graph   string    // display symbol
	rate    string    // progress bar
	percent int       // percent
	current int       // current progress position
	total   int       // total progress
	workers int       // display symbol
	start   time.Time // start time
	work    []Worker
}

func NewStatusBar(current, total, workers int) *StatusBar {
	bar := new(StatusBar)
	bar.current = current
	bar.total = total
	bar.workers = workers
	bar.start = time.Now()
	if bar.graph == "" {
		bar.graph = "█"
	}
	bar.percent = bar.getPercent()
	for i := 0; i < bar.percent; i += 2 {
		bar.rate += bar.graph // initialize the progress bar position
	}
	return bar
}

func NewStatusBarWithGraph(start, total, workers int, graph string) *StatusBar {
	bar := NewStatusBar(start, total, workers)
	bar.graph = graph
	return bar
}
func (bar *StatusBar) getPercent() int {
	return int((float64(bar.current) / float64(bar.total)) * 100)
}
func (bar *StatusBar) getTime() (s string) {
	u := time.Now().Sub(bar.start).Seconds()
	h := int(u) / 3600
	m := int(u) % 3600 / 60
	if h > 0 {
		s += strconv.Itoa(h) + "h "
	}
	if h > 0 || m > 0 {
		s += strconv.Itoa(m) + "m "
	}
	s += strconv.Itoa(int(u)%60) + "s"
	return
}
func (bar *StatusBar) load() {
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.graph
	}
	cyan := color.New(color.FgCyan, color.Bold)
	// d.Printf("\r[%-50s]% 3d%%    %2s   %d/%d", bar.rate, bar.percent, bar.getTime(), bar.current, bar.total)
	// fmt.Printf("\r[%-50s]% 3d%%    %2s   %d/%d", bar.rate, bar.percent, bar.getTime(), bar.current, bar.total)
	// fmt.Printf("\r[%-50s]% 3d%%    %2s   %d/%d", bar.rate, bar.percent, bar.getTime(), bar.current, bar.total)

	// fmt.Printf("\r\033[A\r")
	fmt.Printf("\r\033[A\r")
	for i := 1; i <= bar.workers; i++ {
		fmt.Printf("\r\033[A\r")
	}

	// fmt.Printf("\r[Sarasvati bot dispatcher]\n")
	cyan.Printf("\rTotal worker: %d      Time elapsed: %2s      Jobs: %d/%d \n", bar.workers, bar.getTime(), bar.current, bar.total)
	for _, w := range bar.work { // i := 1; i <= bar.workers; i++ {
		fmt.Printf("\r[worker %d]: %s   \n", w.id, w.status)
	}

}

func (bar *StatusBar) Run() {
	curr := 0
	for !(curr == bar.total) {
		time.Sleep(time.Millisecond * 1)
		// fmt.Printf("\n\ncurrent %d!=%d %v", curr, bar.total, !(curr == bar.total))
		// bar.load()

		if curr == bar.total {
			break
		}
		curr = bar.current
	}
	// time.Sleep(time.Millisecond * 1000)
	// bar.load()
	// fmt.Printf("\n\n%d", curr)
}
func (bar *StatusBar) Reset(current int) {
	bar.mu.Lock()
	defer bar.mu.Unlock()
	bar.current = current
	// bar.load()

}
func (bar *StatusBar) Refresh() {
	// bar.mu.Lock()
	// defer bar.mu.Unlock()
	bar.load()
}
func (bar *StatusBar) Add(i int, worker Worker) {
	bar.mu.Lock()
	defer bar.mu.Unlock()
	bar.current += i
	bar.work[worker.index] = worker
	bar.load()
}
func (bar *StatusBar) AddWorkers(workers []Worker) {
	bar.mu.Lock()
	defer bar.mu.Unlock()
	bar.work = workers
	// bar.load()
}
func startStatusBar() {

	// b := NewStatusBar(0, 1000, 3)
	// for i := 0; i < 1000; i++ {
	// b.Add(1)
	// 	time.Sleep(time.Millisecond * 10)
	// }

}
