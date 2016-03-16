package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server EchoServer) Name() string {
	return "EchoServer"
}

func (server EchoServer) Handle(method, params string) *Response {
	return &Response{"200", method + params}
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)
	resp1, _ := client1.Call("foo", "From client1")
	resp2, _ := client2.Call("foo", "From client2")
	if resp1.Body != "fooFrom client1" ||
		resp1.Code != "200" || resp2.Code != "200" ||
		resp2.Body != "fooFrom client2" {
		t.Error("call failed")
	}
	client1.Close()
	client2.Close()
}
