//go:build windows
// +build windows

package util

import (
	"io"
	"os"
)

func install(current_location string, new_bin io.Reader) error {
	// mv current_location current_location.bak
	bak := current_location + ".bak"
	// if bak exists, delete it
	_, err := os.Stat(bak)
	if err == nil {
		err = os.Remove(bak)
		if err != nil {
			return err
		}
	}
	// rename current_location to bak
	err = os.Rename(current_location, bak)
	if err != nil {
		return err
	}

	// create new file at current_location
	out, err := os.Create(current_location)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, new_bin)
	return err
}
