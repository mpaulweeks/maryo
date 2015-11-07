package main

import (
    "fmt"
)

func main() {
    cred := loadCredentials(credFile)

    _, userData := getUserData(cred)
    fetched := getMiiversePosts(userData.MiiverseNames)

    cache := loadCache(cacheFile)
    newPosts := filterNewPosts(cache, fetched)

    if len(newPosts) > 0 {
        fmt.Println("New posts:", newPosts)
        forumMessage := formatPosts(cred, userData, newPosts)
        client := loginToForum(cred)
        forumKey := getForumKey(cred, client)
        postToForum(cred, client, forumKey, forumMessage)
    }

    updateCache(cache, fetched)
    saveCache(cacheFile, cache)
}
