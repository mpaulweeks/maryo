package main

import (
    "testing"
)

var sampleUserData = UserData{
    MiiverseNames: []string{"MrLuckyWaffles", "WhiteBabylon"},
    NickNames: map[string]string{"Waff": "Mr Lucky Waffles"},
}

func TestPostFormatting(t *testing.T) {
    cred := loadCredentials(credFileTest)
    userData := sampleUserData
    posts := []MiiversePost{waff1}

    expected := "<b>NEW LEVELS</b>\n\nMr Lucky Waffles\nDon't Throw the POW!\nBBD1-0000-00C7-030C\nhttps://miiverse.nintendo.net/posts/AYMHAAACAAADVHkkRVNLAA\n---\nRegister your MiiverseName here: " + cred.RegisterUrl
    result := formatPosts(cred, userData, posts)
    if expected != result {
        t.Errorf("Expected: %q, Result: %q", expected, result)
    }
}

func TestGetForumKey(t *testing.T) {
    cred := loadCredentials(credFileTest)
    client := loginToForum(cred)
    result := getForumKey(cred, client)
    if len(result) != 5 {
        t.Errorf("Bad forum key: %q", result)
    }
}
