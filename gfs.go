package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/ravindersahni/czr"

	kp "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	wobble := kp.Flag("wobble", "Enable wobbling by setting a max leftpad").Default("0").Short('w').Int()
	kp.Version("0.0.1")
	kp.CommandLine.VersionFlag.Short('v')
	kp.CommandLine.HelpFlag.Short('h')
	kp.Parse()

	go trap_and_tidy()

	name, items := get_input()
	loop(name, items, wobble)
}

func get_input() (string, []string) {
	s := bufio.NewScanner(os.Stdin)
	var sucker string

	for {
		fmt.Printf(blue("Who sucks?"))
		fmt.Printf(medium_grey(" [Frank] "))
		czr.Reset()

		s.Scan()
		sucker = strings.TrimSpace(s.Text())

		match, _ := regexp.MatchString("(?i)r.*a.*v.*i", sucker)
		if match == false {
			break
		}

		fmt.Printf(blue("No, Ravi is awesome\n"))
	}

	if len(sucker) < 1 {
		sucker = "Frank"
	}

	fmt.Printf(blue("What does %s suck? "), sucker)
	czr.Reset()
	s.Scan()
	raw := strings.Split(s.Text(), ",")

	var suckables []string
	for _, str := range raw {
		suckables = append(suckables, strings.TrimSpace(str))
	}
	return sucker, suckables
}

func get_padding(spaces *int, is_decreasing *bool, wobble *int) string {
	if *wobble == 0 {
		return ""
	}

	var b bytes.Buffer
	for i := 1; i <= *spaces; i++ {
		b.WriteString(" ")
	}

	if *is_decreasing {
		*spaces--
	} else {
		*spaces++
	}

	if *wobble == *spaces || 0 == *spaces {
		*is_decreasing = !*is_decreasing
	}

	return b.String()
}

func loop(sucker string, suckables []string, wobble *int) {
	spaces := 0
	is_decreasing := false

	for {
		for _, suckable := range suckables {
			fmt.Printf("%s"+green(sucker)+yellow(" sucks ")+red(suckable)+"\n",
				get_padding(&spaces, &is_decreasing, wobble))
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func trap_and_tidy() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals

	fmt.Printf(blue("\n\nThank you for playing '") + yellow("Frank Sucks") + blue("'\n"))
	fmt.Printf(blue("Resetting colors...\n\n"))
	czr.Reset()
	fmt.Printf("Goodbye!\n\n")
	os.Exit(0)
}
