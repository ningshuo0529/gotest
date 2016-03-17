package cg

import (
	"ipc"
	"testing"
)

func TestIpc(t *testing.T) {
	server := ipc.NewIpcServer(&CenterServer{})
	client := ipc.NewIpcClient(server)
	centerclient := CenterClient{client}
	player := NewPlayer()
	player.Name = "ns"
	err := centerclient.AddPlayer(player)
	player.Name = "shit"
	err = centerclient.AddPlayer(player)
	if err != nil {
		t.Error(err)
	}
	ps, err := centerclient.ListPlayer("hehe")
	if err != nil {
		t.Error(err)
	}
	for _, p := range ps {
		t.Log(*p)
	}
	centerclient.Broadcast("hehehe")
	err = centerclient.RemovePlayer("ns")
	if err != nil {
		t.Error(err)
	}
	client.Close()
}
