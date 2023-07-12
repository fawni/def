package main

import (
	lg "github.com/charmbracelet/lipgloss"
)

const (
	RED   = "#F28FAD"
	MAUVE = "#DDB6F2"
	SKY   = "#89DCEB"
	GREEN = "#ABE9B3"
	WHITE = "#F8F8F0"
)

var (
	redColor    = lg.Color(RED)
	skyColor    = lg.Color(SKY)
	greenColor  = lg.Color(GREEN)
	mauveColor  = lg.Color(MAUVE)
	whiteColor  = lg.Color(WHITE)
	dimmedColor = lg.AdaptiveColor{Light: "#C3BAC6", Dark: "#6E6C7E"}

	titleMargin = lg.NewStyle().
			Margin(1, 0, 1, 2)

	titleStyle = lg.NewStyle().
			Background(redColor).
			Foreground(whiteColor).
			Bold(true).
			Padding(0, 1)

	phoneticStyle = lg.NewStyle().
			Foreground(skyColor).
			Italic(true)

	posStyle = lg.NewStyle().
			Foreground(mauveColor).
			Italic(true).
			Bold(true).
			MarginLeft(2)

	textStyle = lg.NewStyle().
			MarginLeft(2)

	synonymStyle = lg.NewStyle().
			Foreground(greenColor)

	antonymStyle = lg.NewStyle().
			Foreground(redColor)

	exampleStyle = lg.NewStyle().
			Foreground(dimmedColor).
			Italic(true).
			MarginLeft(2)

	errorStyle = lg.NewStyle().
			Foreground(redColor).
			Bold(true)
)
