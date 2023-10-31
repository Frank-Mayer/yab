package util

import (
	"os"
	"strconv"
)

func TermWidth() int {
	term_width := os.Getenv("COLUMNS")

	width, err := strconv.Atoi(term_width)
	if err != nil {
		return 80
	}

	return width
}
