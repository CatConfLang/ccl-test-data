package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m tuiModel) View() string {
	if m.loadError != nil {
		errStyle := lipgloss.NewStyle().Foreground(errorColor).Bold(true)
		return errStyle.Render(fmt.Sprintf("Error loading %s: %v\n\nPress q to quit.", m.filename, m.loadError))
	}
	if len(m.tests) == 0 {
		return fmt.Sprintf("Loading... (tests=%d, suite=%s, filename=%s)", len(m.tests), m.suite.Suite, m.filename)
	}

	var content strings.Builder

	header := fmt.Sprintf("📋 %s", m.suite.Suite)
	headerStyle := lipgloss.NewStyle().Foreground(primaryColor).Bold(true)
	content.WriteString(headerStyle.Render(header) + "\n")

	content.WriteString(m.renderFilterBar() + "\n")

	// Calculate pane dimensions
	availableWidth := m.width - 4
	listWidth := listPaneWidth
	if listWidth > availableWidth/2 {
		listWidth = availableWidth / 2
	}
	detailWidth := availableWidth - listWidth - 3

	availableHeight := m.height - 5
	if availableHeight < 5 {
		availableHeight = 5
	}

	listPane := m.renderListPane(listWidth, availableHeight)
	detailPane := m.renderDetailPane(detailWidth, availableHeight)

	splitView := lipgloss.JoinHorizontal(lipgloss.Top, listPane, " ", detailPane)
	content.WriteString(splitView + "\n")

	content.WriteString(m.renderHelpBar())

	return content.String()
}

func (m tuiModel) renderFilterBar() string {
	if m.filterActive {
		modeLabel := ""
		switch m.filterMode {
		case FilterName:
			modeLabel = "Name"
		case FilterFunction:
			modeLabel = "Function"
		case FilterBehavior:
			modeLabel = "Behavior"
		case FilterFeature:
			modeLabel = "Feature"
		}
		return filterLabelStyle.Render(fmt.Sprintf("🔍 Filter by %s: ", modeLabel)) +
			activeFilterStyle.Render(m.filterText+"_")
	}

	if m.filterMode != FilterNone && m.filterText != "" {
		modeLabel := ""
		switch m.filterMode {
		case FilterName:
			modeLabel = "name"
		case FilterFunction:
			modeLabel = "function"
		case FilterBehavior:
			modeLabel = "behavior"
		case FilterFeature:
			modeLabel = "feature"
		}
		return filterLabelStyle.Render(fmt.Sprintf("🔍 Filtering by %s: ", modeLabel)) +
			activeFilterStyle.Render(m.filterText) +
			inactiveFilterStyle.Render(fmt.Sprintf(" (%d/%d tests)", len(m.filteredTests), len(m.tests)))
	}

	return inactiveFilterStyle.Render("/ name • f function • b behavior • F feature • c clear")
}

func (m tuiModel) renderListPane(width, height int) string {
	var content strings.Builder

	maxContentLines := height - 2
	if maxContentLines < 1 {
		maxContentLines = 1
	}

	hasMoreBelow := false
	visibleItems := maxContentLines

	startIdx := m.listScroll
	endIdx := startIdx + visibleItems
	if endIdx > len(m.filteredTests) {
		endIdx = len(m.filteredTests)
	}

	// Show scroll indicator at top if needed (takes 1 line)
	if startIdx > 0 {
		scrollUp := lipgloss.NewStyle().Foreground(subtleColor).Render("↑ more above")
		content.WriteString(scrollUp + "\n")
		visibleItems--
		endIdx = startIdx + visibleItems
		if endIdx > len(m.filteredTests) {
			endIdx = len(m.filteredTests)
		}
	}

	// Reserve 1 line for bottom scroll indicator if needed
	if endIdx < len(m.filteredTests) {
		hasMoreBelow = true
		visibleItems--
		endIdx = startIdx + visibleItems
		if endIdx > len(m.filteredTests) {
			endIdx = len(m.filteredTests)
		}
	}

	for i := startIdx; i < endIdx; i++ {
		test := m.filteredTests[i]
		isSelected := i == m.currentTest

		prefix := "  "
		if isSelected {
			prefix = "► "
		}

		badge := ""
		switch test.Validation {
		case "parse":
			badge = "P"
		case "parse_indented":
			badge = "I"
		default:
			if len(test.Validation) > 0 {
				badge = strings.ToUpper(test.Validation[:1])
			}
		}

		maxNameLen := width - 8
		if maxNameLen < 10 {
			maxNameLen = 10
		}
		name := test.Name
		nameRunes := []rune(name)
		if len(nameRunes) > maxNameLen {
			name = string(nameRunes[:maxNameLen-2]) + ".."
		}

		line := fmt.Sprintf("%s[%s] %s", prefix, badge, name)

		if isSelected {
			style := selectedFileStyle.Width(width - 4)
			content.WriteString(style.Render(line) + "\n")
		} else {
			content.WriteString(line + "\n")
		}
	}

	if hasMoreBelow {
		remaining := len(m.filteredTests) - endIdx
		scrollDown := lipgloss.NewStyle().Foreground(subtleColor).Render(fmt.Sprintf("↓ %d more below", remaining))
		content.WriteString(scrollDown)
	}

	paneStyle := listPaneStyle.Width(width).Height(height)
	if m.focusPane == 0 {
		paneStyle = paneStyle.BorderForeground(primaryColor)
	}

	return paneStyle.Render(content.String())
}

func (m tuiModel) renderDetailPane(width, height int) string {
	maxContentLines := height - 2
	if maxContentLines < 1 {
		maxContentLines = 1
	}

	if len(m.filteredTests) == 0 {
		paneStyle := detailPaneStyle.Width(width).Height(height)
		return paneStyle.Render("No tests match the current filter")
	}

	if m.currentTest >= len(m.filteredTests) {
		return detailPaneStyle.Width(width).Height(height).Render("No test selected")
	}

	test := m.filteredTests[m.currentTest]
	var content strings.Builder

	testHeader := fmt.Sprintf("%s  %s", test.Name, getValidationBadge(test.Validation))
	content.WriteString(testHeaderStyle.Render(testHeader) + "\n\n")

	if test.SourceTest != "" && test.SourceTest != test.Name {
		sourceRef := lipgloss.NewStyle().Foreground(subtleColor).Render(fmt.Sprintf("from: %s", test.SourceTest))
		content.WriteString(sourceRef + "\n\n")
	}

	content.WriteString(inputHeaderStyle.Render("📄 INPUT:") + "\n")
	inputText := ""
	if len(test.Inputs) > 0 {
		inputText = test.Inputs[0]
	}
	inputStyle := inputContentStyle.Width(width - 6)
	content.WriteString(inputStyle.Render(formatInputContent(inputText)) + "\n\n")

	content.WriteString(m.renderExpectedOutput(test, width-4) + "\n")

	content.WriteString(m.renderMetadata(test))

	// Truncate content to fit within available height
	contentStr := content.String()
	lines := strings.Split(contentStr, "\n")
	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	if len(lines) > maxContentLines {
		lines = lines[:maxContentLines]
	}

	paneStyle := detailPaneStyle.Width(width).Height(height)
	if m.focusPane == 1 {
		paneStyle = paneStyle.BorderForeground(primaryColor)
	}

	return paneStyle.Render(strings.Join(lines, "\n"))
}

func (m tuiModel) renderExpectedOutput(test TestCase, width int) string {
	var content strings.Builder
	scrollStyle := lipgloss.NewStyle().Foreground(subtleColor)

	result := renderExpectedContent(test, m.entryScroll, maxEntriesDisplay)

	if result.IsError {
		content.WriteString(errorHeaderStyle.Render(result.Header) + "\n")
		return content.String()
	}
	content.WriteString(successHeaderStyle.Render(result.Header) + "\n")

	if !result.IsObject && !result.IsList {
		content.WriteString(fmt.Sprintf("Count: %d\n", result.TotalItems))
	}
	if result.IsList {
		content.WriteString(fmt.Sprintf("Items: %d\n", result.TotalItems))
	}

	if result.HasMoreAbove {
		content.WriteString(scrollStyle.Render("↑ more above\n"))
	}

	if (result.IsObject || result.IsList) && len(result.Lines) > 0 {
		var boxContent strings.Builder
		for i := result.StartIdx; i < result.EndIdx; i++ {
			boxContent.WriteString(result.Lines[i])
			if i < result.EndIdx-1 {
				boxContent.WriteString("\n")
			}
		}
		content.WriteString(objectBoxStyle.Render(boxContent.String()) + "\n")
	} else {
		for i := result.StartIdx; i < result.EndIdx; i++ {
			content.WriteString(result.Lines[i] + "\n")
		}
	}

	if result.HasMoreBelow {
		remaining := result.TotalItems - result.EndIdx
		content.WriteString(scrollStyle.Render(fmt.Sprintf("↓ %d more (tab→detail, h/l scroll)\n", remaining)))
	}

	return content.String()
}

func (m tuiModel) renderMetadata(test TestCase) string {
	var content strings.Builder

	if len(test.Behaviors) > 0 {
		content.WriteString(metaHeaderStyle.Render("Behaviors: "))
		for i, b := range test.Behaviors {
			if i > 0 {
				content.WriteString(", ")
			}
			content.WriteString(tagStyle.Render(b))
		}
		content.WriteString("\n")
	}

	if len(test.Features) > 0 {
		content.WriteString(metaHeaderStyle.Render("Features: "))
		for i, f := range test.Features {
			if i > 0 {
				content.WriteString(", ")
			}
			content.WriteString(tagStyle.Render(f))
		}
		content.WriteString("\n")
	}

	if len(test.Variants) > 0 {
		content.WriteString(metaHeaderStyle.Render("Variants: "))
		for i, v := range test.Variants {
			if i > 0 {
				content.WriteString(", ")
			}
			content.WriteString(tagStyle.Render(v))
		}
		content.WriteString("\n")
	}

	return content.String()
}

func (m tuiModel) renderHelpBar() string {
	var parts []string

	if m.focusPane == 0 {
		parts = append(parts, "j/k:navigate tests")
	} else {
		parts = append(parts, "j/k:scroll entries")
	}
	parts = append(parts, "g/G:first/last")
	parts = append(parts, "tab:switch pane")

	if m.filterMode != FilterNone {
		parts = append(parts, "c:clear filter")
	} else {
		parts = append(parts, "/:filter")
	}

	if m.directory != "" {
		parts = append(parts, "esc:back")
	}
	parts = append(parts, "q:quit")

	paneIndicator := "◀list"
	if m.focusPane == 1 {
		paneIndicator = "detail▶"
	}
	testNum := 0
	if len(m.filteredTests) > 0 {
		testNum = m.currentTest + 1
	}
	countInfo := fmt.Sprintf("[%d/%d] %s", testNum, len(m.filteredTests), paneIndicator)
	if len(m.filteredTests) != len(m.tests) {
		countInfo = fmt.Sprintf("[%d/%d of %d] %s", testNum, len(m.filteredTests), len(m.tests), paneIndicator)
	}

	help := strings.Join(parts, " • ") + " " + summaryStyle.Render(countInfo)
	return suiteInfoStyle.Render(help)
}
