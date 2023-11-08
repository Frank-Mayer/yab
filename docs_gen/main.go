package main

import (
	"os"

	"selene.frankmayer.dev/docs"
)

func main() {
	docs_str := docs.Markdown()

	filename := "DOCS.md"

	// check if file exists
	if _, err := os.Stat(filename); err == nil {
		// overwrite
		file, err := os.OpenFile(filename, os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = file.WriteString(docs_str)
		if err != nil {
			panic(err)
		}
	} else if os.IsNotExist(err) {
		// create
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = file.WriteString(docs_str)
		if err != nil {
			panic(err)
		}
	} else {
		// other error
		panic(err)
	}
}
