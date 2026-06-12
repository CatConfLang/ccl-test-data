package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/catconflang/ccl-test-data/loader"
	"github.com/charmbracelet/lipgloss"
)

// processTestFile renders a test file to stdout in static (non-TUI) mode
func processTestFile(filename string) error {
	if strings.Contains(filename, "source_tests") {
		warningStyle := lipgloss.NewStyle().Foreground(warningColor).Bold(true)
		fmt.Println(warningStyle.Render("⚠️  WARNING: Source format files have limited display support."))
		fmt.Println(warningStyle.Render("   Use generated format files for full entry details."))
		fmt.Println(warningStyle.Render("   Run 'just generate' to convert source tests to generated format."))
		fmt.Println()
	}

	testLoader := loader.NewTestLoader(".", readerConfig)
	suite, err := testLoader.LoadTestFile(filename, loader.LoadOptions{
		Format:     loader.FormatFlat,
		FilterMode: loader.FilterAll,
	})
	if err != nil {
		return fmt.Errorf("loading test suite: %w", err)
	}

	fmt.Println(suiteHeaderStyle.Render(suite.Suite))
	info := fmt.Sprintf("File: %s | %s", filepath.Base(filename), suite.Description)
	fmt.Println(suiteInfoStyle.Render(info))
	fmt.Println()

	for i, test := range suite.Tests {
		displayTest(test, i+1)
	}

	summary := fmt.Sprintf("📊 Found %d test(s)", len(suite.Tests))
	fmt.Println(summaryStyle.Render(summary))
	fmt.Println()

	return nil
}

func displayTest(test TestCase, index int) {
	header := fmt.Sprintf("Test #%d: %s", index, test.Name)
	fmt.Println(testHeaderStyle.Render(header))

	fmt.Println(inputHeaderStyle.Render("📄 CCL INPUT:"))
	inputText := ""
	if len(test.Inputs) > 0 {
		inputText = test.Inputs[0]
	}
	fmt.Println(inputContentStyle.Render(formatInputContent(inputText)))

	displayExpectedOutput(test)
	displaySelectiveMetadata(test)
	fmt.Println()
}

func displayExpectedOutput(test TestCase) {
	result := renderExpectedContent(test, 0, maxEntriesDisplay)

	if result.IsError {
		fmt.Println(errorHeaderStyle.Render(result.Header))
		return
	}
	fmt.Println(successHeaderStyle.Render(result.Header))

	if result.IsList {
		fmt.Printf("   Items: %d\n", result.TotalItems)
	} else if !result.IsObject {
		fmt.Printf("   Count: %d assertion(s)\n", result.TotalItems)
		if result.TotalItems > 0 {
			fmt.Printf("   Entries (%d total):\n", result.TotalItems)
		}
	}

	if (result.IsObject || result.IsList) && len(result.Lines) > 0 {
		var boxContent strings.Builder
		for i := result.StartIdx; i < result.EndIdx; i++ {
			boxContent.WriteString(result.Lines[i])
			if i < result.EndIdx-1 {
				boxContent.WriteString("\n")
			}
		}
		fmt.Println(objectBoxStyle.Render(boxContent.String()))
	} else {
		for i := result.StartIdx; i < result.EndIdx; i++ {
			fmt.Println(result.Lines[i])
		}
	}

	if result.HasMoreBelow {
		remaining := result.TotalItems - result.EndIdx
		itemType := "entries"
		if result.IsObject {
			itemType = "lines"
		} else if result.IsList {
			itemType = "items"
		}
		truncationMsg := fmt.Sprintf("... and %d more %s (use TUI mode for scrolling)", remaining, itemType)
		truncationStyle := lipgloss.NewStyle().Foreground(subtleColor)
		fmt.Println(truncationStyle.Render("   " + truncationMsg))
	}
}

func displaySelectiveMetadata(test TestCase) {
	variantTags := []string{}

	for _, behavior := range test.Behaviors {
		variantTags = append(variantTags, "behavior:"+behavior)
	}
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
}
