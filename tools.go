package main

import (
    "golang.org/x/net/html"
    "io/ioutil"
    "encoding/json"
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

func readJSONFile(path string, contentsHolder interface{}) {
    fileContents, err := ioutil.ReadFile(path)
    if err != nil {panic(err)}
    err = json.Unmarshal(fileContents, contentsHolder)
    if err != nil {panic(err)}
}

func writeJSONFile(path string, contentsHolder interface{}) {
    fileContents, err := json.Marshal(contentsHolder)
    if err != nil {panic(err)}
    err = ioutil.WriteFile(path, fileContents, 0644)
    if err != nil {panic(err)}
}

func loadCredentials(filePath string) Credentials {
    var cred Credentials
    readJSONFile(filePath, &cred)
    return cred
}
