package main

type MiiversePost struct {
    MiiverseName string
    NickName string
    Description string
    Code string
}

type Credentials struct {
    GoogleAppUrl string
    Login string
    PostUrl string
    GetUrl string
    ThreadId string
    Username string
    Password string
}

type GoogleAppResult struct {
    MiiverseNames []string
    NickNames map[string]string
}
