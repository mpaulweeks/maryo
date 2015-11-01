package main

import (
)

func loadCache(filePath string) map[string]MiiversePost {
    var postCache map[string]MiiversePost
    readJSONFile(filePath, &postCache)
    return postCache
}

func saveCache(filePath string, toSave map[string]MiiversePost){
    writeJSONFile(filePath, toSave)
}

func filterNewPosts(cache map[string]MiiversePost, fetched []MiiversePost) []MiiversePost {
    var cachedNames = make(map[string]bool)
    for _, savedPost := range cache {
        cachedNames[savedPost.MiiverseName] = true
    }

    var out []MiiversePost
    for _, newPost := range fetched {
        _, codeIsPresent := cache[newPost.Code]
        _, oldUser := cachedNames[newPost.MiiverseName]
        if !codeIsPresent && oldUser {
            out = append(out, newPost)
        }
    }
    return out
}

func updateCache(cache map[string]MiiversePost, fetched []MiiversePost) {
    for _, newPost := range fetched {
        _, codeIsPresent := cache[newPost.Code]
        if !codeIsPresent {
            cache[newPost.Code] = newPost
        }
    }
}

