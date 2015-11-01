package main

import (
    "testing"
)

func TestPostFormatting(t *testing.T) {
    posts := []MiiversePost{waff1}

    expected := "<b>NEW LEVELS</b>\n\nMrLuckyWaffles\nDon't Throw the POW!\nBBD1-0000-00C7-030C\n---\nAdd your MiiverseName here: http://notepad.cc/eti-mm"
    result := formatPosts(posts)
    if expected != result {
        t.Errorf("Expected: %q, Result: %q", expected, result)
    }
}