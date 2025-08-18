package luna

import "github.com/charmbracelet/lipgloss"

var fg = lipgloss.Color("#ffffff")
var bg = lipgloss.Color("#000000")

var nameStyle = lipgloss.NewStyle().
	Border(lipgloss.ThickBorder()).
	BorderForeground(fg).
	Padding(0, 3, 0, 3).
	Foreground(fg)
