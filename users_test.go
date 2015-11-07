package main

import (
    "testing"
)

func TestGoogleApp(t *testing.T) {
    cred := loadCredentials(credFileTest)
    ok, res := getUserData(cred)
    if !ok {
        t.Errorf("Error on json fetch")
    }
    expected := "MrLuckyWaffles"
    if !contains(res.MiiverseNames, expected){
        t.Errorf("Expected %q, got %q", expected, res)
    }
}
