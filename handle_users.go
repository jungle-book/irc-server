package main

func handleUserCommand(client *Client, _ IRCMessage) {
	reply := IRCMessage{
		Command:  RPL_USERSSATRT,
		Params:   []string{},
		Trailing: "",
		Source:   "",
	}

	client.ch <- reply.Encode()
}
