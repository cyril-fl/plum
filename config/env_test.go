package config

import "testing"

func TestEnvIs(t *testing.T) {
	tests := []struct {
		name    string
		current string
		valid   []string
		want    bool
	}{
		{
			name:    "exact match",
			current: "development",
			valid:   []string{"development"},
			want:    true,
		},
		{
			name:    "case insensitive match",
			current: "DeVeLoPmEnT",
			valid:   []string{"development"},
			want:    true,
		},
		{
			name:    "match among candidates",
			current: "production",
			valid:   []string{"development", "production", "test"},
			want:    true,
		},
		{
			name:    "no match",
			current: "staging",
			valid:   []string{"development", "production"},
			want:    false,
		},
		{
			name:    "no valid values",
			current: "development",
			want:    false,
		},
		{
			name:    "does not trim whitespace",
			current: " development ",
			valid:   []string{"development"},
			want:    false,
		},
		{
			name:    "empty strings match",
			current: "",
			valid:   []string{"development", ""},
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EnvIs(tt.current, tt.valid...); got != tt.want {
				t.Fatalf("EnvIs(%q, %v) = %v, want %v", tt.current, tt.valid, got, tt.want)
			}
		})
	}
}
