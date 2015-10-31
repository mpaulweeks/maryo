package main

import (
    "code.google.com/p/go.net/publicsuffix"
    "io/ioutil"
    "encoding/json"
    "log"
    "net/http"
    "net/http/cookiejar"
    "net/url"
)

const credFile string = "cred.json"


func readFile(path string) []byte {
    fileContents, err := ioutil.ReadFile(path)
    if err == nil {
        return fileContents
    }
    panic(err)
}

func readJSONFile(path string, contentsHolder interface{}) {
    var fileContents = readFile(path)
    err := json.Unmarshal(fileContents, contentsHolder)
    if err != nil {
        panic(err)
    }
}

func postToForum() {
    var cred map[string]string
    readJSONFile(credFile, &cred)

    options := cookiejar.Options{
        PublicSuffixList: publicsuffix.List,
    }
    jar, err := cookiejar.New(&options)
    if err != nil {
        log.Fatal(err)
    }
    client := http.Client{Jar: jar}
    resp, err := client.PostForm(cred["login"], url.Values{
        "b": {cred["username"]},
        "p" : {cred["password"]},        
    })
    if err != nil {
        log.Fatal(err)
    }

    // resp, err = client.Get(cred["profile"])
    // if err != nil {
    //     log.Fatal(err)
    // }

    resp, err = client.PostForm(cred["post_url"], url.Values{
        "topic": {cred["thread_id"]},
        "h": {"2b996"},
        "message": {"auto\n---\nauto sig"},
        "-ajaxCounter": {"1"},
    })
    if err != nil {
        log.Fatal(err)
    }

    data, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    log.Println(string(data))   // print whole html of user profile data 
}
