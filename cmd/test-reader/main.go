package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Configuration constants
const (
	maxEntriesDisplay = 6 // Maximum entries to show before scrolling/truncation
)

// Color palette and styles
var (
	primaryColor = lipgloss.Color("#00D7FF")
	successColor = lipgloss.Color("#00D787")
	errorColor   = lipgloss.Color("#FF5F87")
	warningColor = lipgloss.Color("#FFAF00")
	subtleColor  = lipgloss.Color("#626262")

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
			Margin(1, 0, 0, 0)

	descriptionStyle = lipgloss.NewStyle().
				Foreground(warningColor).
				Italic(true).
				Margin(0, 0, 0, 2)

	// Input CCL simple style
	inputHeaderStyle = lipgloss.NewStyle().
				Foreground(primaryColor).
				Bold(true)

	inputContentStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(lipgloss.Color("#1A1A1A")).
				Padding(0, 1).
				Margin(0, 0, 1, 2)

	// Compact validation styles (no boxes)
	successHeaderStyle = lipgloss.NewStyle().
				Foreground(successColor).
				Bold(true)

	errorHeaderStyle = lipgloss.NewStyle().
				Foreground(errorColor).
				Bold(true)

	successLabelStyle = lipgloss.NewStyle().
				Foreground(successColor).
				Bold(true)

	errorLabelStyle = lipgloss.NewStyle().
			Foreground(errorColor).
			Bold(true)

	// Compact metadata styles
	metaHeaderStyle = lipgloss.NewStyle().
			Foreground(warningColor).
			Bold(true)

	tagStyle = lipgloss.NewStyle().
			Foreground(warningColor).
			Bold(true)

	conflictStyle = lipgloss.NewStyle().
			Foreground(errorColor).
			Bold(true)

	// Entry display styles
	entryKeyStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true)

	entryValueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF"))

	// Equals sign style
	entryEqualsStyle = lipgloss.NewStyle().
				Foreground(warningColor).
				Bold(true)

	// Entry box style - more visible border with tighter spacing
	entryBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(subtleColor).
			Padding(0, 1).
			Margin(0, 0, 0, 2)

	// Whitespace indicator styles
	whitespaceStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#555555"))

	// Empty indicator styles - brighter for better visibility
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
	Tags      []string `json:"tags"`
	Level     int      `json:"level"`
	Feature   string   `json:"feature"`
	Conflicts []string `json:"conflicts,omitempty"`
}

// ParseValidation represents the parse validation structure
type ParseValidation struct {
	Count        int     `json:"count"`
	Expected     []Entry `json:"expected,omitempty"`
	Error        bool    `json:"error,omitempty"`
	ErrorMessage string  `json:"error_message,omitempty"`
}

// Entry represents a key-value pair
type Entry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// visualizeWhitespaceInline shows spaces as dots and tabs as arrows without styling
func visualizeWhitespaceInline(s string) string {
	result := strings.ReplaceAll(s, " ", "·")
	result = strings.ReplaceAll(result, "\t", "→")
	return result
}

// formatKey handles empty keys with special indicator
func formatKey(key string) string {
	if key == "" {
		return emptyKeyStyle.Render("(empty-key)")
	}
	return entryKeyStyle.Render(visualizeWhitespaceInline(key))
}

// formatValue handles empty values with special indicator
func formatValue(value string) string {
	if value == "" {
		return emptyValueStyle.Render("(empty)")
	}
	return entryValueStyle.Render(visualizeWhitespaceInline(value))
}

// formatInputContent shows input with whitespace indicators and proper background
func formatInputContent(input string) string {
	if input == "" {
		return "(empty)"
	}
	return visualizeWhitespaceInline(input)
}

// FileInfo represents a test file with metadata
type FileInfo struct {
	Path        string
	Name        string
	Description string
	TestCount   int
	ParseTests  int
}

func getJSONFiles(dir string) ([]FileInfo, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var jsonFiles []FileInfo
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") && !strings.Contains(file.Name(), "schema") {
			filePath := filepath.Join(dir, file.Name())
			fileInfo := FileInfo{
				Path: filePath,
				Name: file.Name(),
			}

			// Try to read suite info
			if data, err := os.ReadFile(filePath); err == nil {
				var suite TestSuite
				if err := json.Unmarshal(data, &suite); err == nil {
					fileInfo.Description = suite.Description
					fileInfo.TestCount = len(suite.Tests)

					// Count parse tests
					parseCount := 0
					for _, test := range suite.Tests {
						if _, hasParse := test.Validations["parse"]; hasParse {
							parseCount++
						}
					}
					fileInfo.ParseTests = parseCount
				}
			}

			jsonFiles = append(jsonFiles, fileInfo)
		}
	}

	// Sort by name
	sort.Slice(jsonFiles, func(i, j int) bool {
		return jsonFiles[i].Name < jsonFiles[j].Name
	})

	return jsonFiles, nil
}

func runFileSelectionCLI(dir string) {
	files, err := getJSONFiles(dir)
	if err != nil {
		log.Printf("Error reading directory %s: %v", dir, err)
		return
	}

	if len(files) == 0 {
		fmt.Println("No JSON test files found in directory:", dir)
		return
	}

	// Display header
	header := fmt.Sprintf("📁 Available Test Files in %s", dir)
	fmt.Println(fileHeaderStyle.Render(header))
	fmt.Println()

	// Display files
	nameStyle := lipgloss.NewStyle().Foreground(primaryColor).Bold(true)
	infoStyle := lipgloss.NewStyle().Foreground(subtleColor)

	for i, file := range files {
		fmt.Printf("%2d. %s\n", i+1, nameStyle.Render(file.Name))
		if file.Description != "" {
			fmt.Printf("    %s\n", infoStyle.Render(file.Description))
		}
		fmt.Printf("    %s\n", infoStyle.Render(fmt.Sprintf("Total: %d tests, Parse: %d tests", file.TestCount, file.ParseTests)))
		fmt.Println()
	}

	// Interactive selection
	fmt.Print("Select a file number (1-", len(files), ") or 'q' to quit: ")
	var input string
	fmt.Scanln(&input)

	if input == "q" || input == "Q" {
		return
	}

	var selection int
	if _, err := fmt.Sscanf(input, "%d", &selection); err != nil || selection < 1 || selection > len(files) {
		fmt.Println("Invalid selection")
		return
	}

	selectedFile := files[selection-1]
	fmt.Printf("\nStarting TUI for: %s\n", selectedFile.Name)

	// Run TUI for the selected file
	runTUI(selectedFile.Path)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: test-reader <test-file.json|directory> [--static]")
		fmt.Println("       test-reader tests/                              # Interactive TUI (default)")
		fmt.Println("       test-reader tests/api_essential-parsing.json   # Interactive TUI (default)")
		fmt.Println("       test-reader tests/ --static                     # Static CLI output")
		fmt.Println("       test-reader tests/api_essential-parsing.json --static")
		os.Exit(1)
	}

	path := os.Args[1]
	useStatic := len(os.Args) > 2 && os.Args[2] == "--static"

	// Check if path is a directory
	if info, err := os.Stat(path); err == nil && info.IsDir() {
		if useStatic {
			runFileSelectionCLI(path)
		} else {
			runFileSelectionTUI(path) // TUI is default for directories
		}
	} else {
		// Handle as single file
		if useStatic {
			if err := processTestFile(path); err != nil {
				log.Printf("Error processing %s: %v", path, err)
			}
		} else {
			runTUI(path) // TUI is default for files
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

	// Display input compactly
	fmt.Println(inputHeaderStyle.Render("📄 CCL INPUT:"))
	fmt.Println(inputContentStyle.Render(formatInputContent(test.Input)))

	// Display parse validation if present
	if parseData, ok := test.Validations["parse"]; ok {
		displayParseValidation(parseData)
	}

	// Display selective metadata
	displaySelectiveMetadata(test)
	fmt.Println()
}

func displayParseValidation(parseData interface{}) {
	// Convert to map first to handle the interface{}
	parseMap, ok := parseData.(map[string]interface{})
	if !ok {
		fmt.Println(errorHeaderStyle.Render("❌ Invalid parse validation format"))
		return
	}

	count, _ := parseMap["count"].(float64) // JSON numbers are float64

	// Check if this is an error case
	if errorVal, hasError := parseMap["error"]; hasError && errorVal == true {
		fmt.Println(errorHeaderStyle.Render("❌ EXPECTED: Parse Error"))
		fmt.Printf("   Count: %.0f assertion(s)\n", count)

		if errorMsg, ok := parseMap["error_message"].(string); ok {
			fmt.Printf("   Error: %s\n", errorMsg)
		}
		return
	}

	// Handle successful parse case
	fmt.Println(successHeaderStyle.Render("✅ EXPECTED: Parse Success"))
	fmt.Printf("   Count: %.0f assertion(s)\n", count)

	if expectedData, ok := parseMap["expected"].([]interface{}); ok {
		totalEntries := len(expectedData)
		fmt.Printf("   Entries (%d total):\n", totalEntries)

		// Show up to maxEntriesDisplay entries
		entriesToShow := expectedData
		if totalEntries > maxEntriesDisplay {
			entriesToShow = expectedData[:maxEntriesDisplay]
		}

		for _, entryData := range entriesToShow {
			if entryMap, ok := entryData.(map[string]interface{}); ok {
				key, _ := entryMap["key"].(string)
				value, _ := entryMap["value"].(string)

				// Boxed entry content with key/equals on first line, value on second
				keyLine := fmt.Sprintf("%s %s", formatKey(key), entryEqualsStyle.Render("="))
				valueLine := formatValue(value)
				entryContent := fmt.Sprintf("%s\n%s", keyLine, valueLine)
				fmt.Println(entryBoxStyle.Render(entryContent))
			}
		}

		// Show truncation indicator if there are more entries
		if totalEntries > maxEntriesDisplay {
			remaining := totalEntries - maxEntriesDisplay
			truncationMsg := fmt.Sprintf("... and %d more entries (use TUI mode for scrolling)", remaining)
			truncationStyle := lipgloss.NewStyle().Foreground(subtleColor)
			fmt.Println(truncationStyle.Render("   " + truncationMsg))
		}
	}
}

func displaySelectiveMetadata(test Test) {
	// Show variant tags (behavior:* and variant:*)
	variantTags := []string{}
	for _, tag := range test.Meta.Tags {
		if strings.HasPrefix(tag, "variant:") || strings.HasPrefix(tag, "behavior:") {
			variantTags = append(variantTags, tag)
		}
	}

	if len(variantTags) > 0 {
		fmt.Println(metaHeaderStyle.Render("🔄 VARIANTS:"))
		fmt.Print("   ")
		for i, tag := range variantTags {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(tagStyle.Render(tag))
		}
		fmt.Println()
	}

	// Show conflicts if they exist
	if len(test.Meta.Conflicts) > 0 {
		fmt.Println(metaHeaderStyle.Render("⚠️ CONFLICTS:"))
		fmt.Print("   ")
		for i, conflict := range test.Meta.Conflicts {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(conflictStyle.Render(conflict))
		}
		fmt.Println()
	}
}

// File Selection TUI Model
type fileSelectionModel struct {
	files        []FileInfo
	directory    string
	selectedFile int
	width        int
	height       int
	fileSelected bool
}

func initialFileSelectionModel(dir string) fileSelectionModel {
	return fileSelectionModel{
		directory:    dir,
		selectedFile: 0,
		width:        80,
		height:       24,
		fileSelected: false,
	}
}

func (m fileSelectionModel) Init() tea.Cmd {
	return nil
}

func (m fileSelectionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "j", "down":
			if m.selectedFile < len(m.files)-1 {
				m.selectedFile++
			}
		case "k", "up":
			if m.selectedFile > 0 {
				m.selectedFile--
			}
		case "g":
			m.selectedFile = 0
		case "G":
			if len(m.files) > 0 {
				m.selectedFile = len(m.files) - 1
			}
		case "enter", " ":
			if len(m.files) > 0 && m.selectedFile < len(m.files) {
				m.fileSelected = true
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m fileSelectionModel) View() string {
	if len(m.files) == 0 {
		return "Loading files..."
	}

	var content strings.Builder

	// Header
	header := fmt.Sprintf("📁 Select Test File from %s", m.directory)
	content.WriteString(fileHeaderStyle.Render(header) + "\n\n")

	// File list
	var fileList strings.Builder
	for i, file := range m.files {
		prefix := "  "
		style := lipgloss.NewStyle()

		if i == m.selectedFile {
			prefix = "► "
			style = selectedFileStyle
		}

		// File name line
		fileName := fmt.Sprintf("%s%s", prefix, file.Name)
		fileList.WriteString(style.Render(fileName) + "\n")

		// Description line (if selected or always show brief info)
		if i == m.selectedFile && file.Description != "" {
			desc := fmt.Sprintf("   %s", file.Description)
			descStyle := lipgloss.NewStyle().Foreground(subtleColor)
			fileList.WriteString(descStyle.Render(desc) + "\n")
		}

		// Stats line for selected file
		if i == m.selectedFile {
			stats := fmt.Sprintf("   Total: %d tests, Parse: %d tests", file.TestCount, file.ParseTests)
			statsStyle := lipgloss.NewStyle().Foreground(subtleColor)
			fileList.WriteString(statsStyle.Render(stats) + "\n")
		}

		fileList.WriteString("\n")
	}

	content.WriteString(fileListStyle.Render(fileList.String()) + "\n")

	// Navigation help
	help := "j/k: navigate • g/G: first/last • enter/space: select • q/esc: quit"
	content.WriteString(suiteInfoStyle.Render(help))

	return content.String()
}

func runFileSelectionTUI(dir string) {
	files, err := getJSONFiles(dir)
	if err != nil {
		log.Printf("Error reading directory %s: %v", dir, err)
		return
	}

	if len(files) == 0 {
		fmt.Println("No JSON test files found in directory:", dir)
		return
	}

	model := initialFileSelectionModel(dir)
	model.files = files

	p := tea.NewProgram(model, tea.WithAltScreen())
	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Error running file selection TUI: %v", err)
		os.Exit(1)
	}

	// Check if a file was selected
	if fsModel, ok := finalModel.(fileSelectionModel); ok && fsModel.fileSelected && fsModel.selectedFile >= 0 && fsModel.selectedFile < len(fsModel.files) {
		selectedFile := fsModel.files[fsModel.selectedFile]
		// Run TUI with directory context for back navigation
		runTUIWithBackNav(selectedFile.Path, dir)
	}
}

// TUI Implementation
type tuiModel struct {
	tests       []Test
	suite       TestSuite
	filename    string
	directory   string // For back navigation
	currentTest int
	showAll     bool
	width       int
	height      int
	wantsBack   bool // Track if user wants to go back
	entryScroll int  // Current entry scroll offset
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
	if m.filename != "" {
		return loadTestFileCmd(m.filename)
	}
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
		case "esc":
			if m.directory != "" {
				// Set flag and quit to go back to directory selection
				m.wantsBack = true
				return m, tea.Quit
			}
		case "j", "down":
			if m.currentTest < len(m.tests)-1 {
				m.currentTest++
				m.entryScroll = 0 // Reset scroll when changing tests
			}
		case "k", "up":
			if m.currentTest > 0 {
				m.currentTest--
				m.entryScroll = 0 // Reset scroll when changing tests
			}
		case "g":
			m.currentTest = 0
			m.entryScroll = 0
		case "G":
			m.currentTest = len(m.tests) - 1
			m.entryScroll = 0
		case "a":
			m.showAll = !m.showAll
		case "left", "h":
			if m.entryScroll > 0 {
				m.entryScroll--
			}
		case "right", "l":
			// Scroll entries forward if there are more to show
			if m.currentTest < len(m.tests) {
				if parseData, ok := m.tests[m.currentTest].Validations["parse"]; ok {
					if parseMap, ok := parseData.(map[string]interface{}); ok {
						if expectedData, ok := parseMap["expected"].([]interface{}); ok {
							maxScroll := len(expectedData) - maxEntriesDisplay
							if maxScroll > 0 && m.entryScroll < maxScroll {
								m.entryScroll++
							}
						}
					}
				}
			}
		}
	}
	return m, nil
}

func (m tuiModel) View() string {
	if len(m.tests) == 0 {
		return fmt.Sprintf("Loading... (tests=%d, suite=%s, filename=%s)", len(m.tests), m.suite.Suite, m.filename)
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
	help := "j/k: navigate • g/G: first/last • a: toggle all • h/l: scroll entries • q: quit"
	if m.directory != "" {
		help += " • esc: back to file selection"
	}

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
		content.WriteString(descriptionStyle.Render("📝 "+test.Description) + "\n")
	}

	// Input CCL compact
	content.WriteString(inputHeaderStyle.Render("📄 CCL INPUT:") + "\n")
	content.WriteString(inputContentStyle.Render(formatInputContent(test.Input)) + "\n")

	// Parse validation
	if parseData, ok := test.Validations["parse"]; ok {
		content.WriteString(m.renderParseValidation(parseData, compact) + "\n")
	}

	// Selective metadata (only if not compact)
	if !compact {
		content.WriteString(m.renderSelectiveMetadata(test) + "\n")
	}

	return content.String()
}

func (m tuiModel) renderParseValidation(parseData interface{}, compact bool) string {
	parseMap, ok := parseData.(map[string]interface{})
	if !ok {
		return errorHeaderStyle.Render("❌ Invalid parse validation format")
	}

	count, _ := parseMap["count"].(float64)

	// Check if this is an error case
	if errorVal, hasError := parseMap["error"]; hasError && errorVal == true {
		var content strings.Builder
		content.WriteString(errorHeaderStyle.Render("❌ EXPECTED: Parse Error") + "\n")
		content.WriteString(fmt.Sprintf("   Count: %.0f assertion(s)\n", count))

		if errorMsg, ok := parseMap["error_message"].(string); ok {
			content.WriteString(fmt.Sprintf("   Error: %s\n", errorMsg))
		}
		return content.String()
	}

	// Handle successful parse case
	var content strings.Builder
	content.WriteString(successHeaderStyle.Render("✅ EXPECTED: Parse Success") + "\n")
	content.WriteString(fmt.Sprintf("   Count: %.0f assertion(s)\n", count))

	if expectedData, ok := parseMap["expected"].([]interface{}); ok && !compact {
		totalEntries := len(expectedData)
		content.WriteString(fmt.Sprintf("   Entries (%d total):\n", totalEntries))

		// Handle scrolling
		startIdx := m.entryScroll
		if startIdx >= totalEntries {
			startIdx = totalEntries - 1
			if startIdx < 0 {
				startIdx = 0
			}
		}

		endIdx := startIdx + maxEntriesDisplay
		if endIdx > totalEntries {
			endIdx = totalEntries
		}

		// Show scroll indicators if needed
		if startIdx > 0 {
			scrollStyle := lipgloss.NewStyle().Foreground(subtleColor)
			content.WriteString(scrollStyle.Render("   ↑ More entries above (h/← to scroll up)\n"))
		}

		// Show entries in current scroll window
		for i := startIdx; i < endIdx; i++ {
			entryData := expectedData[i]
			if entryMap, ok := entryData.(map[string]interface{}); ok {
				key, _ := entryMap["key"].(string)
				value, _ := entryMap["value"].(string)

				// Boxed entry content with key/equals on first line, value on second
				keyLine := fmt.Sprintf("%s %s", formatKey(key), entryEqualsStyle.Render("="))
				valueLine := formatValue(value)
				entryContent := fmt.Sprintf("%s\n%s", keyLine, valueLine)
				content.WriteString(entryBoxStyle.Render(entryContent) + "\n")
			}
		}

		// Show scroll indicator if there are more entries below
		if endIdx < totalEntries {
			remaining := totalEntries - endIdx
			scrollStyle := lipgloss.NewStyle().Foreground(subtleColor)
			content.WriteString(scrollStyle.Render(fmt.Sprintf("   ↓ %d more entries below (l/→ to scroll down)\n", remaining)))
		}
	}

	return content.String()
}

func (m tuiModel) renderSelectiveMetadata(test Test) string {
	var content strings.Builder

	// Show variant tags (behavior:* and variant:*)
	variantTags := []string{}
	for _, tag := range test.Meta.Tags {
		if strings.HasPrefix(tag, "variant:") || strings.HasPrefix(tag, "behavior:") {
			variantTags = append(variantTags, tag)
		}
	}

	if len(variantTags) > 0 {
		content.WriteString(metaHeaderStyle.Render("🔄 VARIANTS:") + "\n")
		content.WriteString("   ")
		for i, tag := range variantTags {
			if i > 0 {
				content.WriteString(", ")
			}
			content.WriteString(tagStyle.Render(tag))
		}
		content.WriteString("\n")
	}

	// Show conflicts if they exist
	if len(test.Meta.Conflicts) > 0 {
		content.WriteString(metaHeaderStyle.Render("⚠️ CONFLICTS:") + "\n")
		content.WriteString("   ")
		for i, conflict := range test.Meta.Conflicts {
			if i > 0 {
				content.WriteString(", ")
			}
			content.WriteString(conflictStyle.Render(conflict))
		}
		content.WriteString("\n")
	}

	return content.String()
}

func runTUI(filename string) {
	model := initialTUIModel()
	model.filename = filename

	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running TUI: %v", err)
		os.Exit(1)
	}
}

func runTUIWithBackNav(filename, directory string) {
	for {
		model := initialTUIModel()
		model.filename = filename
		model.directory = directory

		p := tea.NewProgram(model, tea.WithAltScreen())
		finalModel, err := p.Run()
		if err != nil {
			fmt.Printf("Error running TUI: %v", err)
			os.Exit(1)
		}

		// Check if user pressed escape to go back
		if tuiModel, ok := finalModel.(tuiModel); ok && tuiModel.wantsBack {
			// Go back to directory selection
			runFileSelectionTUI(directory)
			break
		} else {
			// User quit normally, exit completely
			break
		}
	}
}
