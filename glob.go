package main

import (
    "fmt"
    "path/filepath"
)

func main() {
    // Get the glob pattern.
    pattern := "*.txt"

    // Get the list of files that match the pattern.
    files, err := filepath.Glob(pattern)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Print the list of files.
    for _, file := range files {
        fmt.Println(file)
    }
}
