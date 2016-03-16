package ipc

import (
	"encoding/json"
)

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	ch := server.Connect()
	return &IpcClient{ch}
}

func (client *IpcClient) Call(method, params string) (rep *Response, err error) {
	req := &Request{method, params}
	b, err := json.Marshal(req)
	if err != nil {
		return
	}
	client.conn <- string(b)
	str := <-client.conn
	var response Response
	err = json.Unmarshal([]byte(str), &response)
	return &response, err
}

func (client *IpcClient) Close() {
	client.conn <- "CLOSE"
}
