package main

type MiiversePost struct {
    MiiverseName string
    NickName string
    Description string
    Code string
    PostId string
    ImgSrc string
    ImgFile string
    ImgUrl string
}

type Credentials struct {
    RegisterUrl string
    GoogleAppUrl string
    Login string
    PostUrl string
    GetUrl string
    ImageUploadUrl string
    ThreadId string
    Username string
    Password string
}

type UserData struct {
    MiiverseNames []string
    NickNames map[string]string
}
