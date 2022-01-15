package ls

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

func TestListDirectory(t *testing.T) {
	var testCases = []struct {
		name       string
		longFormat bool
	}{
		{"Normal Output", false},
		{"Long Output", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tempDir, err := os.MkdirTemp("", "iamtemporary")
			if err != nil {
				t.Error(err)
			}

			var expected []string
			var tempFiles []*os.File
			for i := 0; i < 5; i++ {
				tempFile, err := os.CreateTemp(tempDir, "*.txt")
				if err != nil {
					t.Error(err)
				}

				expected = append(expected, filepath.Base(tempFile.Name()))
				tempFiles = append(tempFiles, tempFile)
			}

			cfg := Config{
				LongFormat: tc.longFormat,
				Directory:  tempDir,
			}

			out := &bytes.Buffer{}
			err = Run(cfg, out)
			if err != nil {
				return
			}

			// fetch the names from ls.Run
			separator := getSeparator(cfg.LongFormat)
			results := strings.Split(out.String(), separator)

			expected = sanitize(expected)
			results = sanitize(results)

			// sort the slices
			sort.Strings(expected)
			sort.Strings(results)

			// compare for equality
			assert.Equal(t, expected, results, "directory contents should be equal")
		})
	}
}

// sanitize removes newlines from strings and delete any empty strings from a slice
func sanitize(slice []string) []string {
	var ret []string

	for i := 0; i < len(slice); i++ {
		s := strings.TrimSuffix(slice[i], "\n")
		if s != "" {
			ret = append(ret, s)
		}
	}

	return ret
}
