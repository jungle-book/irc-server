package main

func handleCommand(client *Client, command string) {
	msg := parseIRCMessage(command)
	msg.Log()

	switch msg.Command {
	case "JOIN":
		handleJoinCommand(client, msg)
	case "USER":
		handleUserCommand(client, msg)
	}
}
