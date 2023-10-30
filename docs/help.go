package docs

import (
	"fmt"
	"os"
)

func Help() {
	bin_name := os.Args[0]
	fmt.Println("Usage:")
    fmt.Println(bin_name + " [configs ...]")
	fmt.Println(bin_name + " [configs ...] -- [args ...]")
}
