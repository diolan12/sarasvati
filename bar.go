package main

import (
	"strconv"
	"sync"
	"time"

	"github.com/fatih/color"
)

type Bar struct {
	mu      sync.Mutex
	graph   string    // display symbol
	rate    string    // progress bar
	percent int       // percent
	current int       // current progress position
	total   int       // total progress
	start   time.Time // start time
}

func NewBar(current, total int) *Bar {
	bar := new(Bar)
	bar.current = current
	bar.total = total
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

func NewBarWithGraph(start, total int, graph string) *Bar {
	bar := NewBar(start, total)
	bar.graph = graph
	return bar
}
func (bar *Bar) getPercent() int {
	return int((float64(bar.current) / float64(bar.total)) * 100)
}
func (bar *Bar) getTime() (s string) {
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
func (bar *Bar) load() {
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.graph
	}
	d := color.New(color.FgGreen, color.Bold)
	d.Printf("\r[%-50s]% 3d%%    %2s   %d/%d", bar.rate, bar.percent, bar.getTime(), bar.current, bar.total)
	// fmt.Printf("\r[%-50s]% 3d%%    %2s   %d/%d", bar.rate, bar.percent, bar.getTime(), bar.current, bar.total)
}
func (bar *Bar) Reset(current int) {
	bar.mu.Lock()
	defer bar.mu.Unlock()
	bar.current = current
	bar.load()

}

func (bar *Bar) Add(i int) {
	bar.mu.Lock()
	defer bar.mu.Unlock()
	bar.current += i
	bar.load()
}
func startBar() {

	b := NewBar(0, 1000)
	for i := 0; i < 1000; i++ {
		b.Add(1)
		time.Sleep(time.Millisecond * 10)
	}

}
