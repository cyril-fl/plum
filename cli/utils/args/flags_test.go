package cliutilsargs

import (
	"testing"

	config "github.com/cyril-fl/plum/cli/config"
	"github.com/spf13/cobra"
)

func TestMakeflagsBool(t *testing.T) {
	cmd := &cobra.Command{}
	Makeflags(cmd, []config.Cmdflags{
		{Name: "json", Type: "bool"},
		{Name: "verbose", Shorthand: "v", Value: "true", Type: "bool"},
	})

	if err := cmd.ParseFlags([]string{"--json"}); err != nil {
		t.Fatalf("ParseFlags() error = %v", err)
	}

	if got, err := cmd.Flags().GetBool("json"); err != nil || !got {
		t.Fatalf("GetBool(json) = %v, %v; want true, nil", got, err)
	}
	if got, err := cmd.Flags().GetBool("verbose"); err != nil || !got {
		t.Fatalf("GetBool(verbose) = %v, %v; want default true, nil", got, err)
	}
}
