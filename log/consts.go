package log

import "github.com/fatih/color"

var (
	red  = color.New(color.FgRed).SprintFunc()
	info = color.New(color.BgBlue, color.FgWhite, color.Bold).SprintFunc()
	// USF - underline sprint func
	USF = color.New(color.Underline).SprintFunc()
)
