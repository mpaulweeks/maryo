package main

import (
    "testing"
)

func TestNames(t *testing.T) {
    names := getMiiverseNames()
    expected := "MrLuckyWaffles"
    if !contains(names, expected){
        t.Errorf("Expected %q, got %q", expected, names)
    }
}
