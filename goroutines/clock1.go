package main

import (
	"io"
	"log"
	"net"
	"time"
)

var count int

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		count++
		log.Printf("新连接：%s, 总连接数：%d\n", conn.RemoteAddr(), count)
		go handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer func() {
		c.Close()
		count--
		log.Printf("已离开：%s，总连接数：%d\n", c.RemoteAddr(), count)
	}()
	for {
		_, err := io.WriteString(c, time.Now().Format("\r现在的时间是：15:04:05\n"))
		if err != nil {
			return // e.g.,
		}
		time.Sleep(1 * time.Second)
	}
}
