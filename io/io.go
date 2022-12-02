// Package io contains input parsers.
package io

import (
	"io"
	"os"
)

func OpenAndReadAll(filename string) ([]byte, error) {
	f, err := os.Open("day1/input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}
