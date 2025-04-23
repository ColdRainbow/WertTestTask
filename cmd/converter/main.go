package main

import (
	"converter/internal/cli"
	"fmt"
	"os"
)

func main() {
	if err := cli.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
