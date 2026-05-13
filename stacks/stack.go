package stacks

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Options struct {
	IncludeModels  bool
	IncludeWidgets bool
}

type Stack interface {
	Name() string
	Generate(cwd, featureName string, opts Options) error
}

// ParseFeatureName accepts any common casing (snake_case, camelCase, PascalCase,
// space-separated, hyphen-separated) and returns all three target forms.
func ParseFeatureName(raw string) (pascal, camel, snake string) {
	words := tokenize(raw)
	return toPascalCase(words), toCamelCase(words), toSnakeCase(words)
}

func tokenize(s string) []string {
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.ReplaceAll(s, "-", " ")
	coarse := strings.Fields(s)

	var words []string
	for _, token := range coarse {
		words = append(words, splitCamel(token)...)
	}
	return words
}

func splitCamel(s string) []string {
	runes := []rune(s)
	var words []string
	start := 0
	for i := 1; i < len(runes); i++ {
		prev := runes[i-1]
		curr := runes[i]
		var next rune
		if i+1 < len(runes) {
			next = runes[i+1]
		}

		cut := false
		if isLower(prev) && isUpper(curr) {
			cut = true
		} else if isUpper(prev) && isUpper(curr) && next != 0 && isLower(next) {
			cut = true
		}

		if cut {
			words = append(words, strings.ToLower(string(runes[start:i])))
			start = i
		}
	}
	words = append(words, strings.ToLower(string(runes[start:])))
	return words
}

func isUpper(r rune) bool { return r >= 'A' && r <= 'Z' }
func isLower(r rune) bool { return r >= 'a' && r <= 'z' }

func toPascalCase(words []string) string {
	var b strings.Builder
	for _, w := range words {
		if len(w) == 0 {
			continue
		}
		b.WriteString(strings.ToUpper(w[:1]) + w[1:])
	}
	return b.String()
}

func toCamelCase(words []string) string {
	if len(words) == 0 {
		return ""
	}
	var b strings.Builder
	for i, w := range words {
		if len(w) == 0 {
			continue
		}
		if i == 0 {
			b.WriteString(strings.ToLower(w))
		} else {
			b.WriteString(strings.ToUpper(w[:1]) + w[1:])
		}
	}
	return b.String()
}

func toSnakeCase(words []string) string {
	return strings.Join(words, "_")
}

// WriteFile creates all parent directories then writes content to path.
func WriteFile(path, content string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("cannot create directory %s: %w", dir, err)
	}
	return os.WriteFile(path, []byte(content), 0644)
}
