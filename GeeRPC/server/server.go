package server

import (
	"GeeRPC/rpc"
	"GeeRPC/service"
	"GeeRPC/util"
	"encoding/json"
	"fmt"
	"net"
)

func Start() {

	// 暴露服务
	rpc.RegisterService("service.hello", service.Hello{})

	// 监听客户端连接
	listen, err := net.Listen("tcp", ":6666")
	if err != nil {
		println("network error !")
	}

	for {
		// 获取连接
		conn, err := listen.Accept()
		if err != nil {
			println("connect error, ", err)
		}

		// 读取套接字
		data := read(&conn)

		// 解码
		request := decode(data)

		// 执行
		invoke(request)
	}

}

func read(conn *net.Conn) []byte {
	var data [1024]byte
	n, err := (*conn).Read(data[0:])
	if err != nil {
		fmt.Println("read data error, ")
	}
	return data[0:n]
}

func decode(data []byte) *rpc.Request {
	request := &rpc.Request{}
	err := json.Unmarshal(data, request)
	if err != nil {
		fmt.Println("Unmarshal error", err)
	}
	return request
}

func invoke(request *rpc.Request) {
	serviceName := request.ServiceName
	service := rpc.FindService(serviceName)
	m := util.MethodByName(service, request.MethodName)
	// 非空判断
	m.Call(nil)
}
