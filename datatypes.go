package main

type MiiversePost struct {
    MiiverseName string
    NickName string
    Description string
    Code string
    PostId string
    ImgSrc string
}

type Credentials struct {
    RegisterUrl string
    GoogleAppUrl string
    Login string
    PostUrl string
    GetUrl string
    ThreadId string
    Username string
    Password string
}

type UserData struct {
    MiiverseNames []string
    NickNames map[string]string
}
