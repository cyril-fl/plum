package debug

import (
	"fmt"

	"github.com/cyril-fl/plum/utils/transform"
)

func Json(data interface{}) {
	json := transform.ParseJson(data)
	fmt.Println(string(json))
}
