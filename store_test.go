package main

import (
    "testing"
    "reflect"
)

var waff1 = MiiversePost{
    MiiverseName: "MrLuckyWaffles",
    Description: "Don't Throw the POW!",
    Code: "BBD1-0000-00C7-030C",
}

var waff2 = MiiversePost{
    MiiverseName: "MrLuckyWaffles",
    Description: "Amiibo Castle",
    Code: "91BB-0000-00CE-4B77",
}

var babylon1 = MiiversePost{
    MiiverseName: "WhiteBabylon",
    Description: "red shell god",
    Code: "74DE-0000-0015-0F80",
}

func TestFileCache(t *testing.T) {
    expected := map[string]MiiversePost{
        waff1.Code: waff1,
    }
    save_cache(cacheFileTest, expected)

    result := load_cache(cacheFileTest)
    if !reflect.DeepEqual(expected, result){
        t.Errorf("Saved != Loaded. Saved: %q, Loaded: %q", expected, result)
    }
}

func TestFilter(t *testing.T) {
    cache := map[string]MiiversePost{
        waff1.Code: waff1,
    }
    fetched := []MiiversePost{waff1, waff2, babylon1}
    expected := []MiiversePost{waff2}
    result := filter_updates(cache, fetched)
    if !reflect.DeepEqual(expected, result){
        t.Errorf("Expected: %q, Result: %q", expected, result)
    }
}

func TestUpdate(t *testing.T) {
    cache := map[string]MiiversePost{
        waff1.Code: waff1,
    }
    fetched := []MiiversePost{waff1, waff2, babylon1}
    expected := map[string]MiiversePost{
        waff1.Code: waff1,
        waff2.Code: waff2,
        babylon1.Code: babylon1,
    }
    update_cache(cache, fetched)
    if !reflect.DeepEqual(expected, cache){
        t.Errorf("Expected: %q, Result: %q", expected, cache)
    }
}
