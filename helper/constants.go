package helper

import (
	"fmt"
	"os"
)

var HomePath string

func init() {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	HomePath = dir
}
