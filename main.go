package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	fmt.Printf("Who sucks? ")
	s.Scan()
	sucker := strings.TrimSpace(s.Text())

	fmt.Printf("What does %s suck? ", sucker)
	s.Scan()
	suckables := strings.Split(s.Text(), ",")

	var trimmed []string
	for _, str := range suckables {
		trimmed = append(trimmed, strings.TrimSpace(str))
	}

	for _, suckable := range trimmed {
		fmt.Printf("%s sucks %s\n", sucker, suckable)
	}
}
