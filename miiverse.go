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
    is_mario := false
    is_container := false

    for {
        tt := z.Next()

        switch tt {
        case html.ErrorToken:
            // End of the document, we're done
            chPosts <- posts
            return
        case html.StartTagToken:
            t := z.Token()
            if t.Data == "p" {
                ok, class := getAttr(t, "class")
                if ok && class == "post-content-text" {
                    found_it = true
                }
            }
            if t.Data == "a" {
                ok, class := getAttr(t, "class")
                if ok && class == "test-community-link" {
                    is_container = true
                }
            }
        case html.TextToken:
            if is_container {
                inner_text := string(z.Text())
                check_text := "Super Mario Maker Community"
                is_mario = strings.Contains(inner_text, check_text)
                is_container = false
            }
            if found_it {
                found_it = false
                if is_mario {
                    new_post := parsePost(miiverseName, string(z.Text()))
                    posts = append(posts, new_post)
                }
            }
        }
    }
}

func getMiiversePosts(names []string) []MiiversePost {
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