package millentries_test

import (
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/cyril-fl/plum/services/factory/millentries"
)

func TestNewReturnsAMillWithoutErrors(t *testing.T) {
	m := millentries.New()

	if m.HasErr() {
		t.Fatalf("New() unexpectedly returned an error: %v", m.GetErr())
	}
}

func TestGetFilesFromSourceHandlesFilesAndDirectories(t *testing.T) {
	root := t.TempDir()
	rootFile := filepath.Join(root, "root.txt")
	nested := filepath.Join(root, "nested")
	nestedFile := filepath.Join(nested, "nested.txt")
	writeMillTestFile(t, rootFile)
	if err := os.Mkdir(nested, 0o755); err != nil {
		t.Fatal(err)
	}
	writeMillTestFile(t, nestedFile)

	m := millentries.New()
	assertMillPaths(t, m.GetFilesFromSource(root, false), rootFile)
	assertMillPaths(t, m.GetFilesFromSource(root, true), rootFile, nestedFile)
	assertMillPaths(t, m.GetFilesFromSource(rootFile, true), rootFile)

	if m.HasErr() {
		t.Fatalf("GetFilesFromSource() unexpectedly returned an error: %v", m.GetErr())
	}
}

func TestGetFilesFromSourceReportsMissingSource(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "missing")
	m := millentries.New()

	if got := m.GetFilesFromSource(missing, false); got != nil {
		t.Fatalf("GetFilesFromSource() = %v, want nil", got)
	}
	if !m.HasErr() || !strings.Contains(m.GetErr().Error(), missing) {
		t.Fatalf("GetFilesFromSource() error = %v, want an error containing %q", m.GetErr(), missing)
	}
}

func TestGetFolderEntriesReturnsOnlyDirectEntries(t *testing.T) {
	root := t.TempDir()
	writeMillTestFile(t, filepath.Join(root, "b.txt"))
	writeMillTestFile(t, filepath.Join(root, "a.txt"))
	nested := filepath.Join(root, "nested")
	if err := os.Mkdir(nested, 0o755); err != nil {
		t.Fatal(err)
	}
	writeMillTestFile(t, filepath.Join(nested, "hidden-from-root.txt"))

	m := millentries.New()
	assertMillPaths(t, m.GetFolderFiles(root), "a.txt", "b.txt")
	assertMillPaths(t, m.GetFolderDirectories(root), "nested")

	if m.HasErr() {
		t.Fatalf("folder methods unexpectedly returned an error: %v", m.GetErr())
	}
}

func TestGetFolderEntriesReportsReadErrors(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "missing")
	m := millentries.New()

	if got := m.GetFolderFiles(missing); got != nil {
		t.Fatalf("GetFolderFiles() = %v, want nil", got)
	}
	if !m.HasErr() || !strings.Contains(m.GetErr().Error(), missing) {
		t.Fatalf("GetFolderFiles() error = %v, want an error containing %q", m.GetErr(), missing)
	}

	m.ClearErr()
	if got := m.GetFolderDirectories(missing); got != nil {
		t.Fatalf("GetFolderDirectories() = %v, want nil", got)
	}
	if !m.HasErr() || !strings.Contains(m.GetErr().Error(), missing) {
		t.Fatalf("GetFolderDirectories() error = %v, want an error containing %q", m.GetErr(), missing)
	}
}

func assertMillPaths(t *testing.T, got []string, want ...string) {
	t.Helper()
	gotSorted := append([]string(nil), got...)
	wantSorted := append([]string(nil), want...)
	sort.Strings(gotSorted)
	sort.Strings(wantSorted)
	if !reflect.DeepEqual(gotSorted, wantSorted) {
		t.Fatalf("paths = %v, want %v", gotSorted, wantSorted)
	}
}

func writeMillTestFile(t *testing.T, path string) {
	t.Helper()
	if err := os.WriteFile(path, []byte("test"), 0o644); err != nil {
		t.Fatal(err)
	}
}
