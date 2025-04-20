package main

func handleNickCommand(client *Client, msg IRCMessage) {
	client.nick = msg.Params[0]

	client.ch <- "001"
}

func handleUserCommand(client *Client, msg IRCMessage) {
	client.username = msg.Params[0]

	reply := IRCMessage{
		Command:  RPL_USERSSATRT,
		Params:   []string{},
		Trailing: "",
		Source:   "",
	}

	client.ch <- reply.Encode()
}
