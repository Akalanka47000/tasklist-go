// Package fixtures provides utility functions to load test fixture files.
// You can directly use these files for example to mock external file downloads in your tests.
package fixtures

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
)

func MustLoadFile(filePath string) io.ReadCloser {
	_, b, _, _ := runtime.Caller(0)
	f, err := os.Open(filepath.Join(b, "../", filePath)) //nolint:gosec
	if err != nil {
		panic(err)
	}
	return f
}
