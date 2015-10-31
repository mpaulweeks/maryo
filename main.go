package main

import (
    "fmt"
)

func main() {
    names := get_names()
    fmt.Println(names)
    // for _, v := range names {
    //     var url = "https://miiverse.nintendo.net/users/" + v + "/posts"
    //     query(url)
    // }
}
