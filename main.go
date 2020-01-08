package main

import (
	"fmt"
	"time"
)

const (
	framesep = "\x1b[H"
	color    = "\033[38;2;127;233;162m"
	reset    = "\033[0m"
	clear    = "\033[H\033[J"
	lpadding = "   "
)

var gradiant = []string{
	"\033[38;2;133;232;166m",
	"\033[38;2;118;210;149m",
	"\033[38;2;096;181;124m",
	"\033[38;2;084;162;110m",
	"\033[38;2;073;145;097m",
	"\033[38;2;063;129;085m",
	"\033[38;2;063;129;085m",
	"\033[38;2;073;145;097m",
	"\033[38;2;084;162;110m",
	"\033[38;2;096;181;124m",
	"\033[38;2;118;210;149m",
	"\033[38;2;133;232;166m",
}

func main() {
	var ng int
	for {
		for _, frame := range globe {
			fmt.Println(clear)
			for _, line := range frame {
				// apply color
				fmt.Printf(gradiant[ng])
				ng++
				if ng >= len(gradiant) {
					ng = 0
				}
				fmt.Println(lpadding + line)
			}
			time.Sleep(150 * time.Millisecond)
		}
	}
}
