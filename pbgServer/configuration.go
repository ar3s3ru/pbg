package pbgServer

type IConfiguration interface {
    GetHTTPPort()    int
    IsLocalDevelop() bool
}
