package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
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

	var (
		rate = flag.String("rate", "100ms", "globe rotation rate")
	)

	flag.Parse()

	duration, err := time.ParseDuration(*rate)
	if err != nil {
		fmt.Printf("bad duration: %s\n", *rate)
		os.Exit(1)
	}

	var (
		sigCh = make(chan os.Signal)
		strCh = make(chan string)
	)

	go func() {
		for {
			select {
			case s := <-strCh:
				fmt.Printf(s)
			case <-sigCh:
				fmt.Println(reset)
				fmt.Println(clear)
				os.Exit(0)
			}
		}
	}()

	signal.Notify(sigCh, os.Interrupt)

	var ng int
	for {
		for _, frame := range globe {
			strCh <- clear
			for _, line := range frame {
				// apply color
				strCh <- gradiant[ng]
				ng++
				if ng >= len(gradiant) {
					ng = 0
				}
				strCh <- lpadding + line + "\n"
			}
			time.Sleep(duration)
		}
	}
}
