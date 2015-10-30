package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "golang.org/x/net/html"
)

func query(url string) (body string){
    res, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    body, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s", body)
    return body
}

func parse_notepad(n *html.Node) {
    if n.Type == html.ElementNode && n.Data == "a" {
        for _, a := range n.Attr {
            if a.Key == "href" {
                fmt.Println(a.Val)
                break
            }
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        parse_notepad(c)
    }
}

func get_names() {
    names_url := "http://notepad.cc/eti-mm-draft"
    notepad_html := query(names_url)
    doc, _ := html.Parse(notepad_html)
    parse_notepad(doc)
}

func main() {
    get_names()
    // for _, v := range names {
    //     var url = "https://miiverse.nintendo.net/users/" + v + "/posts"
    //     query(url)
    // }
}