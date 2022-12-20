package myrpc

import (
	"go-grpc/client"
	"testing"
)

func TestClientSendData(t *testing.T) {

	client.Call("service.hello", "Say", nil)
}
