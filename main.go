package main

import (
    "fmt"
)

func main() {
    names := getMiiverseNames()
    fetched := getMiiversePosts(names)

    cache := loadCache(cacheFile)
    newPosts := filterNewPosts(cache, fetched)

    if len(newPosts) > 0 {
        fmt.Println("New posts:", newPosts)
        forumMessage := formatPosts(newPosts)
        cred := loadCredentials(credFile)
        client := loginToForum(cred)
        forumKey := getForumKey(cred, client)
        postToForum(cred, client, forumKey, forumMessage)
    }

    updateCache(cache, fetched)
    saveCache(cacheFile, cache)
}
