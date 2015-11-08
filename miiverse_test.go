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

func TestDownloadImage(t *testing.T) {
    rawUrl := "https://d3esbfg30x759i.cloudfront.net/ss/WVW69ihSwuo19HigT-"
    post := waff1
    post.ImgSrc = rawUrl
    posts := []MiiversePost{post}
    res := downloadImages(posts)
    if len(res[0].ImgFile) == 0 {
        t.Errorf("ImgFile didn't get saved: %s", res)
    }
}
