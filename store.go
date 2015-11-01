package main

import (
)

func load_cache(filePath string) map[string]MiiversePost {
    var postCache map[string]MiiversePost
    readJSONFile(filePath, &postCache)
    return postCache
}

func save_cache(filePath string, toSave map[string]MiiversePost){
    writeJSONFile(filePath, toSave)
}

func filter_updates(cache map[string]MiiversePost, fetched []MiiversePost) []MiiversePost {
    var cachedNames map[string]bool
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

func update_cache(cache map[string]MiiversePost, fetched []MiiversePost) {
    for _, newPost := range fetched {
        _, codeIsPresent := cache[newPost.Code]
        if !codeIsPresent {
            cache[newPost.Code] = newPost
        }
    }
}

