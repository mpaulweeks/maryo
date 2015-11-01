package main

import (
    "fmt"
    "strings"
    "net/http"
    "golang.org/x/net/html"
)

func parsePost(miiverseName string, text string) MiiversePost {
    split := strings.Split(text, "(")
    return MiiversePost{
        MiiverseName: miiverseName,
        Description: strings.Trim(split[0], " "),
        Code: strings.Trim(split[1], ")"),
    }
}

func crawl_miiverse(miiverseName string, chFinished chan bool, chPosts chan []MiiversePost) {
    url := "https://miiverse.nintendo.net/users/" + miiverseName + "/posts"

    resp, err := http.Get(url)

    var posts []MiiversePost

    defer func() {
        // Notify that we're done after this function
        chFinished <- true
    }()

    if err != nil {
        fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
        return 
    }

    b := resp.Body
    defer b.Close() // close Body when the function returns

    z := html.NewTokenizer(b)

    found_it := false

    for {
        tt := z.Next()

        switch tt {
        case html.ErrorToken:
            // End of the document, we're done
            chPosts <- posts
            return
        case html.StartTagToken:
            t := z.Token()

            if t.Data != "p" {
                continue
            }

            ok, id := getAttr(t, "class")
            if ok && id == "post-content-text" {
                found_it = true
            }
        case html.TextToken:
            if found_it{
                new_post := parsePost(miiverseName, string(z.Text()))
                posts = append(posts, new_post)
                found_it = false
            }
        }
    }
}

func get_miiverse(names []string) []MiiversePost {
    var allPosts []MiiversePost

    // Channels
    chPosts := make(chan []MiiversePost)
    chFinished := make(chan bool) 

    // Kick off the crawl process (concurrently)
    for _, miiverseName := range names {
        go crawl_miiverse(miiverseName, chFinished, chPosts)
    }

    // Subscribe to both channels
    for c := 0; c < len(names); {
        select {
        case posts := <-chPosts:
            allPosts = append(allPosts, posts...)
        case <-chFinished:
            c++
        }
    }

    return allPosts
}