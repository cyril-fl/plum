package millload

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

type loadConfig struct {
	Name  string `yaml:"name" json:"name"`
	Count int    `yaml:"count" json:"count"`
}

func TestLoadFromFileYAML(t *testing.T) {
	file := writeLoadTestFile(t, "config.yaml", "name: campaign\ncount: 3\n")

	m := New(file)
	got := m.LoadFromFile[loadConfig]()

	if m.HasErr() {
		t.Fatalf("LoadFromFile() error = %v", m.GetErr())
	}

	want := &loadConfig{Name: "campaign", Count: 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("LoadFromFile() = %#v, want %#v", got, want)
	}
}

func TestLoadFromFileJSON(t *testing.T) {
	file := filepath.Join(t.TempDir(), "config.json")
	contents, err := json.Marshal(loadConfig{Name: "campaign", Count: 3})
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(file, contents, 0o644); err != nil {
		t.Fatal(err)
	}

	m := New(file)
	got := m.LoadFromFile[loadConfig]()

	if m.HasErr() {
		t.Fatalf("LoadFromFile() error = %v", m.GetErr())
	}

	want := &loadConfig{Name: "campaign", Count: 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("LoadFromFile() = %#v, want %#v", got, want)
	}
}

func TestLoadListFromFile(t *testing.T) {
	file := writeLoadTestFile(t, "items.yml", "- alpha\n- beta\n")

	m := New(file)
	got := m.LoadListFromFile[string]()

	if m.HasErr() {
		t.Fatalf("LoadListFromFile() error = %v", m.GetErr())
	}

	want := []string{"alpha", "beta"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("LoadListFromFile() = %#v, want %#v", got, want)
	}
}

func TestLoadFromFileReturnsNilForMissingFile(t *testing.T) {
	file := filepath.Join(t.TempDir(), "missing.yaml")

	m := New(file)
	got := m.LoadFromFile[loadConfig]()

	if got != nil {
		t.Fatalf("LoadFromFile() = %#v, want nil", got)
	}
	if !m.HasErr() || !strings.Contains(m.GetErr().Error(), file) {
		t.Fatalf("LoadFromFile() error = %v, want an error containing %q", m.GetErr(), file)
	}
}

func TestLoadFromFileReturnsNilForInvalidDocument(t *testing.T) {
	file := writeLoadTestFile(t, "invalid.yaml", "name: [invalid\n")

	m := New(file)
	got := m.LoadFromFile[loadConfig]()

	if got != nil {
		t.Fatalf("LoadFromFile() = %#v, want nil", got)
	}
	if !m.HasErr() || !strings.Contains(m.GetErr().Error(), file) {
		t.Fatalf("LoadFromFile() error = %v, want an error containing %q", m.GetErr(), file)
	}
}

func TestLoadFromFileReturnsNilForUnsupportedExtension(t *testing.T) {
	file := writeLoadTestFile(t, "config.toml", "name = 'campaign'\n")

	m := New(file)
	got := m.LoadFromFile[loadConfig]()

	if got != nil {
		t.Fatalf("LoadFromFile() = %#v, want nil", got)
	}
}

func writeLoadTestFile(t *testing.T, name, contents string) string {
	t.Helper()

	file := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(file, []byte(contents), 0o644); err != nil {
		t.Fatal(err)
	}

	return file
}
