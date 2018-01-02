package main

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

type snowflake struct {
	x, y, speedx, speedy int
}

var snow []*snowflake
var width, height int

const (
	scale = 100
	num   = 100
	speed = 15
)

func main() {
	err := termbox.Init()
	termbox.SetOutputMode(termbox.OutputGrayscale)
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	width, height = termbox.Size()
	ticker := time.NewTicker(speed * time.Millisecond)
	snow = make([]*snowflake, num)
	for i := range snow {
		snow[i] = &snowflake{
			rand.Intn(width * scale),
			rand.Intn(height * scale),
			rand.Intn(3) + 1,
			rand.Intn(20) + 10}
	}
	go tick(ticker.C)
	termbox.PollEvent()
}

func tick(ch <-chan time.Time) {
	for {
		select {
		case <-ch:
			termbox.Clear(termbox.ColorDefault, termbox.Attribute(0))
			for _, sf := range snow {
				if sf.y > height*scale {
					sf.y = 0
				} else if sf.x > width*scale {
					sf.x = 0
				} else {
					sf.x += sf.speedx
					sf.y += sf.speedy
				}
				termbox.SetCell(sf.x/scale, sf.y/scale, '.', termbox.Attribute(sf.speedy+sf.speedx), termbox.ColorDefault)
			}
			termbox.Flush()
		}
	}
}
