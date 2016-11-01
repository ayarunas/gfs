package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type colorizer func(string) string

func make_colorizer(c int) colorizer {
	c_bytes, err := exec.Command("tput", "setaf", strconv.Itoa(c)).Output()

	if err != nil {
		os.Exit(1)
	}

	return func(s string) string {
		var b bytes.Buffer
		b.WriteString(string(c_bytes))
		b.WriteString(s)
		color_str := b.String()
		b.Reset()
		return color_str
	}
}

var colors = map[string]int{
	"blue":        39,
	"red":         9,
	"yellow":      11,
	"green":       10,
	"medium_grey": 240,
}

var blue = make_colorizer(colors["blue"])
var red = make_colorizer(colors["red"])
var yellow = make_colorizer(colors["yellow"])
var green = make_colorizer(colors["green"])
var medium_grey = make_colorizer(colors["medium_grey"])

func reset_colors() {
	reset, err := exec.Command("tput", "sgr0").Output()

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf(string(reset))
}
