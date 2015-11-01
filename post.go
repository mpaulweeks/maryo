package main

import (
    "fmt"
    "code.google.com/p/go.net/publicsuffix"
    "io/ioutil"
    "log"
    "net/http"
    "net/http/cookiejar"
    "net/url"
    "golang.org/x/net/html"
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

func crawlForumTopic(client http.Client, url string) (success bool, out string) {
    resp, err := client.Get(url)

    if err != nil {
        log.Fatal("ERROR: Failed to crawl \"" + url + "\"")
        return
    }

    b := resp.Body
    defer b.Close() // close Body when the function returns

    z := html.NewTokenizer(b)
    for {
        tt := z.Next()

        switch tt {
        case html.ErrorToken:
            // End of the document, we're done
            return
        case html.SelfClosingTagToken:
            t := z.Token()
            if t.Data == "input" {
                ok, name := getAttr(t, "name")
                if ok && name == "h" {
                    ok, key := getAttr(t, "value")
                    if ok {
                        success = true
                        out = key
                        return
                    }
                }
            }
        }
    }
}

func getForumKey(cred map[string]string, client http.Client) string {
    get_url := cred["get_url"] + cred["thread_id"]
    ok, key := crawlForumTopic(client, get_url)
    if !ok {
        log.Fatal("ERROR: Failed to find forum key. Url:", get_url)
    }
    return key
}

func loadForumCredentials(filePath string) map[string]string {
    var cred map[string]string
    readJSONFile(filePath, &cred)
    return cred
}

func loginToForum(cred map[string]string) http.Client {
    options := cookiejar.Options{
        PublicSuffixList: publicsuffix.List,
    }
    jar, err := cookiejar.New(&options)
    if err != nil {
        log.Fatal(err)
    }
    client := http.Client{Jar: jar}
    _, err = client.PostForm(cred["login"], url.Values{
        "b": {cred["username"]},
        "p" : {cred["password"]},        
    })
    if err != nil {
        log.Fatal(err)
    }
    return client
}

func postToForum(cred map[string]string, client http.Client, forumKey string, forumMessage string) {
    resp, err := client.PostForm(cred["post_url"], url.Values{
        "topic": {cred["thread_id"]},
        "h": {forumKey},
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
    log.Println(string(data))   // print resp
}
