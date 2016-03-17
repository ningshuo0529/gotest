package main

import (
	"bufio"
	"cg"
	"fmt"
	"ipc"
	"os"
	"strings"
)

var centerClient *cg.CenterClient

func startCenterService() error {
	server := ipc.NewIpcServer(&cg.CenterServer{})
	client := ipc.NewIpcClient(server)
	centerClient = &cg.CenterClient{client}
	return nil
}

func GetCommandHandlers() map[string]func([]string) int {
	return map[string]func([]string) int{
		"help":        Help,
		"h":           Help,
		"quit":        Quit,
		"q":           Quit,
		"login":       Login,
		"logout":      Logout,
		"listplayers": ListPlayer,
		"send":        Send,
	}
}

func Quit([]string) int {
	return 1
}

func Help([]string) int {
	fmt.Println(`
        Commands:
        login <username>
        logout <username>
        send <message>
        listplayers
        quit(q)
        help(h)
    `)
	return 0
}

func Login(args []string) int {
	if len(args) != 2 {
		fmt.Println("USAGE: login <username>")
	}
	player := cg.NewPlayer()
	player.Name = args[1]
	err := centerClient.AddPlayer(player)
	if err != nil {
		fmt.Println("failed add player", err)
	}
	return 0
}

func Logout(args []string) int {
	if len(args) != 2 {
		fmt.Println("USAGE: logout <username>")
	}
	centerClient.RemovePlayer(args[1])
	return 0
}

func ListPlayer(args []string) int {
	ps, err := centerClient.ListPlayer("")
	if err != nil {
		fmt.Println("failed", err)
	} else {
		for i, v := range ps {
			fmt.Println(i+1, ":", v)
		}
	}
	return 0
}

func Send(args []string) int {
	message := strings.Join(args[1:], " ")
	err := centerClient.Broadcast(message)
	if err != nil {
		fmt.Println("failed", err)
	}
	return 0
}

func main() {
	fmt.Println("Start")
	startCenterService()
	Help(nil)
	r := bufio.NewReader(os.Stdin)
	handlers := GetCommandHandlers()
	for {
		fmt.Print("Command>")
		b, _, _ := r.ReadLine()
		line := string(b)
		tokens := strings.Split(line, " ")
		if handler, ok := handlers[tokens[0]]; ok {
			ret := handler(tokens)
			if ret != 0 {
				break
			}
		} else {
			fmt.Println("command not found", tokens[0])
		}
	}
}
