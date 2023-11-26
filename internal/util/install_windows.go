//go:build windows

package util

import (
	"io"
	"os"
)

func install(currentLocation string, newBin io.Reader) error {
	// mv current_location current_location.bak
	bak := currentLocation + ".bak"
	// if bak exists, delete it
	_, err := os.Stat(bak)
	if err == nil {
		err = os.Remove(bak)
		if err != nil {
			return err
		}
	}
	// rename current_location to bak
	err = os.Rename(currentLocation, bak)
	if err != nil {
		return err
	}

	// create new file at current_location
	out, err := os.Create(currentLocation)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, newBin)
	return err
}
