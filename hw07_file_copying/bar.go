package main

import (
	"fmt"
	"strings"
)

const (
	TealColor  = "\033[1;36m"
	ResetColor = "\033[0m"
)

type Bar struct {
	charFull string
	start    string
	end      string
	percent  int
	total    int64
}

func (bar *Bar) Update(current int64) {
	curPercent := int(current * 100 / bar.total)
	if ((bar.percent + 2) > curPercent) && (current != 0) {
		return
	}

	bar.percent = curPercent

	graph := strings.Repeat(bar.charFull, bar.percent/2)
	fmt.Printf(TealColor+"\r%s%-50s%s"+ResetColor+"%3d%% %10d/%d", bar.start, graph, bar.end, bar.percent, current, bar.total)
}

func (bar *Bar) Finish() {
	fmt.Println()
}

func NewBar(t int64) *Bar {
	return &Bar{
		charFull: "█",
		start:    "[",
		end:      "]",
		percent:  0,
		total:    t,
	}
}
