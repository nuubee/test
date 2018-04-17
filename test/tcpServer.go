package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var host = flag.String("host", "", "host")
var port = flag.String("port", "9999", "port")

type Msg struct {
	Data string `json:"data"`
	Type int    `json:"type"`
}

type Resp struct {
	Data   string `json:"data"`
	Status int    `json:"status"`
}

func main() {
	// 解析参数
	flag.Parse()
	var l net.Listener
	var err error
	// 监听
	l, err = net.Listen("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer l.Close()

	fmt.Println("Listening on " + *host + ":" + *port)

	for {
		// 接收一个client
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}

		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

		// 执行
		go handleRequest(conn)
		//handleRequest(conn)
	}
}

// 处理接收到的connection
//
func handleRequest(conn net.Conn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("Disconnected :" + ipStr)
		conn.Close()
	}()

	// 构建reader和writer
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		// 读取一行数据, 以"\n"结尾
		b, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		// 反序列化数据
		var msg Msg
		json.Unmarshal(b, &msg)
		fmt.Println("GET ==>  data: ", msg.Data, " type: ", msg.Type)

		// 构建回复Msg
		resp := Resp{
			Data:   time.Now().String(),
			Status: 200,
		}
		r, _ := json.Marshal(resp)

		writer.Write(r)
		writer.Write([]byte("\n"))
		writer.Flush()
		//conn.Write(r)
		//conn.Write([]byte("\n"))
	}

	fmt.Println("Done!")
}
