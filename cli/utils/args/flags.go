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
		case "bool":
			cmd.Flags().BoolP(
				f.Name,
				f.Shorthand,
				assertBoolValue(f.Value),
				f.Usage,
			)
		case "stringSlice":
			cmd.Flags().StringSliceP(
				f.Name,
				f.Shorthand,
				nil,
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

// GetBoolFlags reads each bool flag of cmd into its target.
func GetBoolFlags(cmd *cobra.Command, targets map[string]*bool) error {
	for name, target := range targets {
		value, err := cmd.Flags().GetBool(name)
		if err != nil {
			return err
		}
		*target = value
	}

	return nil
}

// Helpers
func assertIntValue(value string) int {
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return v
}

func assertBoolValue(value string) bool {
	v, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	return v
}
