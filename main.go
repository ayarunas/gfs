package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	go trap_and_tidy()
	loop(get_input())
}

func get_input() (string, []string) {
	s := bufio.NewScanner(os.Stdin)

	fmt.Printf("Who sucks? [Frank] ")
	s.Scan()
	sucker := strings.TrimSpace(s.Text())

	if len(sucker) < 1 {
		sucker = "Frank"
	}

	fmt.Printf("What does %s suck? ", sucker)
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
			fmt.Printf("%s sucks %s\n", sucker, suckable)
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
	fmt.Printf("\n\nThank you for playing 'Frank Sucks'\n")
	fmt.Printf("Resetting colors...\n\nGoodbye!\n")
	os.Exit(0)
}
