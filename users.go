package main

import (
    "log"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

func crawl_googLe_app(url string) (ok bool, out UserData) {
    resp, err := http.Get(url)
    if err != nil {
        log.Fatal("ERROR: Failed to crawl \"" + url + "\"")
        return 
    }

    body := resp.Body
    defer body.Close() // close Body when the function returns

    raw_json, err := ioutil.ReadAll(body)
    if err != nil {
        log.Fatal("ERROR: Failed to crawl \"" + url + "\"")
        return 
    }

    json.Unmarshal(raw_json, &out)
    ok = true
    return
}

func getUserData(cred Credentials) (ok bool, out UserData) {
    return crawl_googLe_app(cred.GoogleAppUrl)
}