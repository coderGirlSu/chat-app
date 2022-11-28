package conversation

import (
	context "context"
)

// conversationServer represents the chat server
type ConversationServer struct {
	UnimplementedChatServer
}

// SendMessage implements the SendMessage gRPC function
func (cs *ConversationServer) SendMessage(ctx context.Context, in *SendRequest) (*SendReply, error) {
	return nil, nil
}
