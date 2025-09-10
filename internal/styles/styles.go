package styles

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

var (
	// Color palette
	Blue   = lipgloss.Color("#74C0FC")
	Green  = lipgloss.Color("#51CF66")
	Yellow = lipgloss.Color("#FFD43B")
	Red    = lipgloss.Color("#FF6B6B")
	Cyan   = lipgloss.Color("#339AF0")
)

var (
	// Base styles
	info     = lipgloss.NewStyle().Foreground(Blue).Bold(true)
	success  = lipgloss.NewStyle().Foreground(Green).Bold(true)
	warning  = lipgloss.NewStyle().Foreground(Yellow).Bold(true)
	error_   = lipgloss.NewStyle().Foreground(Red).Bold(true)
	infoLite = lipgloss.NewStyle().Foreground(Cyan)
)

// Direct printing functions - single function call, no fmt.Println needed
func Info(format string, args ...interface{}) {
	fmt.Println(info.Render(fmt.Sprintf(format, args...)))
}

func Success(format string, args ...interface{}) {
	fmt.Println(success.Render(fmt.Sprintf(format, args...)))
}

func Warning(format string, args ...interface{}) {
	fmt.Println(warning.Render(fmt.Sprintf(format, args...)))
}

func Error(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s\n", error_.Render(fmt.Sprintf(format, args...)))
}

func InfoLite(format string, args ...interface{}) {
	fmt.Println(infoLite.Render(fmt.Sprintf(format, args...)))
}

// Specialized message functions
func Status(emoji, message string) {
	Info("%s %s", emoji, message)
}

func Command(command string) {
	InfoLite("📋 Running: %s", command)
}

func FileProcessed(filename string) {
	InfoLite("✓ Generated tests for %s", filename)
}