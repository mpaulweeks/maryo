package main

import (
    "testing"
)

func TestStore(t *testing.T) {
    // result := load_cache()
    // if true{
    //     t.Errorf("%q", result)
    // }


    expected := MiiversePost{
        MiiverseName: "MrLuckyWaffles",
        Description: "Don't Throw the POW!",
        Code: "BBD1-0000-00C7-030C",
    }
    toSave := []MiiversePost{expected}
    save_cache(toSave)
}
