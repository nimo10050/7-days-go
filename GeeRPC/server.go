package myrpc

import (
	"GeeRPC/codec"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
)

const MagicNumber = 0x36789

type Option struct {
	MagicNumber int
	CodeType    codec.Type
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodeType:    codec.GobType,
}

type request struct {
	h            codec.Header
	argv, replyv reflect.Value
}

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

func (server *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			println("rpc server accept error, msg: ", err)
			return
		}
		go server.connection(conn)
	}
}

func Accept(lis net.Listener) {
	DefaultServer.Accept(lis)
}

func (server *Server) connection(conn io.ReadWriteCloser) {
	defer func() { _ = conn.Close() }()
	var opt Option
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		return
	}
	if opt.MagicNumber != MagicNumber {
		return
	}

	codecFunc := codec.NewCodecFuncMap[codec.JsonType]
	c := codecFunc(conn)

	// read header
	var h codec.Header
	if err := c.ReadHeader(&h); err != nil {
		return
	}

	// read body
	req := &request{h: h}
	req.argv = reflect.New(reflect.TypeOf(""))

	if err := c.ReadBody(req.argv.Interface()); err != nil {
		log.Println("rpc server: read argv err:", err)
	}

	// handle request
	fmt.Println("receive : ", req.h, req.argv)
}

//func serverCodec(cc codec.Codec) {
//	for {
//		cc.
//	}
//}
