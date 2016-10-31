package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/maxmclau/gput"
)

var colors = map[string]int{
	"blue":        39,
	"red":         9,
	"yellow":      11,
	"green":       10,
	"medium_grey": 240,
}

func main() {
	go trap_and_tidy()
	loop(get_input())
}

func get_input() (string, []string) {
	s := bufio.NewScanner(os.Stdin)

	gput.Setaf(colors["blue"])
	fmt.Printf("Who sucks?")

	gput.Setaf(colors["medium_grey"])
	fmt.Printf(" [Frank] ")

	gput.Sgr0()
	s.Scan()
	sucker := strings.TrimSpace(s.Text())

	if len(sucker) < 1 {
		sucker = "Frank"
	}

	gput.Setaf(colors["blue"])
	fmt.Printf("What does %s suck? ", sucker)

	gput.Sgr0()
	s.Scan()
	raw := strings.Split(s.Text(), ",")

	var suckables []string
	for _, str := range raw {
		suckables = append(suckables, strings.TrimSpace(str))
	}
	return sucker, suckables
}

func get_padding(spaces *int, is_decreasing *bool) string {
	var b bytes.Buffer

	for i := 1; i <= *spaces; i++ {
		b.WriteString(" ")
	}

	if *is_decreasing {
		*spaces--
	} else {
		*spaces++
	}

	if 5 == *spaces || 0 == *spaces {
		*is_decreasing = !*is_decreasing
	}

	return b.String()
}

func loop(sucker string, suckables []string) {
	spaces := 0
	is_decreasing := false

	for {
		for _, suckable := range suckables {
			// Need to figure out how to build a multicolored string before printing
			gput.Setaf(colors["green"])
			fmt.Printf(get_padding(&spaces, &is_decreasing))
			fmt.Printf("%s", sucker)
			gput.Setaf(colors["yellow"])
			fmt.Printf(" sucks ")
			gput.Setaf(colors["red"])
			fmt.Printf("%s\n", suckable)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func trap_and_tidy() {
	signals := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signals
		done <- true
	}()

	<-done

	gput.Setaf(colors["blue"])
	fmt.Printf("\n\nThank you for playing 'Frank Sucks'\n")
	fmt.Printf("Resetting colors...\n\n")
	gput.Sgr0()
	fmt.Printf("Goodbye!\n\n")
	os.Exit(0)
}
