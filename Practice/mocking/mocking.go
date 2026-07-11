package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func Countdown(out io.Writer) {
	countDownTime := 3
	finalWord := "Go!"
	for i := countDownTime; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}