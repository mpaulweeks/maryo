package main

import (
    "testing"
    "reflect"
)

func TestMiiverse(t *testing.T) {
    names := []string{"MrLuckyWaffles", "WhiteBabylon", "_"}
    result := getMiiversePosts(names)
    found := false
    expected := MiiversePost{
        MiiverseName: "MrLuckyWaffles",
        NickName: "Waff",
        Description: "Don't Throw the POW!",
        Code: "BBD1-0000-00C7-030C",
    }
    for _, mp := range result {
        found = found || reflect.DeepEqual(expected, mp)
    }
    if !found {
        t.Errorf("Expected :%q, got: %q", expected, result)
    }
}
