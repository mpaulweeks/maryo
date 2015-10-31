package main

import (
    "testing"
)

func contains(slice []string, item string) bool {
    set := make(map[string]struct{}, len(slice))
    for _, s := range slice {
        set[s] = struct{}{}
    }

    _, ok := set[item]
    return ok
}

func TestNames(t *testing.T) {
    names := get_names()
    expected := "MrLuckyWaffles"
    if !contains(names, expected){
        t.Errorf("Expected %q, got %q", expected, names)
    }
}
