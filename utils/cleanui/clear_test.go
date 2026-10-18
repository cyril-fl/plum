package cleanui

import (
	"io"
	"os"
	"testing"
)

func TestRunPrintsClearSequence(t *testing.T) {
	oldStdout := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = writer
	t.Cleanup(func() {
		os.Stdout = oldStdout
	})

	Run()
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}

	got, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != CLEAR {
		t.Fatalf("Run() output = %q, want %q", got, CLEAR)
	}
}
