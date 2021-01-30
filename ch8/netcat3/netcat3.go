package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:32005")
	if err != nil {
		log.Fatal(err)
	}

	// 此处设计成无缓冲通道，发送操作会阻塞，是顺序发送
	// 导致发送和接收goroutine同步化
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} //指示主 goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // 等待后台 goroutine 完成
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
