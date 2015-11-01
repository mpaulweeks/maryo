package main

import (
    "testing"
    "reflect"
)

func TestFileCache(t *testing.T) {
    example := MiiversePost{
        MiiverseName: "MrLuckyWaffles",
        Description: "Don't Throw the POW!",
        Code: "BBD1-0000-00C7-030C",
    }
    expected := map[string]MiiversePost{
        example.Code: example,
    }
    save_cache(cacheFileTest, expected)

    result := load_cache(cacheFileTest)
    if !reflect.DeepEqual(expected, result){
        t.Errorf("Saved != Loaded. Saved: %q, Loaded: %q", expected, result)
    }
}

func TestCompareCache(t *testing.T) {
    example := MiiversePost{
        MiiverseName: "MrLuckyWaffles",
        Description: "Don't Throw the POW!",
        Code: "BBD1-0000-00C7-030C",
    }
    expected := map[string]MiiversePost{
        example.Code: example,
    }
    save_cache(cacheFileTest, expected)

    result := load_cache(cacheFileTest)
    if !reflect.DeepEqual(expected, result){
        t.Errorf("Saved != Loaded. Saved: %q, Loaded: %q", expected, result)
    }
}
