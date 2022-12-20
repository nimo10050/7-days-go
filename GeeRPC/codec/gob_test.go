package codec

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

var buf bytes.Buffer

func TestEncodeAndDecode(t *testing.T) {
	//file, err := os.Open("/Users/zhanggaopei/Downloads/123.txt")
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(Person{Name: "zhangsan", Age: 10})
	if err != nil {
		println(err.Error())
	}
	//println(buf.String())

	decoder := gob.NewDecoder(&buf)
	var p1 = &Person{}
	decoder.Decode(p1)
	println("p1.name: ", p1.Name)
}

func TestAAA(t *testing.T) {
	p := Person{Name: "justin", Age: 30}
	buf := Encoder(p)

	p1 := Person{}
	err := Decoder(buf, p1)
	if err != nil {
		fmt.Println(err)
	} else {
		println("p1.name: ", p1.Name)
	}
}

func Encoder(inter interface{}) bytes.Buffer {
	var buf bytes.Buffer
	// writer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(inter)
	if err != nil {
		fmt.Println(err)
	}
	return buf
}

func Decoder(buf bytes.Buffer, inter interface{}) error {
	// reader
	decoder := gob.NewDecoder(&buf)
	decoder.Decode(inter)
	println(inter)
	return nil
}
