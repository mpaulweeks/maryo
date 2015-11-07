package main

import (
    "testing"
    "reflect"
)

var waff1 = MiiversePost{
    MiiverseName: "MrLuckyWaffles",
    NickName: "Waff",
    Description: "Don't Throw the POW!",
    Code: "BBD1-0000-00C7-030C",
}

var waff2 = MiiversePost{
    MiiverseName: "MrLuckyWaffles",
    NickName: "Waff",
    Description: "Amiibo Castle",
    Code: "91BB-0000-00CE-4B77",
}

var babylon1 = MiiversePost{
    MiiverseName: "WhiteBabylon",
    NickName: "Babylon",
    Description: "red shell god",
    Code: "74DE-0000-0015-0F80",
}

func TestFileCache(t *testing.T) {
    expected := map[string]MiiversePost{
        waff1.Code: waff1,
    }
    saveCache(cacheFileTest, expected)
    result := loadCache(cacheFileTest)
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
    result := filterNewPosts(cache, fetched)
    if !reflect.DeepEqual(expected, result){
        t.Errorf("Expected: %q, Result: %q", expected, result)
    }
}

func TestFilterEmpty(t *testing.T) {
    cache := map[string]MiiversePost{
        waff1.Code: waff1,
    }
    fetched := []MiiversePost{waff1, babylon1}
    result := filterNewPosts(cache, fetched)
    if len(result) > 0 {
        t.Errorf("Expected: [], Result: %q", result)
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
    updateCache(cache, fetched)
    if !reflect.DeepEqual(expected, cache){
        t.Errorf("Expected: %q, Result: %q", expected, cache)
    }
}

func TestUpdateEmpty(t *testing.T) {
    cache := map[string]MiiversePost{}
    fetched := []MiiversePost{waff1, waff2, babylon1}
    expected := map[string]MiiversePost{
        waff1.Code: waff1,
        waff2.Code: waff2,
        babylon1.Code: babylon1,
    }
    updateCache(cache, fetched)
    if !reflect.DeepEqual(expected, cache){
        t.Errorf("Expected: %q, Result: %q", expected, cache)
    }
}
