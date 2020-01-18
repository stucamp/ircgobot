package main

import (
	"fmt"
	"log"
	"net"

	"gopkg.in/irc.v3"
)

func main() {
	conn, err := net.Dial("tcp", "chat.freenode.net:6667")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		log.Fatalln(err)
	}

	hasCreds, stucreds := GetCredsFromJSON()

	var config irc.ClientConfig

	if hasCreds {
		config = irc.ClientConfig{
			Nick: stucreds[0],
			Pass: stucreds[1],
			User: stucreds[2],
			Name: stucreds[3],
			Handler: irc.HandlerFunc(func(c *irc.Client, m *irc.Message) {
				if m.Command == "001" {
					// 001 is a welcome event, so we join channels there
					//c.Write("JOIN #ubuntu-us-az")
					c.Write("JOIN #bot-testing-chan")
				} else if m.Command == "PRIVMSG" && c.FromChannel(m) {
					// Create a handler on all messages.
					c.WriteMessage(&irc.Message{
						Command: "PRIVMSG",
						Params: []string{
							m.Params[0],
							m.Trailing(),
						},
					})
				}
			}),
		}
	}

	// Create the client
	client := irc.NewClient(conn, config)
	err = client.Run()
	if err != nil {
		fmt.Printf("Error: %s", err)
		log.Fatalln(err)
	}
}
