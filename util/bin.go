package util

import "os"

func BinName() string {
	bin_name := os.Args[0]
	if len(bin_name) > 24 {
		return "selene"
	}
	return bin_name
}
