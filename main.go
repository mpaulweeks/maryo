package main

import (
    "fmt"
)

func main() {
    fmt.Println("main")
    names := getMiiverseNames()
    fetched := getMiiversePosts(names)

    cache := loadCache(cacheFile)
    newPosts := filterNewPosts(cache, fetched)

    if len(newPosts) > 0 {
        forumMessage := formatPosts(newPosts)
        fmt.Println("New message: %q:", forumMessage)
        postToForum(forumMessage)
    }
    fmt.Println("New posts: %q:", newPosts)

    updateCache(cache, fetched)
    saveCache(cacheFile, cache)
}
