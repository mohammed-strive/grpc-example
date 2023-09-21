package utils

import (
	"path/filepath"
	"runtime"
)

var basepath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}

func Path(relative string) string {
	if filepath.IsAbs(relative) {
		return relative
	}

	return filepath.Join(basepath, relative)
}
