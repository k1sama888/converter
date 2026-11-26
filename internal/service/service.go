package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", fmt.Errorf("service: empty input")
	}
	if isMorse(trimmed) {
		return morse.ToText(trimmed), nil
	}
	return morse.ToMorse(trimmed), nil
}
func isMorse(s string) bool {
	if !strings.ContainsAny(s, ".-") {
		return false
	}
	replacer := strings.NewReplacer(".", "", "-", "", " ", "", "/", "")
	return replacer.Replace(s) == ""
}
func ToText(morseStr string) (string, error) {
	var unknown []string
	c := morse.NewConverter(
		morse.DefaultMorse,
		morse.WithCharSeparator(" "),
		morse.WithWordSeparator("   "),
		morse.WithTrailingSeparator(false),
		morse.WithHandler(func(err error) string {
			if e, ok := err.(morse.ErrNoEncoding); ok {
				unknown = append(unknown, e.Text)
			}
			return ""
		}),
	)

	result := strings.TrimSpace(c.ToText(morseStr))

	if len(unknown) > 0 {
		return "", fmt.Errorf("service: unknown morse codes: %s", strings.Join(unknown, ", "))
	}
	if result == "" {
		return "", fmt.Errorf("service: morse string decoded to empty result")
	}

	return result, nil
}
func ToMorse(text string) (string, error) {
	var unknown []string

	c := morse.NewConverter(
		morse.DefaultMorse,
		morse.WithCharSeparator(" "),
		morse.WithWordSeparator("   "),
		morse.WithLowercaseHandling(true),
		morse.WithTrailingSeparator(false),
		morse.WithHandler(func(err error) string {
			if e, ok := err.(morse.ErrNoEncoding); ok {
				unknown = append(unknown, e.Text)
			}
			return ""
		}),
	)

	result := strings.TrimSpace(c.ToMorse(text))

	if len(unknown) > 0 {
		return "", fmt.Errorf("service: unsupported characters: %s", strings.Join(unknown, ", "))
	}
	if result == "" {
		return "", fmt.Errorf("service: text encoded to empty result")
	}

	return result, nil
}
