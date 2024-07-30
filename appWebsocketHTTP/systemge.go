package appWebsocketHTTP

import (
	"SystemgeOneNodePing/topics"

	"github.com/neutralusername/Systemge/Message"
	"github.com/neutralusername/Systemge/Node"
)

func (app *AppWebsocketHTTP) GetAsyncMessageHandlers() map[string]Node.AsyncMessageHandler {
	return map[string]Node.AsyncMessageHandler{
		topics.PING: func(node *Node.Node, message *Message.Message) error {
			println("PING")
			node.AsyncMessage(topics.PONG, "")
			return nil
		},
		topics.PONG: func(node *Node.Node, message *Message.Message) error {
			println("PONG")
			return nil
		},
	}
}

func (app *AppWebsocketHTTP) GetSyncMessageHandlers() map[string]Node.SyncMessageHandler {
	return map[string]Node.SyncMessageHandler{}
}
