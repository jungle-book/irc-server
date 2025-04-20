package main

func handleCommand(client *Client, command string) {
	msg := parseIRCMessage(command)
	msg.Log()

	switch msg.Command {
	case "JOIN":
		handleJoinCommand(client, msg)
	case "USER":
		handleUserCommand(client, msg)
	case "NICK":
		handleNickCommand(client, msg)
	case "PING":
		handlePingCommand(client, msg)
	case "PRIVMSG":
		handleNotifyCommand(client, msg)
	}
}
