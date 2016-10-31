package main

import (
	"bufio"
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

func loop(sucker string, suckables []string) {
	for {
		for _, suckable := range suckables {
			gput.Setaf(colors["green"])
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
		fmt.Println(sig)
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
