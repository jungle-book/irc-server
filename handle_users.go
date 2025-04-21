package main

func handleNickCommand(client *Client, msg IRCMessage) {
	client.nick = msg.Params[0]

	reply := IRCMessage{
		Source:   "",
		Command:  "NICK",
		Params:   msg.Params,
		Trailing: msg.Trailing,
	}

	mutex.Lock()
	_, ok := clients[client.nick]
	if !ok {
		clients[client.nick] = client
	}
	mutex.Unlock()

	client.ch <- reply.Encode()
}

func handleUserCommand(client *Client, msg IRCMessage) {
	client.username = msg.Params[0]
	client.ch <- "001"
}
