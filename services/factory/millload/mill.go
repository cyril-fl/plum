package millload

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cyril-fl/plum/primitive"
	"gopkg.in/yaml.v3"
)

type mill struct {
	primitive.Mill
	src  string
	info os.FileInfo
}

func New(src string) *mill {
	m := &mill{
		src: src,
	}

	info, err := os.Stat(m.src)
	if err != nil {
		m.SetErr(fmt.Errorf("stat %s: %w", m.src, err))
	}

	m.info = info

	return m
}

// LoadFromFile decodes a YAML or JSON document into T based on its extension.
func (m *mill) LoadFromFile[T any]() *T {
	var params T

	switch strings.ToLower(filepath.Ext(m.src)) {
	case ".yaml", ".yml":
		m.loadYAMLFromFile(&params)
	case ".json":
		m.loadJSONFromFile(&params)
	default:
		m.SetErr(fmt.Errorf("unsupported file extension: %s", m.src))
		return nil
	}

	if m.HasErr() {
		return nil
	}

	return &params
}

// LoadListFromFile decodes a YAML or JSON sequence into []T.
func (m *mill) LoadListFromFile[T any]() []T {
	params := m.LoadFromFile[[]T]()

	if params == nil {
		return nil
	}

	return *params
}

// loadYAMLFromFile decodes the source file as YAML into target.
func (m *mill) loadYAMLFromFile(target any) {
	data, err := os.ReadFile(m.src)
	if err != nil {
		m.SetErr(fmt.Errorf("read %s: %w", m.src, err))
	} else if err := yaml.Unmarshal(data, target); err != nil {
		m.SetErr(fmt.Errorf("unmarshal %s: %w", m.src, err))
	}
}

// loadJSONFromFile decodes the source file as JSON into target.
func (m *mill) loadJSONFromFile(target any) {
	data, err := os.ReadFile(m.src)
	if err != nil {
		m.SetErr(fmt.Errorf("read %s: %w", m.src, err))
	} else if err := json.Unmarshal(data, target); err != nil {
		m.SetErr(fmt.Errorf("unmarshal %s: %w", m.src, err))
	}
}
