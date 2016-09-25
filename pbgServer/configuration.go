package pbgServer

type IConfiguration interface {
    GetListenAndServe() string
    GetHTTPPort()       int
}
