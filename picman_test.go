package main

import "testing"

func TestIsPathAnImage(t *testing.T) {
    paths := []string{"a.png", "b.jpeg"}
    filter := new(ImagePathFilter)
    filter.addImageRegexp("png", `(?i)\.png$`)
    filter.addImageRegexp("jpeg", `(?i)\.jpeg$`)
    for _, path := range paths {
        if !filter.isImage(path) {
            t.Errorf("Expected isImage(%v) to be true.", path)
        }
    }
}