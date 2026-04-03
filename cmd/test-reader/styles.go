package main

import "github.com/charmbracelet/lipgloss"

// Configuration constants
const (
	maxEntriesDisplay = 6  // Maximum entries to show before scrolling/truncation
	listPaneWidth     = 40 // Width of the test list pane in split view
)

// Color palette
var (
	primaryColor = lipgloss.Color("#00D7FF")
	successColor = lipgloss.Color("#00D787")
	errorColor   = lipgloss.Color("#FF5F87")
	warningColor = lipgloss.Color("#FFAF00")
	subtleColor  = lipgloss.Color("#626262")
)

// Application styles
var (
	suiteHeaderStyle = lipgloss.NewStyle().
				Foreground(primaryColor).
				Bold(true).
				Border(lipgloss.DoubleBorder()).
				BorderForeground(primaryColor).
				Padding(1, 2).
				Margin(1, 0)

	suiteInfoStyle = lipgloss.NewStyle().
			Foreground(subtleColor).
			Italic(true).
			Margin(0, 2)

	testHeaderStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			Margin(1, 0, 0, 0)

	inputHeaderStyle = lipgloss.NewStyle().
				Foreground(primaryColor).
				Bold(true)

	inputContentStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(lipgloss.Color("#1A1A1A")).
				Padding(0, 1).
				Margin(0, 0, 1, 2)

	successHeaderStyle = lipgloss.NewStyle().
				Foreground(successColor).
				Bold(true)

	errorHeaderStyle = lipgloss.NewStyle().
				Foreground(errorColor).
				Bold(true)

	metaHeaderStyle = lipgloss.NewStyle().
			Foreground(warningColor).
			Bold(true)

	tagStyle = lipgloss.NewStyle().
			Foreground(warningColor).
			Bold(true)

	entryKeyStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true)

	entryValueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF"))

	entryEqualsStyle = lipgloss.NewStyle().
				Foreground(warningColor).
				Bold(true)

	entryBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(subtleColor).
			Padding(0, 1).
			Margin(0, 0, 0, 2)

	emptyKeyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#999999")).
			Italic(true)

	emptyValueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#999999")).
			Italic(true)

	summaryStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			Margin(1, 0)

	// File selection styles
	selectedFileStyle = lipgloss.NewStyle().
				Foreground(primaryColor).
				Background(lipgloss.Color("#1A1A1A")).
				Bold(true).
				Padding(0, 1)

	fileListStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(subtleColor).
			Padding(1).
			Margin(1, 0)

	fileHeaderStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			Align(lipgloss.Center).
			Border(lipgloss.DoubleBorder()).
			BorderForeground(primaryColor).
			Padding(1, 2).
			Margin(1, 0)

	// Split view styles
	listPaneStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(subtleColor).
			Padding(0, 1)

	detailPaneStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(primaryColor).
			Padding(0, 1)

	filterLabelStyle = lipgloss.NewStyle().
				Foreground(warningColor).
				Bold(true)

	activeFilterStyle = lipgloss.NewStyle().
				Foreground(primaryColor).
				Background(lipgloss.Color("#1A1A1A")).
				Bold(true).
				Padding(0, 1)

	inactiveFilterStyle = lipgloss.NewStyle().
				Foreground(subtleColor).
				Padding(0, 1)

	// Validation type badge styles
	parseValidationStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#000000")).
				Background(successColor).
				Bold(true).
				Padding(0, 1)

	parseIndentedValidationStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#000000")).
					Background(primaryColor).
					Bold(true).
					Padding(0, 1)

	otherValidationStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#000000")).
				Background(warningColor).
				Bold(true).
				Padding(0, 1)

	// Object visualization styles (for build_hierarchy, load, etc.)
	objectKeyStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true)

	objectValueStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF"))

	objectBranchStyle = lipgloss.NewStyle().
				Foreground(subtleColor)

	arrayIndexStyle = lipgloss.NewStyle().
			Foreground(subtleColor)

	objectBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(subtleColor).
			Padding(0, 1).
			Margin(0, 0, 0, 2)
)
