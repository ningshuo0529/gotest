package cg

import "fmt"

type Player struct {
	Name string
	mq   chan *Message
}

func NewPlayer() *Player {
	m := make(chan *Message, 0)
	player := &Player{"", m}

	go func(p *Player) {
		for {
			msg := <-p.mq
			fmt.Println(p.Name, "receive msg", msg.Content)
		}
	}(player)
	return player
}
