package debug

import (
	"io"
	"os"
	"testing"
)

func TestJsonPrintsIndentedJSON(t *testing.T) {
	oldStdout := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = writer
	t.Cleanup(func() {
		os.Stdout = oldStdout
	})

	Json(map[string]string{"name": "plum"})
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}

	got, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	want := "{\n  \"name\": \"plum\"\n}\n"
	if string(got) != want {
		t.Fatalf("Json() output = %q, want %q", got, want)
	}
}
