package lib

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func LogConverted(inputPath string, outputPath string) {
	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("70")).Render("Converted " + inputPath + " to " + outputPath))
}

func LogConversionError(inputPath string, text string) {
	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render("Error converting file " + inputPath + " : " + text))
}

func LogInfo(text string) {
	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("70")).Render(text))
}

func LogError(text string) {
	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render(text))
}
