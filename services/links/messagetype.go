package links

import (
	"net/url"
	"strings"
)

type MessageType string

const (
	Command MessageType = "Command"
	Link                = "Link"
	Message             = "Message"
)

func getMessageType(msg string) MessageType {
	isCmd := strings.HasPrefix(msg, "/cmd")
	if isCmd {
		return Command
	}

	_, err := url.ParseRequestURI(msg)
	if err != nil {
		return Message
	}

	return Link
}
