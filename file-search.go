package filesearch

import (
    //"path/filepath"
    "fmt"
    "os"
    "strings"
)
// Find this
func Find(stringInFilename string) (filesFoundReturn []FileInfo, err error){

    path := "\\\\vault\\movies\\"
    d, err := os.Open(path)
    defer d.Close()
    fi, err := d.Readdir(-1)
    if err != nil {
        fmt.Println("Error reading folder "+ err.Error())
    }

    fmt.Println("Searching for: " +stringInFilename)

    matchInfo := make([]FileInfo, 0, 5)
    numMatches := 0
    for _, file := range fi {
        if strings.Contains(strings.ToLower(file.Name()), strings.ToLower(stringInFilename)) {
              matchInfo = append(matchInfo, FileInfo{Name: file.Name(), Path: path})
              numMatches++
        }

    }

    return matchInfo, err

}

// FileInfo does
type FileInfo struct {
    Name string `json:"name"`
    Path string `json:"path"`
}