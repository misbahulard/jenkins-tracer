package util

import (
	"os"
	"strings"
)

func CreateDirectoryByFile(path string) error {
	dirs := strings.Split(path, "/")

	// If only file name or the string is empty end the process
	if len(dirs) <= 1 {
		return nil
	}

	dirs = dirs[:len(dirs)-1]
	dir := strings.Join(dirs, "/")

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
