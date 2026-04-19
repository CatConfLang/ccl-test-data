// Command validate-tags enforces the tag-URL contract between this test suite
// and the CCL documentation site.
//
// It walks source_tests/, collects every function/feature/behavior/variant tag
// referenced by tests, fetches the canonical tag-index.json emitted by the
// website build, and fails if any tag used in tests is missing from the index.
//
// The canonical tag-index lives at https://catconflang.com/tag-index.json.
// If the fetch fails, a cached copy at docs/tag-index.json is used as a
// fallback so CI doesn't break when the website is mid-deploy.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"time"
)

const (
	defaultIndexURL  = "https://catconflang.com/tag-index.json"
	defaultCachePath = "docs/tag-index.json"
	fetchTimeout     = 10 * time.Second
)

type tagEntry struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type tagIndex struct {
	Version int                 `json:"version"`
	Site    string              `json:"site"`
	Tags    map[string]tagEntry `json:"tags"`
}

func main() {
	var (
		indexURL  = flag.String("url", defaultIndexURL, "Canonical tag-index.json URL")
		cachePath = flag.String("cache", defaultCachePath, "Fallback cached tag-index.json path")
		testsDir  = flag.String("tests", "source_tests", "Directory containing test JSON files")
		offline   = flag.Bool("offline", false, "Skip the fetch and use the cache directly")
	)
	flag.Parse()

	index, source, err := loadIndex(*indexURL, *cachePath, *offline)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: loading tag-index: %v\n", err)
		os.Exit(2)
	}
	fmt.Printf("tag-index loaded from %s (%d tags, site=%s)\n", source, len(index.Tags), index.Site)

	usedTags, err := collectTags(*testsDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: walking tests: %v\n", err)
		os.Exit(2)
	}
	fmt.Printf("collected %d distinct tags from %s/\n", len(usedTags), *testsDir)

	missing := setDifference(usedTags, index.Tags)
	unused := setDifference(index.Tags, usedTags)

	if len(missing) > 0 {
		fmt.Fprintf(os.Stderr, "\nFAIL: %d tag(s) used in tests have no canonical documentation URL:\n", len(missing))
		for _, tag := range missing {
			fmt.Fprintf(os.Stderr, "  %s\n", tag)
		}
		fmt.Fprintln(os.Stderr, "\nEither document the tag on the CCL website (update src/data/tags.ts)")
		fmt.Fprintln(os.Stderr, "or remove the tag from the affected test files.")
		os.Exit(1)
	}

	if len(unused) > 0 {
		fmt.Printf("\nWARN: %d tag(s) documented but not used in any test (informational):\n", len(unused))
		for _, tag := range unused {
			fmt.Printf("  %s\n", tag)
		}
	}

	fmt.Println("\nOK: every test tag resolves to a canonical documentation URL.")
}

func loadIndex(url, cachePath string, offline bool) (*tagIndex, string, error) {
	if !offline {
		index, err := fetchIndex(url)
		if err == nil {
			return index, url, nil
		}
		fmt.Fprintf(os.Stderr, "warning: fetch %s failed (%v); falling back to %s\n", url, err, cachePath)
	}
	data, err := os.ReadFile(cachePath)
	if err != nil {
		return nil, "", fmt.Errorf("no cached index at %s: %w", cachePath, err)
	}
	var index tagIndex
	if err := json.Unmarshal(data, &index); err != nil {
		return nil, "", fmt.Errorf("parsing %s: %w", cachePath, err)
	}
	return &index, cachePath, nil
}

func fetchIndex(url string) (*tagIndex, error) {
	client := &http.Client{Timeout: fetchTimeout}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var index tagIndex
	if err := json.Unmarshal(body, &index); err != nil {
		return nil, err
	}
	return &index, nil
}

// collectTags walks every JSON file under dir and returns the set of
// function/feature/behavior/variant tags used in tests, in canonical
// "category:name" form.
func collectTags(dir string) (map[string]struct{}, error) {
	tags := map[string]struct{}{}
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".json" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading %s: %w", path, err)
		}
		var raw any
		if err := json.Unmarshal(data, &raw); err != nil {
			return fmt.Errorf("parsing %s: %w", path, err)
		}
		walkTags(raw, tags)
		return nil
	})
	return tags, err
}

// tagPrefixByKey maps metadata array keys to their canonical tag prefix.
var tagPrefixByKey = map[string]string{
	"functions": "function",
	"features":  "feature",
	"behaviors": "behavior",
	"variants":  "variant",
}

// walkTags recurses into arbitrary JSON values and harvests tag strings from
// known tag-bearing keys.
//
// Critically, it does NOT recurse into data-carrying fields like `expect`,
// `input`, `inputs`, or `args`: a `features` key inside `expect.data` is
// test output (the *content* of a CCL document being parsed), not test
// metadata. Recursing into them causes false-positive "tags" like the
// literal keys of a test fixture.
func walkTags(v any, out map[string]struct{}) {
	switch x := v.(type) {
	case map[string]any:
		for key, val := range x {
			if prefix, ok := tagPrefixByKey[key]; ok {
				addStrings(val, prefix, out)
				continue
			}
			switch key {
			case "conflicts":
				if c, ok := val.(map[string]any); ok {
					for k, v := range c {
						if prefix, ok := tagPrefixByKey[k]; ok {
							addStrings(v, prefix, out)
						}
					}
				}
			case "expect", "input", "inputs", "args":
				// Test data, not metadata — skip recursion.
			default:
				walkTags(val, out)
			}
		}
	case []any:
		for _, item := range x {
			walkTags(item, out)
		}
	}
}

func addStrings(v any, prefix string, out map[string]struct{}) {
	arr, ok := v.([]any)
	if !ok {
		return
	}
	for _, item := range arr {
		if s, ok := item.(string); ok {
			out[prefix+":"+s] = struct{}{}
		}
	}
}

// setDifference returns keys present in a but not b, sorted.
func setDifference[A, B any](a map[string]A, b map[string]B) []string {
	var out []string
	for k := range a {
		if _, ok := b[k]; !ok {
			out = append(out, k)
		}
	}
	sort.Strings(out)
	return out
}
