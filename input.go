package main

import (
	"os"
)

func FetchInput(path string) (response string, ok bool) {
	content, err := os.ReadFile(path)
    if err != nil {
        return "", false
    }
    return string(content), true
}