package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"
)

const (
	framesep  = "\x1b[H"
	reset     = "\033[0m"
	clear     = "\033[H\033[J"
	lpadding  = "   "
	tspadding = "        "
	tsfmt     = "2006-01-02 15:04:05.00000 -0700 MST "
)

func hexToRGB(h string) (rgb [3]uint8) {
	h = strings.Replace(h, "#", "", -1)

	var hb []byte
	hb, err := hex.DecodeString(h)
	if err != nil {
		fmt.Printf("bad hex code: %s\n", h)
		os.Exit(1)
	}

	if len(hb) > 0 {
		rgb[0] = uint8(hb[0])
	}
	if len(hb) > 1 {
		rgb[1] = uint8(hb[1])
	}
	if len(hb) > 2 {
		rgb[2] = uint8(hb[2])
	}
	return
}

func main() {

	var (
		startColor = flag.String("startColor", "#3f8155", "start hex code for gradiant colorizer")
		endColor   = flag.String("endColor", "#85e8a6", "ending hex code for gradiant colorizer")
		shades     = flag.Int("shades", 5, "number of shades for gradiant colorizer")
		random     = flag.Bool("random", false, "use randomized colors")
		rate       = flag.String("rate", "100ms", "globe rotation rate")
		clock      = flag.Bool("clock", false, "show clock below globe")
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

	var colorizer Colorizer

	if *random {
		colorizer = Random{}
	} else {
		colorizer = NewGradiant(*shades, hexToRGB(*startColor), hexToRGB(*endColor))
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		for _, frame := range globe {
			strCh <- clear + "\n"
			for _, line := range frame {
				// apply color
				strCh <- colorizer.Next()
				strCh <- lpadding + line + "\n"
			}
			if *clock {
				strCh <- colorizer.Base()
				strCh <- "\n" + tspadding + time.Now().Format(tsfmt)
			}
			time.Sleep(duration)
		}
	}
}
