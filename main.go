package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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
		}
	}
}
