package transform

import "testing"

func TestCapitalize(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "empty", input: "", want: ""},
		{name: "lowercase", input: "plum", want: "Plum"},
		{name: "already capitalized", input: "Plum", want: "Plum"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Capitalize(tt.input); got != tt.want {
				t.Errorf("Capitalize(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestUppercase(t *testing.T) {
	if got := Uppercase("Plum"); got != "PLUM" {
		t.Errorf("Uppercase(%q) = %q, want %q", "Plum", got, "PLUM")
	}
}

func TestCleanJoinSkipsEmptyParts(t *testing.T) {
	if got := CleanJoin(", ", "one", "", "two", ""); got != "one, two" {
		t.Errorf("CleanJoin() = %q, want %q", got, "one, two")
	}
}

func TestParseJson(t *testing.T) {
	if got := ParseJson(map[string]string{"name": "plum"}); got != "{\n  \"name\": \"plum\"\n}" {
		t.Errorf("ParseJson() = %q, want indented JSON", got)
	}
}

func TestParseJsonReportsMarshalErrors(t *testing.T) {
	if got := ParseJson(func() {}); got != "json: unsupported type: func()" {
		t.Errorf("ParseJson() = %q, want marshal error", got)
	}
}

func TestParseAsString(t *testing.T) {
	if got := ParseAsString(nil); got != "" {
		t.Errorf("ParseAsString(nil) = %q, want empty string", got)
	}
	if got := ParseAsString(42); got != "42" {
		t.Errorf("ParseAsString(42) = %q, want 42", got)
	}
}

func TestStringBool(t *testing.T) {
	for _, value := range []string{"true", "1", "yes"} {
		if !StringBool(value) {
			t.Errorf("StringBool(%q) = false, want true", value)
		}
	}
	for _, value := range []string{"false", "0", "no", "other"} {
		if StringBool(value) {
			t.Errorf("StringBool(%q) = true, want false", value)
		}
	}
}
