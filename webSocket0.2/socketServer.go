package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)
//https://github.com/Jiashengp/GoHttpServer
// https://github.com/Jiashengp/Python_httpServer
func main() {
	//确定协议及绑定IP及端口
	netListen, err := net.Listen("tcp", "localhost:8080")
	CheckError(err)
	defer netListen.Close()
	Log("waiting for client request")
	for {
		//接受请求连接
		conn, err := netListen.Accept()
		if err != nil {
			CheckError(err)
			break
		} else {
			Log(conn.RemoteAddr().String(), "tcp connect success")
			//处理请求连接
			handleConnection(conn)
		}
		conn.Close()
	}
}

//处理请求连接函数
func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)
	if err != nil {
		Log(conn.RemoteAddr().String(), " connection error: ", err)
		conn.Close()
	} else {
		Log(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
		responseInfoToClient(conn, string(buffer[:n]), err)
	}
}

//返回数据的函数
func responseInfoToClient(conn net.Conn, requestInfo string, err error) {
	//获取http协议头
	conn.Write([]byte(getFileContent("head.md")))
	conn.Write([]byte("\n"))
	var path string = strings.Replace(getMidStr(requestInfo, "GET /", "HTTP"), " ", "", -1)
	fmt.Println(path)
	if path != "" {
		if path == "favicon.ico" {
			fmt.Println("every connect hava favicon.ico resource request")
		} else {
			_, err = os.Open(path)
			if err != nil {
				fmt.Println("RESTful")
			} else {
				conn.Write([]byte(getFileContent(path)))
			}
		}
	} else {
		conn.Write([]byte(getFileContent("index.html")))
	}
}

func getMidStr(data string, startStr string, endStr string) (reqSouce string) {
	var startIndex int = strings.Index(data, startStr)
	var info string
	if startIndex >= 0 {
		startIndex += len(startStr)
		var endIndex int = strings.Index(data, endStr)
		info = data[startIndex:endIndex]
	}
	return info
}

func getFileContent(path string) (fileInfo string) {
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
		fileInfo += line
	}
	return fileInfo
}

func Log(v ...interface{}) {
	log.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
