package main

import (
    "fmt"
    "strings"
    "net/http"
    "golang.org/x/net/html"
)

type MiiversePost struct {
    Username string
    Description string
    Code string
}

func crawl_miiverse(url string, chFinished chan bool, out chan MiiversePost) {
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
            return
        case html.StartTagToken:
            t := z.Token()

            if t.Data != "textarea" {
                continue
            }

            ok, id := getAttr(t, "id")
            if ok && id == "contents" {
                found_it = true
            }
        case html.TextToken:
            if found_it{
                post := string(z.Text())
                posts = append(posts, post)
            }
        }
    }

    out <- posts
    return
}

func get_miiverse(names []string) []MiiversePost {
    var allPosts []MiiversePost

    // Channels
    chPosts := make(chan MiiversePost)
    chFinished := make(chan bool) 

    // Kick off the crawl process (concurrently)
    for _, url := range names {
        go crawl_miiverse(url, chPosts, chFinished)
    }

    // Subscribe to both channels
    for c := 0; c < len(names); {
        select {
        case posts := <-chPosts:
            allPosts = append(allPosts, posts)
        case <-chFinished:
            c++
        }
    }

    return allPosts
}