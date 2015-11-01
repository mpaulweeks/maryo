package main

import (
    "fmt"
    "strings"
    "net/http"
    "golang.org/x/net/html"
)

func crawl_notepad(url string) (ok bool, out string) {
    resp, err := http.Get(url)

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
                ok = true
                out = string(z.Text())
                return                
            }
        }
    }
}

func getMiiverseNames() []string {
    names_url := "http://notepad.cc/eti-mm"
    var names []string
    ok, notepad := crawl_notepad(names_url)
    if !ok {
        return names
    }
    for _, line := range strings.Split(notepad, "\n"){
        trimmed := strings.Trim(line, " ")
        if len(trimmed) == 0 || trimmed[0:1] == "#" {
            continue
        }
        names = append(names, trimmed)
    }
    return names
}