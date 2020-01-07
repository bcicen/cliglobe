package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const framesep = "\x1b[H"
const color = "\033[38;2;127;233;162m"
const reset = "\033[0m"
const clear = "\033[H\033[J"

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

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func main() {
	//lines, err := readLines("./ascsaver.art/globe.vt")
	//if err != nil {
	//panic(err)
	//}

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
				fmt.Printf("   ")
				fmt.Println(line)
			}
			time.Sleep(150 * time.Millisecond)
		}
	}
}
