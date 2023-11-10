//go:build unix
// +build unix

package util

import (
	"io"
	"os"
)

func install(current_location string, new_bin io.Reader) error {
	out, err := os.OpenFile(current_location, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, new_bin)
	return err
}
