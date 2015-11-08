package main

import (
    "testing"
    "reflect"
)

func TestMiiverse(t *testing.T) {
    names := []string{"MrLuckyWaffles", "WhiteBabylon", "_"}
    result := getMiiversePosts(names)
    found := false
    expected := waff1
    for _, mp := range result {
        found = found || reflect.DeepEqual(expected, mp)
    }
    if !found {
        t.Errorf("Expected :%q, got: %q", expected, result)
    }
}
