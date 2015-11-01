package main

import (
    "fmt"
)

const cacheFile string = "cache.json"

func load_cache() []MiiversePost {
    var postCache []MiiversePost
    readJSONFile(cacheFile, &postCache)
    fmt.Println(postCache)
    return postCache
}

func save_cache(toSave []MiiversePost){
    writeJSONFile(cacheFile, toSave)
}

func filter_updates(raw []MiiversePost) []MiiversePost {
    var out []MiiversePost

    return out
}

