package main

import (
)

func load_cache(filePath string) []MiiversePost {
    var postCache []MiiversePost
    readJSONFile(filePath, &postCache)
    return postCache
}

func save_cache(filePath string, toSave []MiiversePost){
    writeJSONFile(filePath, toSave)
}

func filter_updates(raw []MiiversePost) []MiiversePost {
    var out []MiiversePost

    return out
}

