package myrpc

import (
	"go-grpc/codec"
	"net"
	"testing"
	"time"
)

func startServer(addr chan string) {
	listen, err := net.Listen("tcp", ":0")
	if err != nil {
		println("network error !")
	}

	addr <- listen.Addr().String()
	Accept(listen)
}

type User struct {
	Name     string `json: "name"`
	Password string `json: "password"`
}

func TestOne2One(t *testing.T) {
	// start server
	addr := make(chan string)
	go startServer(addr)

	conn, _ := net.Dial("tcp", <-addr)
	defer func() { _ = conn.Close() }()

	time.Sleep(time.Second)
	println("==> client write data start")
	jsonCodec := codec.NewJsonCodec(conn)
	// client write data
	jsonCodec.Write(nil, &User{Name: "zhangfei", Password: "123456"})
	println("==> client write data done")

}
