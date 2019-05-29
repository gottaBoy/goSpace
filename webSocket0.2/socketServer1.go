package main

import (
    "bufio"
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
)

func main() {
    //http请求处理
    http.HandleFunc("/", handler1)
    //绑定监听地址和端口
    http.ListenAndServe("localhost:8080", nil)
}

//请求处理函数
func handler1(w http.ResponseWriter, r *http.Request) {
    //获取请求资源
    path := r.URL.Path
    if strings.Contains(path[1:], "") {
        //返回请求资源
        fmt.Fprintf(w, getHtmlFile("index.html"))
    } else {
        if strings.Contains(path[1:], ".html") {
            w.Header().Set("content-type", "text/html")
            fmt.Fprintf(w, getHtmlFile(path[1:]))
        }
        if strings.Contains(path[1:], ".css") {
            w.Header().Set("content-type", "text/css")
            fmt.Fprintf(w, getHtmlFile(path[1:]))
        }
        if strings.Contains(path[1:], ".js") {
            w.Header().Set("content-type", "text/javascript")
            fmt.Fprintf(w, getHtmlFile(path[1:]))
        }
        if strings.Contains(path[1:], "") {
            fmt.Print(strings.Contains(path[1:], ""))
        }
    }

}

func getHtmlFile(path string) (fileHtml string) {
    file, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    rd := bufio.NewReader(file)
    for {
        line, err := rd.ReadString('\n')

        if err != nil || io.EOF == err {
            break
        }
        fileHtml += line
    }
    return fileHtml
}