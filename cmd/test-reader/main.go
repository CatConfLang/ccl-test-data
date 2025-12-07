// CCL Test Reader - Interactive viewer for CCL test files
//
// Designed for generated format files (generated_tests/) with split-view layout.
// Source format files (source_tests/) have limited display support.
//
// The test-reader displays CCL test cases with their input, expected output,
// and metadata. Features a split-view TUI with filterable test list and detail pane.
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
	"github.com/tylerbu/ccl-test-lib/config"
	"github.com/tylerbu/ccl-test-lib/loader"
	"github.com/tylerbu/ccl-test-lib/types"
)

// Configuration constants
const (
	maxEntriesDisplay = 6  // Maximum entries to show before scrolling/truncation
	listPaneWidth     = 40 // Width of the test list pane in split view
	minDetailWidth    = 60 // Minimum width for detail pane
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

	// Split view styles
	listPaneStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(subtleColor).
			Padding(0, 1)

	detailPaneStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(primaryColor).
			Padding(0, 1)

	filterInputStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(warningColor).
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
)

// Use types from ccl-test-lib - these are aliases for convenience
type TestSuite = types.TestSuite
type TestCase = types.TestCase
type Entry = types.Entry

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

					// Count parse tests (parse and parse_indented)
					parseCount := 0
					for _, test := range suite.Tests {
						// Check if test has parse or parse_indented validation
						hasParse := false
						if test.Validations != nil {
							if test.Validations.Parse != nil || test.Validations.ParseIndented != nil {
								hasParse = true
							}
						} else if test.Validation == "parse" || test.Validation == "parse_indented" {
							hasParse = true
						}
						if hasParse {
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
		fmt.Printf("    %s\n", infoStyle.Render(fmt.Sprintf("Total: %d tests, Parse/ParseIndented: %d tests", file.TestCount, file.ParseTests)))
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
	// TODO: Remove support for source format files completely.
	// The test-reader should only work with generated flat format files.
	// Source format files should be converted to generated format first.

	// Detect if this is a source format file and emit warning
	isSourceFormat := strings.Contains(filename, "source_tests")
	if isSourceFormat {
		warningStyle := lipgloss.NewStyle().Foreground(warningColor).Bold(true)
		fmt.Println(warningStyle.Render("⚠️  WARNING: Source format files have limited display support."))
		fmt.Println(warningStyle.Render("   Use generated format files for full entry details."))
		fmt.Println(warningStyle.Render("   Run 'just generate' to convert source tests to generated format."))
		fmt.Println()
	}

	// Use ccl-test-lib loader to handle both source and flat formats
	impl := config.ImplementationConfig{
		Name:               "test-reader",
		Version:            "1.0.0",
		SupportedFunctions: []config.CCLFunction{config.FunctionParse, config.FunctionParseIndented}, // Support parse and parse_indented tests for display
	}
	testLoader := loader.NewTestLoader(".", impl)
	suite, err := testLoader.LoadTestFile(filename, loader.LoadOptions{
		Format:     loader.FormatFlat,
		FilterMode: loader.FilterAll, // Load all tests for viewer
	})
	if err != nil {
		return fmt.Errorf("loading test suite: %w", err)
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
		summary := "📋 No parse tests (parse/parse_indented) found in this file"
		fmt.Println(summaryStyle.Render(summary))
	} else {
		summary := fmt.Sprintf("📊 Found %d parse test(s) (parse/parse_indented)", parseOnlyCount)
		fmt.Println(summaryStyle.Render(summary))
	}
	fmt.Println()

	return nil
}

func hasParseValidation(test TestCase) bool {
	// ccl-test-lib TestCase should have parse functionality built-in
	// For now, assume all loaded tests are valid parse tests
	return true
}

func displayTest(test TestCase, index int) {
	// Test header
	header := fmt.Sprintf("Test #%d: %s", index, test.Name)
	fmt.Println(testHeaderStyle.Render(header))

	// Note: ccl-test-lib TestCase doesn't have Description field
	// Description would be in suite metadata if needed

	// Display input compactly (use first input for single-input tests)
	fmt.Println(inputHeaderStyle.Render("📄 CCL INPUT:"))
	inputText := ""
	if len(test.Inputs) > 0 {
		inputText = test.Inputs[0]
	}
	fmt.Println(inputContentStyle.Render(formatInputContent(inputText)))

	// Display parse validation - use the Expected field from ccl-test-lib
	displayParseValidationFromTestCase(test)

	// Display selective metadata
	displaySelectiveMetadata(test)
	fmt.Println()
}

func displayParseValidationFromTestCase(test TestCase) {
	// Handle successful parse case using ccl-test-lib TestCase
	fmt.Println(successHeaderStyle.Render("✅ EXPECTED: Parse Success"))

	// Handle Expected field which can be array of entries or map with entries
	var entries []Entry
	var count int

	// Try to extract entries from Expected interface{}
	if test.Expected != nil {
		switch expected := test.Expected.(type) {
		case []interface{}:
			// Direct array of entries
			for _, item := range expected {
				if entryMap, ok := item.(map[string]interface{}); ok {
					key, _ := entryMap["key"].(string)
					value, _ := entryMap["value"].(string)
					entries = append(entries, Entry{Key: key, Value: value})
				}
			}
			count = len(entries)
		case map[string]interface{}:
			// Map with count and entries fields
			if c, ok := expected["count"]; ok {
				if countFloat, ok := c.(float64); ok {
					count = int(countFloat)
				}
			}
			if entriesArray, ok := expected["entries"].([]interface{}); ok {
				for _, item := range entriesArray {
					if entryMap, ok := item.(map[string]interface{}); ok {
						key, _ := entryMap["key"].(string)
						value, _ := entryMap["value"].(string)
						entries = append(entries, Entry{Key: key, Value: value})
					}
				}
			}
		}
	}

	fmt.Printf("   Count: %d assertion(s)\n", count)

	if len(entries) > 0 {
		totalEntries := len(entries)
		fmt.Printf("   Entries (%d total):\n", totalEntries)

		// Show up to maxEntriesDisplay entries
		entriesToShow := entries
		if totalEntries > maxEntriesDisplay {
			entriesToShow = entries[:maxEntriesDisplay]
		}

		for _, entry := range entriesToShow {
			// Display key=value on same line unless value contains newlines
			if strings.Contains(entry.Value, "\n") {
				// Multiline value: key on first line, value on subsequent lines
				keyLine := fmt.Sprintf("%s %s", formatKey(entry.Key), entryEqualsStyle.Render("="))
				valueLine := formatValue(entry.Value)
				entryContent := fmt.Sprintf("%s\n%s", keyLine, valueLine)
				fmt.Println(entryBoxStyle.Render(entryContent))
			} else {
				// Single line: key = value
				entryContent := fmt.Sprintf("%s %s %s", formatKey(entry.Key), entryEqualsStyle.Render("="), formatValue(entry.Value))
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

func displaySelectiveMetadata(test TestCase) {
	// Show behavior tags if available in ccl-test-lib TestCase
	variantTags := []string{}

	// Add behaviors to variant tags
	for _, behavior := range test.Behaviors {
		variantTags = append(variantTags, "behavior:"+behavior)
	}

	// Add any other relevant metadata from TestCase
	if len(test.Features) > 0 {
		for _, feature := range test.Features {
			variantTags = append(variantTags, "feature:"+feature)
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

	// Note: ccl-test-lib TestCase might not have Conflicts field
	// This functionality may need to be adapted based on the actual structure
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
			stats := fmt.Sprintf("   Total: %d tests, Parse/ParseValue: %d tests", file.TestCount, file.ParseTests)
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

// FilterMode represents the active filter type
type FilterMode int

const (
	FilterNone FilterMode = iota
	FilterName
	FilterFunction
	FilterBehavior
	FilterFeature
)

// TUI Implementation - Split View with Filter
type tuiModel struct {
	// Test data
	tests         []TestCase
	filteredTests []TestCase // Tests matching current filter
	suite         TestSuite
	filename      string
	directory     string // For back navigation

	// Selection state
	currentTest int
	listScroll  int // Scroll offset for test list
	entryScroll int // Scroll offset for entries in detail view

	// Filter state
	filterMode   FilterMode
	filterText   string
	filterActive bool // True when typing in filter

	// Layout
	width     int
	height    int
	focusPane int // 0=list, 1=detail

	// Navigation
	wantsBack bool
}

type testLoadedMsg struct {
	suite    TestSuite
	filename string
}

func initialTUIModel() tuiModel {
	return tuiModel{
		currentTest:   0,
		width:         120,
		height:        40,
		focusPane:     0,
		filterMode:    FilterNone,
		filteredTests: nil,
	}
}

// applyFilter updates filteredTests based on current filter settings
func (m *tuiModel) applyFilter() {
	if m.filterText == "" || m.filterMode == FilterNone {
		m.filteredTests = m.tests
		return
	}

	filterLower := strings.ToLower(m.filterText)
	var filtered []TestCase

	for _, test := range m.tests {
		match := false
		switch m.filterMode {
		case FilterName:
			match = strings.Contains(strings.ToLower(test.Name), filterLower)
		case FilterFunction:
			if test.Validation != "" {
				match = strings.Contains(strings.ToLower(test.Validation), filterLower)
			}
			for _, fn := range test.Functions {
				if strings.Contains(strings.ToLower(fn), filterLower) {
					match = true
					break
				}
			}
		case FilterBehavior:
			for _, b := range test.Behaviors {
				if strings.Contains(strings.ToLower(b), filterLower) {
					match = true
					break
				}
			}
		case FilterFeature:
			for _, f := range test.Features {
				if strings.Contains(strings.ToLower(f), filterLower) {
					match = true
					break
				}
			}
		}
		if match {
			filtered = append(filtered, test)
		}
	}

	m.filteredTests = filtered
	// Reset selection if current is out of bounds
	if m.currentTest >= len(m.filteredTests) {
		m.currentTest = 0
	}
}

// getValidationBadge returns a styled badge for the validation type
func getValidationBadge(validation string) string {
	switch validation {
	case "parse":
		return parseValidationStyle.Render("parse")
	case "parse_indented":
		return parseIndentedValidationStyle.Render("parse_indented")
	default:
		return otherValidationStyle.Render(validation)
	}
}

func loadTestFileCmd(filename string) tea.Cmd {
	return func() tea.Msg {
		// Use ccl-test-lib loader for generated flat format files
		impl := config.ImplementationConfig{
			Name:               "test-reader",
			Version:            "1.0.0",
			SupportedFunctions: []config.CCLFunction{config.FunctionParse, config.FunctionParseIndented},
		}
		testLoader := loader.NewTestLoader(".", impl)
		suite, err := testLoader.LoadTestFile(filename, loader.LoadOptions{
			Format:     loader.FormatFlat,
			FilterMode: loader.FilterAll,
		})
		if err != nil {
			return nil
		}

		return testLoadedMsg{
			suite:    *suite,
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
		m.tests = msg.suite.Tests
		m.filteredTests = m.tests // Initialize filtered to all tests
		return m, nil

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		// Handle filter input mode
		if m.filterActive {
			switch msg.String() {
			case "enter":
				m.filterActive = false
				m.applyFilter()
			case "esc":
				m.filterActive = false
				m.filterText = ""
				m.filterMode = FilterNone
				m.applyFilter()
			case "backspace":
				if len(m.filterText) > 0 {
					m.filterText = m.filterText[:len(m.filterText)-1]
					m.applyFilter()
				}
			default:
				// Accept printable characters
				if len(msg.String()) == 1 {
					m.filterText += msg.String()
					m.applyFilter()
				}
			}
			return m, nil
		}

		// Normal navigation mode
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			if m.filterMode != FilterNone {
				// Clear filter first
				m.filterText = ""
				m.filterMode = FilterNone
				m.applyFilter()
			} else if m.directory != "" {
				m.wantsBack = true
				return m, tea.Quit
			}
		case "j", "down":
			if m.currentTest < len(m.filteredTests)-1 {
				m.currentTest++
				m.entryScroll = 0
				// Adjust list scroll to keep selection visible
				visibleItems := m.height - 12 // Account for header/footer
				if m.currentTest >= m.listScroll+visibleItems {
					m.listScroll = m.currentTest - visibleItems + 1
				}
			}
		case "k", "up":
			if m.currentTest > 0 {
				m.currentTest--
				m.entryScroll = 0
				if m.currentTest < m.listScroll {
					m.listScroll = m.currentTest
				}
			}
		case "g":
			m.currentTest = 0
			m.listScroll = 0
			m.entryScroll = 0
		case "G":
			if len(m.filteredTests) > 0 {
				m.currentTest = len(m.filteredTests) - 1
				m.entryScroll = 0
				visibleItems := m.height - 12
				if m.currentTest >= visibleItems {
					m.listScroll = m.currentTest - visibleItems + 1
				}
			}
		case "tab":
			// Toggle focus between panes
			m.focusPane = (m.focusPane + 1) % 2
		case "h", "left":
			if m.focusPane == 1 && m.entryScroll > 0 {
				m.entryScroll--
			}
		case "l", "right":
			if m.focusPane == 1 {
				m.scrollEntriesDown()
			}
		// Filter shortcuts
		case "/":
			m.filterMode = FilterName
			m.filterActive = true
			m.filterText = ""
		case "f":
			m.filterMode = FilterFunction
			m.filterActive = true
			m.filterText = ""
		case "b":
			m.filterMode = FilterBehavior
			m.filterActive = true
			m.filterText = ""
		case "F":
			m.filterMode = FilterFeature
			m.filterActive = true
			m.filterText = ""
		case "c":
			// Clear filter
			m.filterText = ""
			m.filterMode = FilterNone
			m.applyFilter()
		}
	}
	return m, nil
}

func (m *tuiModel) scrollEntriesDown() {
	if m.currentTest >= len(m.filteredTests) {
		return
	}
	currentTest := m.filteredTests[m.currentTest]
	entryCount := 0
	if currentTest.Expected != nil {
		if expectedMap, ok := currentTest.Expected.(map[string]interface{}); ok {
			if entriesArray, ok := expectedMap["entries"].([]interface{}); ok {
				entryCount = len(entriesArray)
			}
		}
	}
	maxScroll := entryCount - maxEntriesDisplay
	if maxScroll > 0 && m.entryScroll < maxScroll {
		m.entryScroll++
	}
}

func (m tuiModel) View() string {
	if len(m.tests) == 0 {
		return fmt.Sprintf("Loading... (tests=%d, suite=%s, filename=%s)", len(m.tests), m.suite.Suite, m.filename)
	}

	var content strings.Builder

	// Compact header
	header := fmt.Sprintf("📋 %s", m.suite.Suite)
	headerStyle := lipgloss.NewStyle().Foreground(primaryColor).Bold(true)
	content.WriteString(headerStyle.Render(header) + "\n")

	// Filter bar
	content.WriteString(m.renderFilterBar() + "\n")

	// Calculate pane dimensions
	availableWidth := m.width - 4 // Account for borders
	listWidth := listPaneWidth
	if listWidth > availableWidth/2 {
		listWidth = availableWidth / 2
	}
	detailWidth := availableWidth - listWidth - 3 // 3 for gap between panes

	availableHeight := m.height - 6 // Account for header, filter, and help

	// Render split view panes
	listPane := m.renderListPane(listWidth, availableHeight)
	detailPane := m.renderDetailPane(detailWidth, availableHeight)

	// Join panes horizontally
	splitView := lipgloss.JoinHorizontal(lipgloss.Top, listPane, " ", detailPane)
	content.WriteString(splitView + "\n")

	// Help bar
	content.WriteString(m.renderHelpBar())

	return content.String()
}

func (m tuiModel) renderFilterBar() string {
	var filterInfo string

	if m.filterActive {
		// Show active filter input
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
		filterInfo = filterLabelStyle.Render(fmt.Sprintf("🔍 Filter by %s: ", modeLabel)) +
			activeFilterStyle.Render(m.filterText+"_")
	} else if m.filterMode != FilterNone && m.filterText != "" {
		// Show active filter summary
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
		filterInfo = filterLabelStyle.Render(fmt.Sprintf("🔍 Filtering by %s: ", modeLabel)) +
			activeFilterStyle.Render(m.filterText) +
			inactiveFilterStyle.Render(fmt.Sprintf(" (%d/%d tests)", len(m.filteredTests), len(m.tests)))
	} else {
		// Show filter hints
		filterInfo = inactiveFilterStyle.Render("/ name • f function • b behavior • F feature • c clear")
	}

	return filterInfo
}

func (m tuiModel) renderListPane(width, height int) string {
	var content strings.Builder

	// Calculate visible items
	visibleItems := height - 2 // Account for border
	if visibleItems < 1 {
		visibleItems = 1
	}

	// Determine which tests to show
	startIdx := m.listScroll
	endIdx := startIdx + visibleItems
	if endIdx > len(m.filteredTests) {
		endIdx = len(m.filteredTests)
	}

	// Show scroll indicator at top if needed
	if startIdx > 0 {
		scrollUp := lipgloss.NewStyle().Foreground(subtleColor).Render("↑ more above")
		content.WriteString(scrollUp + "\n")
		visibleItems--
		endIdx = startIdx + visibleItems
		if endIdx > len(m.filteredTests) {
			endIdx = len(m.filteredTests)
		}
	}

	// Render test list items
	for i := startIdx; i < endIdx; i++ {
		test := m.filteredTests[i]
		isSelected := i == m.currentTest

		// Build compact test summary
		prefix := "  "
		if isSelected {
			prefix = "► "
		}

		// Validation badge (short form)
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

		// Truncate name to fit
		maxNameLen := width - 6 // Account for prefix and badge
		name := test.Name
		if len(name) > maxNameLen {
			name = name[:maxNameLen-2] + ".."
		}

		line := fmt.Sprintf("%s[%s] %s", prefix, badge, name)

		if isSelected {
			style := selectedFileStyle.Width(width - 2)
			content.WriteString(style.Render(line) + "\n")
		} else {
			content.WriteString(line + "\n")
		}
	}

	// Show scroll indicator at bottom if needed
	if endIdx < len(m.filteredTests) {
		remaining := len(m.filteredTests) - endIdx
		scrollDown := lipgloss.NewStyle().Foreground(subtleColor).Render(fmt.Sprintf("↓ %d more below", remaining))
		content.WriteString(scrollDown)
	}

	// Apply list pane style
	paneStyle := listPaneStyle.Width(width).Height(height)
	if m.focusPane == 0 {
		paneStyle = paneStyle.BorderForeground(primaryColor)
	}

	return paneStyle.Render(content.String())
}

func (m tuiModel) renderDetailPane(width, height int) string {
	if len(m.filteredTests) == 0 {
		emptyMsg := "No tests match the current filter"
		paneStyle := detailPaneStyle.Width(width).Height(height)
		return paneStyle.Render(emptyMsg)
	}

	if m.currentTest >= len(m.filteredTests) {
		return detailPaneStyle.Width(width).Height(height).Render("No test selected")
	}

	test := m.filteredTests[m.currentTest]
	var content strings.Builder

	// Test header with validation badge
	testHeader := fmt.Sprintf("%s  %s", test.Name, getValidationBadge(test.Validation))
	content.WriteString(testHeaderStyle.Render(testHeader) + "\n\n")

	// Source test reference if available
	if test.SourceTest != "" && test.SourceTest != test.Name {
		sourceRef := lipgloss.NewStyle().Foreground(subtleColor).Render(fmt.Sprintf("from: %s", test.SourceTest))
		content.WriteString(sourceRef + "\n\n")
	}

	// Input section
	content.WriteString(inputHeaderStyle.Render("📄 INPUT:") + "\n")
	inputText := ""
	if len(test.Inputs) > 0 {
		inputText = test.Inputs[0]
	}
	// Constrain input width
	inputStyle := inputContentStyle.Width(width - 6)
	content.WriteString(inputStyle.Render(formatInputContent(inputText)) + "\n\n")

	// Expected output section
	content.WriteString(m.renderExpectedOutput(test, width-4) + "\n")

	// Metadata section
	content.WriteString(m.renderMetadata(test))

	// Apply detail pane style
	paneStyle := detailPaneStyle.Width(width).Height(height)
	if m.focusPane == 1 {
		paneStyle = paneStyle.BorderForeground(primaryColor)
	}

	return paneStyle.Render(content.String())
}

func (m tuiModel) renderExpectedOutput(test TestCase, width int) string {
	var content strings.Builder

	// Determine success/error status
	isError := test.ExpectError
	if isError {
		content.WriteString(errorHeaderStyle.Render("❌ EXPECTED: Error") + "\n")
	} else {
		header := "✅ EXPECTED"
		if test.Validation == "parse" || test.Validation == "parse_indented" {
			header = "✅ EXPECTED: Entries"
		}
		content.WriteString(successHeaderStyle.Render(header) + "\n")
	}

	// Extract entries from Expected
	// Note: ccl-test-lib loader transforms expected to just the entries array
	// for parse/parse_indented validation types
	var entries []Entry

	if test.Expected != nil {
		switch expected := test.Expected.(type) {
		case []interface{}:
			// Direct array of entries (from loader transformation)
			for _, item := range expected {
				if entryMap, ok := item.(map[string]interface{}); ok {
					key, _ := entryMap["key"].(string)
					value, _ := entryMap["value"].(string)
					entries = append(entries, Entry{Key: key, Value: value})
				}
			}
		case map[string]interface{}:
			// Structured format with count/entries (raw JSON)
			if entriesArray, ok := expected["entries"].([]interface{}); ok {
				for _, item := range entriesArray {
					if entryMap, ok := item.(map[string]interface{}); ok {
						key, _ := entryMap["key"].(string)
						value, _ := entryMap["value"].(string)
						entries = append(entries, Entry{Key: key, Value: value})
					}
				}
			}
		}
	}

	count := len(entries)
	content.WriteString(fmt.Sprintf("Count: %d\n", count))

	if len(entries) > 0 {
		totalEntries := len(entries)

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

		// Scroll indicator up
		if startIdx > 0 {
			scrollStyle := lipgloss.NewStyle().Foreground(subtleColor)
			content.WriteString(scrollStyle.Render("↑ more above\n"))
		}

		// Show entries
		for i := startIdx; i < endIdx; i++ {
			entry := entries[i]
			if strings.Contains(entry.Value, "\n") {
				keyLine := fmt.Sprintf("%s %s", formatKey(entry.Key), entryEqualsStyle.Render("="))
				valueLine := formatValue(entry.Value)
				entryContent := fmt.Sprintf("%s\n%s", keyLine, valueLine)
				content.WriteString(entryBoxStyle.Render(entryContent) + "\n")
			} else {
				entryContent := fmt.Sprintf("%s %s %s", formatKey(entry.Key), entryEqualsStyle.Render("="), formatValue(entry.Value))
				content.WriteString(entryBoxStyle.Render(entryContent) + "\n")
			}
		}

		// Scroll indicator down
		if endIdx < totalEntries {
			remaining := totalEntries - endIdx
			scrollStyle := lipgloss.NewStyle().Foreground(subtleColor)
			content.WriteString(scrollStyle.Render(fmt.Sprintf("↓ %d more (tab→detail, h/l scroll)\n", remaining)))
		}
	}

	return content.String()
}

func (m tuiModel) renderMetadata(test TestCase) string {
	var content strings.Builder

	// Behaviors
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

	// Features
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

	// Variants
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

	// Navigation
	parts = append(parts, "j/k:navigate")
	parts = append(parts, "g/G:first/last")
	parts = append(parts, "tab:switch pane")

	// Scrolling (when in detail pane)
	if m.focusPane == 1 {
		parts = append(parts, "h/l:scroll entries")
	}

	// Filter status
	if m.filterMode != FilterNone {
		parts = append(parts, "c:clear filter")
	}

	// Exit
	if m.directory != "" {
		parts = append(parts, "esc:back")
	}
	parts = append(parts, "q:quit")

	// Test count
	countInfo := fmt.Sprintf("[%d/%d]", m.currentTest+1, len(m.filteredTests))
	if len(m.filteredTests) != len(m.tests) {
		countInfo = fmt.Sprintf("[%d/%d of %d]", m.currentTest+1, len(m.filteredTests), len(m.tests))
	}

	help := strings.Join(parts, " • ") + " " + summaryStyle.Render(countInfo)
	return suiteInfoStyle.Render(help)
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
