package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Color palette and styles
var (
	primaryColor   = lipgloss.Color("#00D7FF")
	successColor   = lipgloss.Color("#00D787")
	errorColor     = lipgloss.Color("#FF5F87")
	warningColor   = lipgloss.Color("#FFAF00")
	subtleColor    = lipgloss.Color("#626262")
	
	// Suite header style
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

	// Test case styles
	testHeaderStyle = lipgloss.NewStyle().
		Foreground(primaryColor).
		Bold(true).
		Border(lipgloss.RoundedBorder(), true, false, false, false).
		BorderForeground(primaryColor).
		Padding(0, 1).
		Margin(1, 0, 0, 0)

	descriptionStyle = lipgloss.NewStyle().
		Foreground(warningColor).
		Italic(true).
		Margin(0, 1)

	// Input CCL box
	inputHeaderStyle = lipgloss.NewStyle().
		Foreground(primaryColor).
		Bold(true)

	inputBoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(subtleColor).
		Padding(1).
		Margin(1, 1)

	// Success/Error validation boxes
	successBoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(successColor).
		Padding(1).
		Margin(1, 1)

	errorBoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(errorColor).
		Padding(1).
		Margin(1, 1)

	successLabelStyle = lipgloss.NewStyle().
		Foreground(successColor).
		Bold(true)

	errorLabelStyle = lipgloss.NewStyle().
		Foreground(errorColor).
		Bold(true)

	// Metadata styles
	metaBoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(warningColor).
		Padding(1).
		Margin(1, 1)

	tagStyle = lipgloss.NewStyle().
		Foreground(warningColor).
		Background(lipgloss.Color("#2A2A00")).
		Padding(0, 1).
		Margin(0, 1, 0, 0).
		Bold(true)

	conflictStyle = lipgloss.NewStyle().
		Foreground(errorColor).
		Background(lipgloss.Color("#2A0000")).
		Padding(0, 1).
		Margin(0, 1, 0, 0).
		Bold(true)

	summaryStyle = lipgloss.NewStyle().
		Foreground(primaryColor).
		Bold(true).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(primaryColor).
		Padding(0, 1).
		Margin(1, 0)
)

// TestSuite represents the structure of a test file
type TestSuite struct {
	Schema      string `json:"$schema"`
	Suite       string `json:"suite"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Tests       []Test `json:"tests"`
}

// Test represents a single test case
type Test struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Input       string                 `json:"input"`
	Validations map[string]interface{} `json:"validations"`
	Meta        Meta                   `json:"meta"`
}

// Meta contains test metadata
type Meta struct {
	Tags     []string `json:"tags"`
	Level    int      `json:"level"`
	Feature  string   `json:"feature"`
	Conflicts []string `json:"conflicts,omitempty"`
}

// ParseValidation represents the parse validation structure
type ParseValidation struct {
	Count    int     `json:"count"`
	Expected []Entry `json:"expected,omitempty"`
	Error    bool    `json:"error,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

// Entry represents a key-value pair
type Entry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: test-reader <test-file.json> [--tui]")
		fmt.Println("       test-reader tests/api-essential-parsing.json")
		fmt.Println("       test-reader tests/api-essential-parsing.json --tui")
		os.Exit(1)
	}

	// Check if TUI mode is requested
	useTUI := false
	filename := os.Args[1]
	if len(os.Args) > 2 && os.Args[2] == "--tui" {
		useTUI = true
	}

	if useTUI {
		runTUI(filename)
	} else {
		if err := processTestFile(filename); err != nil {
			log.Printf("Error processing %s: %v", filename, err)
		}
	}
}

func processTestFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	var suite TestSuite
	if err := json.Unmarshal(data, &suite); err != nil {
		return fmt.Errorf("parsing JSON: %w", err)
	}

	// Suite header with styled box
	header := fmt.Sprintf("%s", suite.Suite)
	info := fmt.Sprintf("File: %s | %s", filepath.Base(filename), suite.Description)
	
	fmt.Println(suiteHeaderStyle.Render(header))
	fmt.Println(suiteInfoStyle.Render(info))
	fmt.Println()

	parseOnlyCount := 0
	for _, test := range suite.Tests {
		if hasParseValidation(test) {
			parseOnlyCount++
			displayTest(test, parseOnlyCount)
		}
	}

	// Summary with styled box
	if parseOnlyCount == 0 {
		summary := "📋 No parse-only tests found in this file"
		fmt.Println(summaryStyle.Render(summary))
	} else {
		summary := fmt.Sprintf("📊 Found %d parse-only test(s)", parseOnlyCount)
		fmt.Println(summaryStyle.Render(summary))
	}
	fmt.Println()

	return nil
}

func hasParseValidation(test Test) bool {
	_, hasParse := test.Validations["parse"]
	return hasParse
}

func displayTest(test Test, index int) {
	// Test header
	header := fmt.Sprintf("Test #%d: %s", index, test.Name)
	fmt.Println(testHeaderStyle.Render(header))
	
	if test.Description != "" {
		fmt.Println(descriptionStyle.Render("📝 " + test.Description))
	}

	// Display input in styled box
	inputHeader := inputHeaderStyle.Render("📄 CCL INPUT")
	inputContent := test.Input
	if inputContent == "" {
		inputContent = "(empty)"
	}
	inputBox := inputBoxStyle.Render(inputHeader + "\n\n" + inputContent)
	fmt.Println(inputBox)

	// Display parse validation if present
	if parseData, ok := test.Validations["parse"]; ok {
		displayParseValidation(parseData)
	}

	// Display metadata in styled box
	displayMetadata(test)
	fmt.Println()
}

func displayParseValidation(parseData interface{}) {
	// Convert to map first to handle the interface{}
	parseMap, ok := parseData.(map[string]interface{})
	if !ok {
		errorContent := errorLabelStyle.Render("❌ Invalid parse validation format")
		fmt.Println(errorBoxStyle.Render(errorContent))
		return
	}

	count, _ := parseMap["count"].(float64) // JSON numbers are float64
	
	// Check if this is an error case
	if errorVal, hasError := parseMap["error"]; hasError && errorVal == true {
		var content strings.Builder
		content.WriteString(errorLabelStyle.Render("❌ EXPECTED: Parse Error") + "\n")
		content.WriteString(fmt.Sprintf("Count: %.0f assertion(s)\n", count))
		
		if errorMsg, ok := parseMap["error_message"].(string); ok {
			content.WriteString(fmt.Sprintf("Error: %s", errorMsg))
		}
		
		fmt.Println(errorBoxStyle.Render(content.String()))
		return
	}

	// Handle successful parse case
	var content strings.Builder
	content.WriteString(successLabelStyle.Render("✅ EXPECTED: Parse Success") + "\n")
	content.WriteString(fmt.Sprintf("Count: %.0f assertion(s)\n", count))

	if expectedData, ok := parseMap["expected"].([]interface{}); ok {
		content.WriteString("\nEntries:\n")
		for i, entryData := range expectedData {
			if entryMap, ok := entryData.(map[string]interface{}); ok {
				key, _ := entryMap["key"].(string)
				value, _ := entryMap["value"].(string)
				content.WriteString(fmt.Sprintf("  %d. key=%q, value=%q\n", i+1, key, value))
			}
		}
	}
	
	fmt.Println(successBoxStyle.Render(content.String()))
}

func displayMetadata(test Test) {
	var content strings.Builder
	
	content.WriteString("🏷️  TAGS:\n")
	for _, tag := range test.Meta.Tags {
		content.WriteString(tagStyle.Render(tag))
	}
	content.WriteString("\n\n")
	
	content.WriteString(fmt.Sprintf("📊 LEVEL: %d | FEATURE: %s\n", test.Meta.Level, test.Meta.Feature))
	
	if len(test.Meta.Conflicts) > 0 {
		content.WriteString("\n⚠️  CONFLICTS:\n")
		for _, conflict := range test.Meta.Conflicts {
			content.WriteString(conflictStyle.Render(conflict))
		}
	}
	
	fmt.Println(metaBoxStyle.Render(content.String()))
}

// TUI Implementation
type tuiModel struct {
	tests       []Test
	suite       TestSuite
	filename    string
	currentTest int
	showAll     bool
	width       int
	height      int
}

type testLoadedMsg struct {
	suite    TestSuite
	filename string
}

func initialTUIModel() tuiModel {
	return tuiModel{
		currentTest: 0,
		showAll:     false,
		width:       80,
		height:      24,
	}
}

func loadTestFileCmd(filename string) tea.Cmd {
	return func() tea.Msg {
		data, err := os.ReadFile(filename)
		if err != nil {
			return nil
		}

		var suite TestSuite
		if err := json.Unmarshal(data, &suite); err != nil {
			return nil
		}

		return testLoadedMsg{
			suite:    suite,
			filename: filename,
		}
	}
}

func (m tuiModel) Init() tea.Cmd {
	return nil
}

func (m tuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case testLoadedMsg:
		m.suite = msg.suite
		m.filename = msg.filename
		// Filter to only parse tests
		parseTests := []Test{}
		for _, test := range msg.suite.Tests {
			if _, hasParse := test.Validations["parse"]; hasParse {
				parseTests = append(parseTests, test)
			}
		}
		m.tests = parseTests
		return m, nil

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "j", "down":
			if m.currentTest < len(m.tests)-1 {
				m.currentTest++
			}
		case "k", "up":
			if m.currentTest > 0 {
				m.currentTest--
			}
		case "g":
			m.currentTest = 0
		case "G":
			m.currentTest = len(m.tests) - 1
		case "a":
			m.showAll = !m.showAll
		}
	}
	return m, nil
}

func (m tuiModel) View() string {
	if len(m.tests) == 0 {
		return "Loading..."
	}

	var content strings.Builder

	// Header
	header := fmt.Sprintf("%s", m.suite.Suite)
	info := fmt.Sprintf("File: %s | %s", filepath.Base(m.filename), m.suite.Description)
	
	content.WriteString(suiteHeaderStyle.Render(header) + "\n")
	content.WriteString(suiteInfoStyle.Render(info) + "\n\n")

	if m.showAll {
		// Show all tests
		for i, test := range m.tests {
			if i == m.currentTest {
				content.WriteString(m.renderTest(test, i+1, true) + "\n")
			} else {
				content.WriteString(m.renderTestSummary(test, i+1) + "\n")
			}
		}
	} else {
		// Show current test only
		if m.currentTest < len(m.tests) {
			content.WriteString(m.renderTest(m.tests[m.currentTest], m.currentTest+1, false) + "\n")
		}
	}

	// Navigation info
	navInfo := fmt.Sprintf("Test %d of %d", m.currentTest+1, len(m.tests))
	help := "j/k: navigate • g/G: first/last • a: toggle all • q: quit"
	
	content.WriteString("\n")
	content.WriteString(summaryStyle.Render(navInfo) + "\n")
	content.WriteString(suiteInfoStyle.Render(help))

	return content.String()
}

func (m tuiModel) renderTestSummary(test Test, index int) string {
	prefix := "  "
	if index == m.currentTest+1 {
		prefix = "► "
	}
	
	status := "✅"
	if parseData, ok := test.Validations["parse"].(map[string]interface{}); ok {
		if errorVal, hasError := parseData["error"]; hasError && errorVal == true {
			status = "❌"
		}
	}
	
	summary := fmt.Sprintf("%s%s %s", prefix, status, test.Name)
	if index == m.currentTest+1 {
		return testHeaderStyle.Render(summary)
	}
	return summary
}

func (m tuiModel) renderTest(test Test, index int, compact bool) string {
	var content strings.Builder

	// Test header
	header := fmt.Sprintf("Test #%d: %s", index, test.Name)
	content.WriteString(testHeaderStyle.Render(header) + "\n")
	
	if test.Description != "" {
		content.WriteString(descriptionStyle.Render("📝 " + test.Description) + "\n")
	}

	// Input CCL box
	inputHeader := inputHeaderStyle.Render("📄 CCL INPUT")
	inputContent := test.Input
	if inputContent == "" {
		inputContent = "(empty)"
	}
	inputBox := inputBoxStyle.Render(inputHeader + "\n\n" + inputContent)
	content.WriteString(inputBox + "\n")

	// Parse validation
	if parseData, ok := test.Validations["parse"]; ok {
		content.WriteString(m.renderParseValidation(parseData, compact) + "\n")
	}

	// Metadata box (only if not compact)
	if !compact {
		content.WriteString(m.renderMetadata(test) + "\n")
	}

	return content.String()
}

func (m tuiModel) renderParseValidation(parseData interface{}, compact bool) string {
	parseMap, ok := parseData.(map[string]interface{})
	if !ok {
		errorContent := errorLabelStyle.Render("❌ Invalid parse validation format")
		return errorBoxStyle.Render(errorContent)
	}

	count, _ := parseMap["count"].(float64)
	
	// Check if this is an error case
	if errorVal, hasError := parseMap["error"]; hasError && errorVal == true {
		var content strings.Builder
		content.WriteString(errorLabelStyle.Render("❌ EXPECTED: Parse Error") + "\n")
		content.WriteString(fmt.Sprintf("Count: %.0f assertion(s)\n", count))
		
		if errorMsg, ok := parseMap["error_message"].(string); ok {
			content.WriteString(fmt.Sprintf("Error: %s", errorMsg))
		}
		
		return errorBoxStyle.Render(content.String())
	}

	// Handle successful parse case
	var content strings.Builder
	content.WriteString(successLabelStyle.Render("✅ EXPECTED: Parse Success") + "\n")
	content.WriteString(fmt.Sprintf("Count: %.0f assertion(s)\n", count))

	if expectedData, ok := parseMap["expected"].([]interface{}); ok && !compact {
		content.WriteString("\nEntries:\n")
		for i, entryData := range expectedData {
			if entryMap, ok := entryData.(map[string]interface{}); ok {
				key, _ := entryMap["key"].(string)
				value, _ := entryMap["value"].(string)
				content.WriteString(fmt.Sprintf("  %d. key=%q, value=%q\n", i+1, key, value))
			}
		}
	}
	
	return successBoxStyle.Render(content.String())
}

func (m tuiModel) renderMetadata(test Test) string {
	var content strings.Builder
	
	content.WriteString("🏷️  TAGS:\n")
	for _, tag := range test.Meta.Tags {
		content.WriteString(tagStyle.Render(tag))
	}
	content.WriteString("\n\n")
	
	content.WriteString(fmt.Sprintf("📊 LEVEL: %d | FEATURE: %s\n", test.Meta.Level, test.Meta.Feature))
	
	if len(test.Meta.Conflicts) > 0 {
		content.WriteString("\n⚠️  CONFLICTS:\n")
		for _, conflict := range test.Meta.Conflicts {
			content.WriteString(conflictStyle.Render(conflict))
		}
	}
	
	return metaBoxStyle.Render(content.String())
}

func runTUI(filename string) {
	model := initialTUIModel()
	p := tea.NewProgram(model, tea.WithAltScreen())
	
	// Load the test file
	if data, err := os.ReadFile(filename); err == nil {
		var suite TestSuite
		if err := json.Unmarshal(data, &suite); err == nil {
			// Filter to only parse tests
			parseTests := []Test{}
			for _, test := range suite.Tests {
				if _, hasParse := test.Validations["parse"]; hasParse {
					parseTests = append(parseTests, test)
				}
			}
			model.tests = parseTests
			model.suite = suite
			model.filename = filename
		}
	}
	
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running TUI: %v", err)
		os.Exit(1)
	}
}