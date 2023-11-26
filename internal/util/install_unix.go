//go:build unix

package util

import (
	"io"
	"os"
)

func install(currentLocation string, newBin io.Reader) error {
	out, err := os.OpenFile(currentLocation, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, newBin)
	return err
}
