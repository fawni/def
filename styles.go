package main

import (
	"github.com/charmbracelet/lipgloss"
	lg "github.com/charmbracelet/lipgloss"
)

const (
	red      = "#F28FAD"
	flamingo = "#F2CDCD"
	mauve    = "#DDB6F2"
	sky      = "#89DCEB"
	green    = "#ABE9B3"
	white    = "#F8F8F0"
)

var (
	redColor    = lg.Color(red)
	skyColor    = lg.Color(sky)
	greenColor  = lg.Color(green)
	mauveColor  = lg.Color(mauve)
	whiteColor  = lg.Color(white)
	dimmedColor = lipgloss.AdaptiveColor{Light: "#C3BAC6", Dark: "#6E6C7E"}

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
