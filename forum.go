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

const messageSigBase string = "\n---\nRegister your MiiverseName here: "
const miiversePostBase string = "https://miiverse.nintendo.net/posts/"

func formatPost(userData UserData, post MiiversePost) string {
    displayName := post.MiiverseName + " aka " + post.NickName
    forumName, ok := userData.NickNames[post.NickName]
    if ok {
        displayName = forumName
    }
    return fmt.Sprintf("%s\n%s\n%s\n%s%s", displayName, post.Description, post.Code, miiversePostBase, post.PostId)
}

func formatPosts(cred Credentials, userData UserData, newPosts []MiiversePost) string {
    html := "<b>NEW LEVELS</b>"
    for _, post := range newPosts{
        html = fmt.Sprintf("%s\n\n%s", html, formatPost(userData, post))
    }
    html = html + messageSigBase + cred.RegisterUrl
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

func getForumKey(cred Credentials, client http.Client) string {
    get_url := cred.GetUrl + cred.ThreadId
    ok, key := crawlForumTopic(client, get_url)
    if !ok {
        log.Fatal("ERROR: Failed to find forum key. Url:", get_url)
    }
    return key
}

func loginToForum(cred Credentials) http.Client {
    options := cookiejar.Options{
        PublicSuffixList: publicsuffix.List,
    }
    jar, err := cookiejar.New(&options)
    if err != nil {
        log.Fatal(err)
    }
    client := http.Client{Jar: jar}
    _, err = client.PostForm(cred.Login, url.Values{
        "b": {cred.Username},
        "p" : {cred.Password},
    })
    if err != nil {
        log.Fatal(err)
    }
    return client
}

func postToForum(cred Credentials, client http.Client, forumKey string, forumMessage string) {
    resp, err := client.PostForm(cred.PostUrl, url.Values{
        "topic": {cred.ThreadId},
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
