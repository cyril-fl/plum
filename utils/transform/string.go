// Package transform provides value and string conversion helpers.
package transform

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
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

// ParseYaml returns data as YAML indented with two spaces or the marshal error message.
func ParseYaml(data interface{}) (out string) {
	// yaml.v3 panics on unsupported types where encoding/json returns an error.
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr)
			out = fmt.Sprint(r)
		}
	}()

	// Two-space indent, like ParseJson.
	var buffer bytes.Buffer
	encoder := yaml.NewEncoder(&buffer)
	encoder.SetIndent(2)
	if err := encoder.Encode(data); err != nil {
		fmt.Fprintln(os.Stderr)
		return err.Error()
	}
	encoder.Close()

	// Trimmed like ParseJson output, which has no trailing newline.
	return strings.TrimSuffix(buffer.String(), "\n")
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
