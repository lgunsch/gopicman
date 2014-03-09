package main

import (
    "fmt"
    "regexp"
)

type ImagePathFilter struct {
    patterns map[string]*regexp.Regexp
}

func (filter *ImagePathFilter) isImage(path string) bool {
    for _, regex := range filter.patterns {
        if (regex.MatchString(path)) {
            return true;
        }
    }
    return false
}

func (filter *ImagePathFilter) addImageRegexp(extension, expr string) {
    if filter.patterns == nil {
        filter.patterns = make(map[string]*regexp.Regexp)
    }
    r := regexp.MustCompile(expr)
    filter.patterns[extension] = r
}

func main() {
    fmt.Println("picman")
}