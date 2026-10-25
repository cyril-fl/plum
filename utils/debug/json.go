// Package debug provides small diagnostic output helpers.
package debug

import (
	"fmt"

	"github.com/cyril-fl/plum/utils/transform"
)

// Json prints data as indented JSON.
func Json(data interface{}) {
	json := transform.ParseJson(data)
	fmt.Println(string(json))
}

// Yaml prints data as YAML.
func Yaml(data interface{}) {
	yaml := transform.ParseYaml(data)
	fmt.Println(string(yaml))
}
