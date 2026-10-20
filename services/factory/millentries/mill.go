package millentries

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/cyril-fl/plum/primitive"
)

type Mill = mill

type mill struct {
	primitive.Mill
	fileAddress string
}

func New() *mill {
	return &mill{}
}

// GetFilesFromSource returns the source file or files in the source directory,
// including subdirectories if recursive is true.
func (m *mill) GetFilesFromSource(src string, recursive bool) []string {
	info, err := os.Stat(src)
	if err != nil {
		m.SetErr(fmt.Errorf("%q: %w", src, err))
		return nil
	}

	if !info.IsDir() {
		return []string{src}
	}

	var result []string

	m.SetErr(filepath.WalkDir(src, func(path string, entry fs.DirEntry, err error) error {
		switch {
		case err != nil:
			m.SetErr(fmt.Errorf("%q: %w", path, err))
		case !entry.IsDir():
			result = append(result, path)
		case path != src && !recursive:
			return fs.SkipDir
		}

		return nil
	}))

	return result
}

// GetFolderFiles returns the non-directory entries in src.
func (m *mill) GetFolderFiles(src string) []string {
	return m.crawlFolderEntries(src, isFile)
}

// GetFolderDirectories returns the directories in src.
func (m *mill) GetFolderDirectories(src string) []string {
	return m.crawlFolderEntries(src, isDirectory)
}

// crawlFolderEntries returns entries in src matching predicate.
func (m *mill) crawlFolderEntries(src string, predicate predicate) []string {
	entries, err := os.ReadDir(src)
	if err != nil {
		m.SetErr(fmt.Errorf("%q: %w", src, err))
		return nil
	}

	result := make([]string, 0, len(entries))
	for _, entry := range entries {
		if predicate(entry) {
			result = append(result, entry.Name())
		}
	}

	return result
}

// predicate reports whether an entry should be kept.
type predicate func(entry os.DirEntry) bool

// isDirectory reports whether an entry is a directory.
func isDirectory(entry os.DirEntry) bool {
	return entry.IsDir()
}

// isFile reports whether an entry is not a directory.
func isFile(entry os.DirEntry) bool {
	return !entry.IsDir()
}
