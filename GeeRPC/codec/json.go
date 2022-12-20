package codec

import (
	"encoding/json"
	"io"
)

type JsonCodec struct {
	conn io.ReadWriteCloser
}

type User struct {
	Name     string `json: "name"`
	Password string `json: "password"`
}

func (j JsonCodec) Close() error {
	err := j.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func (j JsonCodec) ReadHeader(header *Header) error {
	panic("not me")
}

func (j *JsonCodec) ReadBody(body interface{}) error {
	//TODO implement me
	var data [1024]byte
	n, err2 := j.conn.Read(data[0:])
	if err2 != nil {
		println("read error: ", err2)
	}
	println("<== server receive: " + string(data[0:n]))
	err := json.Unmarshal(data[0:n], body)
	if err != nil {
		return err
	}
	return nil
}

func (j *JsonCodec) Write(header *Header, body interface{}) error {
	// 序列化
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}
	j.conn.Write(data)
	println("write json str: " + string(data))
	return nil
}

func NewJsonCodec(conn io.ReadWriteCloser) Codec {
	return &JsonCodec{
		conn: conn,
	}
}
