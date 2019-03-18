package main

import czr "github.com/ravindersahni/colorizer"

var colors = map[string]int{
	"blue":        39,
	"red":         9,
	"yellow":      11,
	"green":       10,
	"medium_grey": 240,
}

var blue = czr.Make(colors["blue"])
var red = czr.Make(colors["red"])
var yellow = czr.Make(colors["yellow"])
var green = czr.Make(colors["green"])
var medium_grey = czr.Make(colors["medium_grey"])
