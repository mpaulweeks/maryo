package main

import (
    "strings"
    "bytes"
    "io"
    "mime/multipart"
    "os"
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
const bookmarkBase string = "https://supermariomakerbookmark.nintendo.net/courses/"

func formatPost(userData UserData, post MiiversePost) string {
    displayName := post.MiiverseName + " aka " + post.NickName
    forumName, ok := userData.NickNames[post.NickName]
    if ok {
        displayName = forumName
    }
    return fmt.Sprintf(
        "<img src=\"%s\"/>\n%s\n<b>%s</b>\n%s\n<b>Miiverse:</b> %s%s\n<b>Bookmark:</b> %s%s",
        post.ImgUrl, displayName, post.Description, post.Code,
        miiversePostBase, post.PostId,
        bookmarkBase, post.Code,
    )
}

func formatPosts(cred Credentials, userData UserData, newPosts []MiiversePost) string {
    var post_htmls []string
    for _, post := range newPosts{
        post_htmls = append(post_htmls, formatPost(userData, post))
    }
    return strings.Join(post_htmls, "\n\n") + messageSigBase + cred.RegisterUrl
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

func parseImageUploadResponse(resp *http.Response) (ok bool, imgUrl string){
    body := resp.Body
    defer body.Close() // close Body when the function returns

    rawOut, err := ioutil.ReadAll(body)
    if err != nil {
        log.Fatal("ERROR: Failed to read response")
        return
    }
    strOut := string(rawOut)
    imgUrl = strings.Split(strOut, "&quot;")[1]
    ok = true
    return
}

func uploadImage(cred Credentials, client http.Client, file string) (ok bool, imgUrl string) {
    url := cred.ImageUploadUrl

    // Prepare a form that you will submit to that URL.
    var b bytes.Buffer
    w := multipart.NewWriter(&b)
    // Add your image file
    f, err := os.Open(file)
    if err != nil {
        fmt.Print("err opening file")
        return
    }
    fw, err := w.CreateFormFile("file", file)
    if err != nil {
        fmt.Print("err creating form")
        return
    }
    if _, err = io.Copy(fw, f); err != nil {
        fmt.Print("err copying form")
        return
    }
    // Don't forget to close the multipart writer.
    // If you don't close it, your request will be missing the terminating boundary.
    w.Close()

    // Now that you have a form, you can submit it to your handler.
    req, err := http.NewRequest("POST", url, &b)
    if err != nil {
        fmt.Print("err creating request file")
        return
    }
    // Don't forget to set the content type, this will contain the boundary.
    req.Header.Set("Content-Type", w.FormDataContentType())

    // Submit the request
    res, err := client.Do(req)
    if err != nil {
        fmt.Print("err submiting request file")
        return
    }

    // Check the response
    if res.StatusCode != http.StatusOK {
        err = fmt.Errorf("bad status: %s", res.Status)
    }

    ok, imgUrl = parseImageUploadResponse(res)
    return
}

func uploadImages(cred Credentials, client http.Client, miiversePosts []MiiversePost) []MiiversePost {
    var out []MiiversePost
    for _, post := range miiversePosts {
        _, imgUrl := uploadImage(cred, client, post.ImgFile)
        post.ImgUrl = imgUrl
        out = append(out, post)
    }
    return out
}
