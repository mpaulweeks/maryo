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

    // postToForum()
    fmt.Println(newPosts)

    updateCache(cache, fetched)
    saveCache(cacheFile, cache)
}
