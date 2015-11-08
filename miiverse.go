package main

import (
    "io"
    "net/url"
    "os"
    "fmt"
    "strings"
    "net/http"
    "golang.org/x/net/html"
)

func parsePost(miiverseName string, nickname string, text string, postUrl string, imgSrc string) MiiversePost {
    splitText := strings.Split(text, "(")
    postId := strings.Trim(postUrl, "/posts/")
    return MiiversePost{
        MiiverseName: miiverseName,
        NickName: nickname,
        Description: strings.Trim(splitText[0], " "),
        Code: strings.Trim(splitText[1], ")"),
        PostId: postId,
        ImgSrc: imgSrc,
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
    is_nickname := false
    nickname := ""
    is_post := false
    postUrl := ""
    postImgSrc := ""

    for {
        tt := z.Next()

        switch tt {
        case html.ErrorToken:
            // End of the document, we're done
            chPosts <- posts
            return
        case html.StartTagToken:
            t := z.Token()
            switch t.Data {
            case "p":
                ok, class := getAttr(t, "class")
                if ok && class == "post-content-text" {
                    found_it = true
                }
            case "a":
                ok, class := getAttr(t, "class")
                if ok && class == "test-community-link" {
                    is_container = true
                }
                if ok && class == "nick-name" {
                    is_nickname = true
                }
                if ok && class == "screenshot-container still-image" {
                    _, postUrl = getAttr(t, "href")
                    is_post = true
                }
            case "img":
                ok, imgSrc := getAttr(t, "src")
                if ok && is_post {
                    postImgSrc = imgSrc
                    is_post = false
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
                    new_post := parsePost(
                        miiverseName,
                        nickname,
                        string(z.Text()),
                        postUrl,
                        postImgSrc,
                    )
                    posts = append(posts, new_post)
                }
            }
            if is_nickname && len(nickname) == 0 {
                nickname = string(z.Text())
                is_nickname = false
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

func downloadImage(rawUrl string) string {
    fmt.Println("Downloading file...")

    fileURL, err := url.Parse(rawUrl)
    if err != nil {
         panic(err)
    }

    path := fileURL.Path
    segments := strings.Split(path, "/")
    fileName := segments[2]
    fileName = fmt.Sprintf("img/%s.jpeg", fileName)

    file, err := os.Create(fileName)
    if err != nil {
         fmt.Println(err)
         panic(err)
    }
    defer file.Close()

    check := http.Client{
         CheckRedirect: func(r *http.Request, via []*http.Request) error {
                 r.URL.Opaque = r.URL.Path
                 return nil
         },
    }

    resp, err := check.Get(rawUrl) // add a filter to check redirect
    if err != nil {
         fmt.Println(err)
         panic(err)
    }
    defer resp.Body.Close()

    fmt.Println(resp.Status)

    size, err := io.Copy(file, resp.Body)
    if err != nil {
         panic(err)
    }

    fmt.Printf("%s with %v bytes downloaded", fileName, size)

    return fileName
}