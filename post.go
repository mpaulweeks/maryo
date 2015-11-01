package main

import (
    "fmt"
    "code.google.com/p/go.net/publicsuffix"
    "io/ioutil"
    "log"
    "net/http"
    "net/http/cookiejar"
    "net/url"
)

const messageSig string = "\n---\nAdd your MiiverseName here: " + namesUrl

func formatPosts(newPosts []MiiversePost) string {
    html := "<b>NEW LEVELS</b>"
    for _, post := range newPosts{
        html = fmt.Sprintf("%s\n\n%s\n%s\n%s", html, post.MiiverseName, post.Description, post.Code)
    }
    html = html + messageSig
    return html
}

func postToForum(forumMessage string) {
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

    resp, err = client.PostForm(cred["post_url"], url.Values{
        "topic": {cred["thread_id"]},
        "h": {"2b996"},
        "message": {forumMessage},
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
