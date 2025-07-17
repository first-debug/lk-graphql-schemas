package main

import (
	"fmt"
	"os"

	protolint "github.com/yoheimuta/protolint/lib"
)

func main() {
	args := []string{"lint",
		"user-provider/proto/",
	}
	err := protolint.Lint(args, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
