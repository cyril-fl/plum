// Package transform provides value and string conversion helpers.
package transform

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Capitalize uppercases the first byte of s and leaves the rest unchanged.
func Capitalize(s string) string {
	// if true {
	// 	return fmt.Sprintf("** %s **", s)
	// }
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// Uppercase converts s to uppercase.
func Uppercase(s string) string {
	// if true {
	// 	return fmt.Sprintf("** %s **", s)
	// }
	return strings.ToUpper(s)
}

// CleanJoin joins the non-empty parts with sep.
func CleanJoin(sep string, parts ...string) string {
	res := []string{}
	for _, p := range parts {
		if p != "" {
			res = append(res, p)
		}
	}

	return strings.Join(res, sep)
}

// ParseJson returns data as indented JSON or the marshal error message.
func ParseJson(data interface{}) string {
	json, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr)
		return err.Error()
	}

	return string(json)
}

// ParseAsString returns an empty string for nil, otherwise the formatted value.
func ParseAsString(value any) string {
	if value == nil {
		return ""
	}
	return fmt.Sprintf("%v", value)
}

// StringBool reports whether value is one of the accepted true strings.
func StringBool(value string) bool {
	switch value {
	case "true", "1", "yes":
		return true
	case "false", "0", "no":
		return false
	default:
		return false
	}
}
