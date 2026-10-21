package config

import "strings"

func EnvIs(current string, valid ...string) bool {
	for _, env := range valid {
		if strings.EqualFold(current, env) {
			return true
		}
	}
	return false
}
