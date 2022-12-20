package client

import (
	myrpc "GeeRPC"
	"GeeRPC/codec"
	"GeeRPC/rpc"
	"encoding/json"
	"fmt"
	"net"
	"sync"
)

// Client RPC 客户端
type Client struct {
	pending map[uint64]*Call
	mu      sync.Mutex
	codec   codec.Codec
}

// Call 定义一个承载一次 RPC 调用的信息
type Call struct {
	Seq           uint64
	ServiceMethod string
	Args          interface{}
	reply         interface{}
}

type Option struct {
}

func Dial(addr string, option myrpc.Option) *Client {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("拨号失败, 哈哈!")
	}
	return NewClient(conn, option)
}

func NewClient(conn net.Conn, option myrpc.Option) *Client {
	// 给服务端发送消息， 沟通一下协议格式
	codecFunc := codec.NewCodecFuncMap[option.CodeType]
	c := codecFunc(conn)
	c.Write(nil, nil)
	// 创建 client
	client := &Client{codec: c}
	return client
}

func call(serviceName string, methodName string, reply interface{}) error {
	// 连接
	conn, err := net.Dial("tcp", "localhost:6666")
	if err != nil {
		return err
	}
	body := rpc.Request{ServiceName: serviceName, MethodName: methodName}
	// 写入数据
	data, err := json.Marshal(body)
	conn.Write(data)
	return nil
}
