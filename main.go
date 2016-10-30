package main

import (
    "fmt"
)

func main() {
    cred := loadCredentials(credFile)

    _, userData := getUserData(cred)
    fetched := getMiiversePosts(userData.MiiverseNames)
    fmt.Printf("Fetched %v posts for %v users\n", len(fetched), len(userData.MiiverseNames))

    cache := loadCache(cacheFile)
    newPosts := filterNewPosts(cache, fetched)

    if len(newPosts) > 0 {
        fmt.Println("New posts:", newPosts)
        newPosts = downloadImages(newPosts)
        client := loginToForum(cred)
        newPosts = uploadImages(cred, client, newPosts)
        forumKey := getForumKey(cred, client)
        forumMessage := formatPosts(cred, userData, newPosts)
        postToForum(cred, client, forumKey, forumMessage)
    } else {
        fmt.Println("No new posts found")
    }

    updateCache(cache, fetched)
    saveCache(cacheFile, cache)
}
