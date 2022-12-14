package chat

import (
	"context"
	"time"

	chat_message "github.com/decentralized-chat/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (c *Client) MessagePeer(msg *chat_message.ContentMessage, peer Peer) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := peer.client.SendMessage(ctx, msg)
	if err != nil {
		c.zk.SendMessageToQueue(c.channel, peer.user.Username, c.User.Username, msg.Content)
		return err
	}

	return nil
}

func (c *Client) BroadcastMessage(content string) error {
	msg := &chat_message.ContentMessage{
		From:    &c.User,
		Content: content,
		SentAt:  timestamppb.Now(),
	}

	for _, peer := range c.peers {
		go c.MessagePeer(msg, peer)
	}

	return nil
}
