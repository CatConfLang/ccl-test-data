package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/catconflang/ccl-test-data/config"
	"github.com/catconflang/ccl-test-data/loader"
	"github.com/catconflang/ccl-test-data/types"
	tea "github.com/charmbracelet/bubbletea"
)

// Type aliases for convenience
type TestSuite = types.TestSuite
type TestCase = types.TestCase
type Entry = types.Entry

// readerConfig is the shared implementation config for loading tests
var readerConfig = config.ImplementationConfig{
	Name:               "test-reader",
	Version:            "1.0.0",
	SupportedFunctions: []config.CCLFunction{config.FunctionParse, config.FunctionParseIndented},
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

// TUI message types

type testLoadedMsg struct {
	suite    TestSuite
	filename string
}

type testLoadErrorMsg struct {
	err      error
	filename string
}

// tuiModel is the main split-view test viewer
type tuiModel struct {
	// Test data
	tests         []TestCase
	filteredTests []TestCase
	suite         TestSuite
	filename      string
	directory     string // For back navigation

	// Selection state
	currentTest int
	listScroll  int
	entryScroll int

	// Filter state
	filterMode   FilterMode
	filterText   string
	filterActive bool

	// Layout
	width     int
	height    int
	focusPane int // 0=list, 1=detail

	// Navigation
	wantsBack bool

	// Error state
	loadError error
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

// applyFilter updates filteredTests based on current filter settings.
// Called from within Update() since tuiModel uses value receivers for tea.Model.
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
	if m.currentTest >= len(m.filteredTests) {
		m.currentTest = 0
	}
}

// scrollEntriesDown scrolls the detail pane entries down by one.
// Called from within Update() since tuiModel uses value receivers for tea.Model.
func (m *tuiModel) scrollEntriesDown() {
	if m.currentTest >= len(m.filteredTests) {
		return
	}
	maxScroll := countExpectedItems(m.filteredTests[m.currentTest]) - maxEntriesDisplay
	if maxScroll > 0 && m.entryScroll < maxScroll {
		m.entryScroll++
	}
}

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

// Tea commands for async file loading

func loadTestFileCmd(filename string) tea.Cmd {
	return func() tea.Msg {
		testLoader := loader.NewTestLoader(".", readerConfig)
		suite, err := testLoader.LoadTestFile(filename, loader.LoadOptions{
			Format:     loader.FormatFlat,
			FilterMode: loader.FilterAll,
		})
		if err != nil {
			return testLoadErrorMsg{err: err, filename: filename}
		}
		return testLoadedMsg{suite: *suite, filename: filename}
	}
}

func loadAllTestFilesCmd(dir string) tea.Cmd {
	return func() tea.Msg {
		files, err := os.ReadDir(dir)
		if err != nil {
			return testLoadErrorMsg{err: err, filename: dir}
		}

		testLoader := loader.NewTestLoader(".", readerConfig)

		var allTests []TestCase
		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") && !strings.Contains(file.Name(), "schema") {
				filePath := filepath.Join(dir, file.Name())
				suite, err := testLoader.LoadTestFile(filePath, loader.LoadOptions{
					Format:     loader.FormatFlat,
					FilterMode: loader.FilterAll,
				})
				if err != nil {
					continue
				}
				for i := range suite.Tests {
					suite.Tests[i].Name = fmt.Sprintf("[%s] %s", file.Name(), suite.Tests[i].Name)
				}
				allTests = append(allTests, suite.Tests...)
			}
		}

		return testLoadedMsg{
			suite:    TestSuite{Description: "All tests combined", Tests: allTests},
			filename: "All Tests",
		}
	}
}

// Bubbletea interface

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
		m.filteredTests = m.tests
		return m, nil

	case testLoadErrorMsg:
		m.loadError = msg.err
		m.filename = msg.filename
		return m, nil

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
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
					runes := []rune(m.filterText)
					m.filterText = string(runes[:len(runes)-1])
					m.applyFilter()
				}
			default:
				if len([]rune(msg.String())) == 1 {
					m.filterText += msg.String()
					m.applyFilter()
				}
			}
			return m, nil
		}

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			if m.filterMode != FilterNone {
				m.filterText = ""
				m.filterMode = FilterNone
				m.applyFilter()
			} else if m.directory != "" {
				m.wantsBack = true
				return m, tea.Quit
			}
		case "j", "down":
			if m.focusPane == 0 {
				if m.currentTest < len(m.filteredTests)-1 {
					m.currentTest++
					m.entryScroll = 0
					visibleItems := m.height - 12
					if m.currentTest >= m.listScroll+visibleItems {
						m.listScroll = m.currentTest - visibleItems + 1
					}
				}
			} else {
				m.scrollEntriesDown()
			}
		case "k", "up":
			if m.focusPane == 0 {
				if m.currentTest > 0 {
					m.currentTest--
					m.entryScroll = 0
					if m.currentTest < m.listScroll {
						m.listScroll = m.currentTest
					}
				}
			} else {
				if m.entryScroll > 0 {
					m.entryScroll--
				}
			}
		case "g":
			if m.focusPane == 0 {
				m.currentTest = 0
				m.listScroll = 0
			}
			m.entryScroll = 0
		case "G":
			if m.focusPane == 0 && len(m.filteredTests) > 0 {
				m.currentTest = len(m.filteredTests) - 1
				m.entryScroll = 0
				visibleItems := m.height - 12
				if m.currentTest >= visibleItems {
					m.listScroll = m.currentTest - visibleItems + 1
				}
			}
		case "tab":
			m.focusPane = (m.focusPane + 1) % 2
		case "h", "left":
			if m.entryScroll > 0 {
				m.entryScroll--
			}
		case "l", "right":
			m.scrollEntriesDown()
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
			m.filterText = ""
			m.filterMode = FilterNone
			m.applyFilter()
		}
	}
	return m, nil
}

// runTUI launches the test viewer for a single file
func runTUI(filename string) {
	model := initialTUIModel()
	model.filename = filename

	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running TUI: %v", err)
		os.Exit(1)
	}
}

// allTestsModel wraps tuiModel to load all tests from a directory on init
type allTestsModel struct {
	tuiModel tuiModel
	dir      string
}

func (m *allTestsModel) Init() tea.Cmd {
	return loadAllTestFilesCmd(m.dir)
}

func (m *allTestsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	newModel, cmd := m.tuiModel.Update(msg)
	if tm, ok := newModel.(tuiModel); ok {
		m.tuiModel = tm
	}
	return m, cmd
}

func (m *allTestsModel) View() string {
	return m.tuiModel.View()
}
