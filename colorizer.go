package main

import (
	"fmt"
	"math/rand"
	"time"
)

// EscapeCode returns a new escape code with RGB values
func EscapeCode(r, g, b uint8) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

// Colorizer returns an ASCII escape code
// for colorizing a string
type Colorizer interface {
	Next() string
	Base() string
}

// Gradiant implements a colorizer that returns a
type Gradiant struct {
	next   int
	colors [][3]uint8
}

// NewGradiant returns a new gradiant colorizer
func NewGradiant(steps int, start, end [3]uint8) *Gradiant {
	var colors [][3]uint8
	factor := 1 / (float64(steps - 1))
	for i := 0; i < steps; i++ {
		f := factor * float64(i)
		r, rMax := float64(start[0]), float64(end[0])
		g, gMax := float64(start[1]), float64(end[1])
		b, bMax := float64(start[2]), float64(end[2])
		colors = append(colors, [3]uint8{
			uint8(r + f*(rMax-r)),
			uint8(g + f*(gMax-g)),
			uint8(b + f*(bMax-b)),
		})
	}
	// add the same colors back in reverse for
	// the day-to-night effect
	for i := len(colors) - 1; i > 0; i-- {
		fmt.Println(i)
		colors = append(colors, colors[i])
	}
	return &Gradiant{colors: colors}
}

func (g *Gradiant) Next() string {
	if g.next+1 == len(g.colors) {
		g.next = 0
	}
	next := EscapeCode(g.colors[g.next][0], g.colors[g.next][1], g.colors[g.next][2])
	g.next++
	return next

}

func (g *Gradiant) Base() string {
	return EscapeCode(g.colors[0][0], g.colors[0][1], g.colors[0][2])
}

// Random implements a randomized Colorizer
type Random struct{}

func (r Random) Next() string {
	return EscapeCode(uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)))
}

func (r Random) Base() string {
	return EscapeCode(uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)))
}

func init() {
	rand.Seed(time.Now().Unix())
}
