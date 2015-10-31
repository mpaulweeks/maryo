package main

import (
    "testing"
)

func TestNames(t *testing.T) {
    names := get_names()
    expected := "MrLuckyWaffles"
    if !contains(names, expected){
        t.Errorf("Expected %q, got %q", expected, names)
    }
}
