package main

import (
	lg "github.com/charmbracelet/lipgloss"
)

const (
	BLACK  = 0
	RED    = 1
	GREEN  = 2
	PURPLE = 5
	CYAN   = 6
	WHITE  = 7
	GRAY   = 8
)

var (
	red    = lg.ANSIColor(RED)
	cyan   = lg.ANSIColor(CYAN)
	green  = lg.ANSIColor(GREEN)
	purple = lg.ANSIColor(PURPLE)
	white  = lg.ANSIColor(WHITE)
	gray   = lg.ANSIColor(GRAY)

	titleMargin = lg.NewStyle().
			Margin(1, 0, 1, 2)

	titleStyle = lg.NewStyle().
			Background(red).
			Foreground(white).
			Bold(true).
			Padding(0, 1)

	phoneticStyle = lg.NewStyle().
			Foreground(cyan).
			Italic(true)

	posStyle = lg.NewStyle().
			Foreground(purple).
			Italic(true).
			Bold(true).
			MarginLeft(2)

	textStyle = lg.NewStyle().
			MarginLeft(2)

	synonymStyle = lg.NewStyle().
			Foreground(green)

	antonymStyle = lg.NewStyle().
			Foreground(red)

	exampleStyle = lg.NewStyle().
			Foreground(gray).
			Italic(true).
			MarginLeft(2)

	errorStyle = lg.NewStyle().
			Foreground(red)
)
