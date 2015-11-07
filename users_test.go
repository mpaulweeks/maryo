package main

import (
    "testing"
)

func TestGoogleAppMiiverseNames(t *testing.T) {
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

func TestGoogleAppNickNames(t *testing.T) {
    cred := loadCredentials(credFileTest)
    ok, res := getUserData(cred)
    if !ok {
        t.Errorf("Error on json fetch")
    }
    expectedKey := "Waff"
    expectedValue := "Mr Lucky Waffles"
    if res.NickNames[expectedKey] != expectedValue {
        t.Errorf("Expected %q: %q, got %q", expectedKey,expectedValue, res)
    }
}
