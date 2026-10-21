package cliutilsargs

import (
	"strconv"

	config "github.com/cyril-fl/plum/cli/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func Makeflags(cmd *cobra.Command, flags []config.Cmdflags) []*pflag.FlagSet {
	var result []*pflag.FlagSet
	for _, f := range flags {

		switch f.Type {
		case "int":
			cmd.Flags().IntP(
				f.Name,
				f.Shorthand,
				assertIntValue(f.Value),
				f.Usage,
			)
		case "string":
			cmd.Flags().StringP(
				f.Name,
				f.Shorthand,
				f.Value,
				f.Usage,
			)
		default:
			cmd.Flags().StringP(
				f.Name,
				f.Shorthand,
				"",
				f.Usage,
			)
		}
	}
	return result
}

// Helpers
func assertIntValue(value string) int {
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return v
}
