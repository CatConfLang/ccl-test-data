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

// FileInfo represents a test file with metadata
type FileInfo struct {
	Path        string
	Name        string
	Description string
	TestCount   int
	ParseTests  int
	IsVirtual   bool // True for the "All" entry that combines all files
}

func getJSONFiles(dir string) ([]FileInfo, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var jsonFiles []FileInfo
	totalTestCount := 0
	totalParseTests := 0

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") && !strings.Contains(file.Name(), "schema") {
			filePath := filepath.Join(dir, file.Name())
			fileInfo := FileInfo{
				Path: filePath,
				Name: file.Name(),
			}

			if data, err := os.ReadFile(filePath); err == nil {
				var suite TestSuite
				if err := json.Unmarshal(data, &suite); err == nil {
					fileInfo.Description = suite.Description
					fileInfo.TestCount = len(suite.Tests)
					totalTestCount += len(suite.Tests)

					parseCount := 0
					for _, test := range suite.Tests {
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
					totalParseTests += parseCount
				}
			}

			jsonFiles = append(jsonFiles, fileInfo)
		}
	}

	sort.Slice(jsonFiles, func(i, j int) bool {
		return jsonFiles[i].Name < jsonFiles[j].Name
	})

	allEntry := FileInfo{
		Path:        "",
		Name:        "📊 All Tests",
		Description: "Combined view of all test files",
		TestCount:   totalTestCount,
		ParseTests:  totalParseTests,
		IsVirtual:   true,
	}
	jsonFiles = append([]FileInfo{allEntry}, jsonFiles...)

	return jsonFiles, nil
}

// runFileSelectionCLI displays a static file listing and prompts for selection
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

	header := fmt.Sprintf("📁 Available Test Files in %s", dir)
	fmt.Println(fileHeaderStyle.Render(header))
	fmt.Println()

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
	runTUI(selectedFile.Path)
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

	header := fmt.Sprintf("📁 Select Test File from %s", m.directory)
	content.WriteString(fileHeaderStyle.Render(header) + "\n\n")

	var fileList strings.Builder
	for i, file := range m.files {
		prefix := "  "
		style := lipgloss.NewStyle()

		if i == m.selectedFile {
			prefix = "► "
			style = selectedFileStyle
		}

		fileName := fmt.Sprintf("%s%s", prefix, file.Name)
		if file.IsVirtual {
			virtualStyle := style.Bold(true)
			if i != m.selectedFile {
				virtualStyle = lipgloss.NewStyle().Bold(true).Foreground(primaryColor)
			}
			fileList.WriteString(virtualStyle.Render(fileName) + "\n")
		} else {
			fileList.WriteString(style.Render(fileName) + "\n")
		}

		if i == m.selectedFile && file.Description != "" {
			desc := fmt.Sprintf("   %s", file.Description)
			descStyle := lipgloss.NewStyle().Foreground(subtleColor)
			fileList.WriteString(descStyle.Render(desc) + "\n")
		}

		if i == m.selectedFile {
			stats := fmt.Sprintf("   Total: %d tests, Parse/ParseIndented: %d tests", file.TestCount, file.ParseTests)
			statsStyle := lipgloss.NewStyle().Foreground(subtleColor)
			fileList.WriteString(statsStyle.Render(stats) + "\n")
		}

		if file.IsVirtual {
			fileList.WriteString(lipgloss.NewStyle().Foreground(subtleColor).Render("   ────────────────────────") + "\n")
		}

		fileList.WriteString("\n")
	}

	content.WriteString(fileListStyle.Render(fileList.String()) + "\n")

	help := "j/k: navigate • g/G: first/last • enter/space: select • q/esc: quit"
	content.WriteString(suiteInfoStyle.Render(help))

	return content.String()
}

// runFileSelectionTUI runs the file selection → test viewer loop.
// Uses iterative navigation to avoid unbounded recursion.
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

	for {
		model := initialFileSelectionModel(dir)
		model.files = files

		p := tea.NewProgram(model, tea.WithAltScreen())
		finalModel, err := p.Run()
		if err != nil {
			fmt.Printf("Error running file selection TUI: %v", err)
			os.Exit(1)
		}

		fsModel, ok := finalModel.(fileSelectionModel)
		if !ok || !fsModel.fileSelected || fsModel.selectedFile < 0 || fsModel.selectedFile >= len(fsModel.files) {
			return
		}

		selectedFile := fsModel.files[fsModel.selectedFile]
		wantsBack := false
		if selectedFile.IsVirtual {
			wantsBack = runTUIAndCheckBack(dir)
		} else {
			wantsBack = runTUIAndCheckBack(selectedFile.Path)
		}
		if !wantsBack {
			return
		}
	}
}

// runTUIAndCheckBack runs the test viewer TUI and returns true if the user
// wants to go back to file selection, false if they quit.
func runTUIAndCheckBack(fileOrDir string) bool {
	info, err := os.Stat(fileOrDir)
	isDir := err == nil && info.IsDir()

	if isDir {
		model := initialTUIModel()
		model.filename = ""
		model.directory = fileOrDir

		p := tea.NewProgram(&allTestsModel{tuiModel: model, dir: fileOrDir}, tea.WithAltScreen())
		finalModel, err := p.Run()
		if err != nil {
			fmt.Printf("Error running TUI: %v", err)
			os.Exit(1)
		}
		if m, ok := finalModel.(*allTestsModel); ok {
			return m.tuiModel.wantsBack
		}
		return false
	}

	model := initialTUIModel()
	model.filename = fileOrDir
	model.directory = filepath.Dir(fileOrDir)

	p := tea.NewProgram(model, tea.WithAltScreen())
	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Error running TUI: %v", err)
		os.Exit(1)
	}
	if tm, ok := finalModel.(tuiModel); ok {
		return tm.wantsBack
	}
	return false
}
