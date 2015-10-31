package main

import (
    "golang.org/x/net/html"
)

func getAttr(t html.Token, attr string) (ok bool, val string) {
    for _, a := range t.Attr {
        if a.Key == attr {
            val = a.Val
            ok = true
        }
    }
    return
}

func contains(slice []string, item string) bool {
    set := make(map[string]struct{}, len(slice))
    for _, s := range slice {
        set[s] = struct{}{}
    }

    _, ok := set[item]
    return ok
}
