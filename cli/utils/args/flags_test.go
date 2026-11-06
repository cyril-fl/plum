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

func TestGetBoolFlags(t *testing.T) {
	cmd := &cobra.Command{}
	Makeflags(cmd, []config.Cmdflags{{Name: "json", Type: "bool"}, {Name: "pretty", Type: "bool"}})
	if err := cmd.ParseFlags([]string{"--pretty"}); err != nil {
		t.Fatalf("ParseFlags() error = %v", err)
	}

	var asJSON, asPretty bool
	if err := GetBoolFlags(cmd, map[string]*bool{"json": &asJSON, "pretty": &asPretty}); err != nil {
		t.Fatalf("GetBoolFlags() error = %v", err)
	}
	if asJSON || !asPretty {
		t.Fatalf("json = %v, pretty = %v; want false, true", asJSON, asPretty)
	}

	if err := GetBoolFlags(cmd, map[string]*bool{"missing": &asJSON}); err == nil {
		t.Fatal("GetBoolFlags() error = nil, want undefined flag error")
	}
}

func TestMakeflagsStringSlice(t *testing.T) {
	cmd := &cobra.Command{}
	Makeflags(cmd, []config.Cmdflags{{Name: "tag", Type: "stringSlice"}})
	if err := cmd.ParseFlags([]string{"--tag", "a,b", "--tag", "c"}); err != nil {
		t.Fatalf("ParseFlags() error = %v", err)
	}

	got, err := cmd.Flags().GetStringSlice("tag")
	if err != nil || len(got) != 3 || got[0] != "a" || got[1] != "b" || got[2] != "c" {
		t.Fatalf("GetStringSlice(tag) = %v, %v; want [a b c], nil", got, err)
	}
}
