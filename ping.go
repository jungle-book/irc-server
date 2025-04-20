package main

func handlePingCommand(client *Client, msg IRCMessage) {
	resp := IRCMessage{
		Command:  "PONG",
		Source:   "",
		Params:   msg.Params,
		Trailing: "",
	}

	client.ch <- resp.Encode()
}
