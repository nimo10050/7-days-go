package codec

import (
	"bufio"
	"encoding/gob"
	"io"
)

type GobCodec struct {
	conn io.ReadWriteCloser
	dec  *gob.Decoder
	enc  *gob.Encoder
	buf  *bufio.Writer
}

func (g *GobCodec) Close() error {
	err := g.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func (g *GobCodec) ReadHeader(header *Header) error {
	err := g.dec.Decode(header)
	if err != nil {
		return err
	}
	return nil
}

func (g *GobCodec) ReadBody(body interface{}) error {
	err := g.dec.Decode(body)
	if err != nil {
		return err
	}
	return nil
}

func (g *GobCodec) Write(header *Header, body interface{}) (err error) {
	defer func() {
		_ = g.buf.Flush()
		if err != nil {
			_ = g.Close()
		}
	}()
	if err := g.enc.Encode(header); err != nil {
		return nil
	}

	if err := g.enc.Encode(body); err != nil {
		return err
	}
	return nil
}

func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
		buf:  buf,
	}
}

var _ Codec = (*GobCodec)(nil)
