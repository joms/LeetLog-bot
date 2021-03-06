package main

import(
	"irc"
	"strings"

	_ "commands/top"
)

// Set up and run our IRC connection
func main() {
	irc.Run(&irc.Config{
		Admins:		strings.Split("JoMs", ","),
		Server:		"localhost:6667",
		Channels: 	strings.Split("#sandkas.se", ","),
		User:		"gLeet",
		Nick:		"gLeet",
	})
}