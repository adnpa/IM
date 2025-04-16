package conversation

import (
	"context"

	"github.com/adnpa/IM/api/pb"
)

type ConversationService struct {
	pb.UnimplementedConversationServer
}

// s *ConversationService pb.ConversationServer
func (s *ConversationService) CreateConversation(_ context.Context, _ *pb.CreateConversationReq) (*pb.CreateConversationResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *ConversationService) GetSelfConversationList(_ context.Context, _ *pb.GetSelfConversationListListReq) (*pb.GetSelfConversationListListResp, error) {
	panic("not implemented") // TODO: Implement
}
