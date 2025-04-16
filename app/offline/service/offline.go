package service

import (
	"context"
	"time"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/offline/global"
	"github.com/adnpa/IM/app/offline/model"
	imodel "github.com/adnpa/IM/internal/model"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// s *OfflineService pb.OfflineServer

const (
	DB_Name     = "im"
	COLL_NAME   = "offline"
	INBOX_KEY   = "uid"
	INBOX_FIELD = "inbox"
)

type OfflineService struct {
	pb.UnimplementedOfflineServer
}

func (s *OfflineService) GetOfflineMsg(_ context.Context, in *pb.GetOfflineMsgReq) (*pb.GetOfflineMsgResp, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	coll := global.DB.Database(DB_Name).Collection(COLL_NAME)

	filter := bson.M{INBOX_KEY: in.Uid}

	inbox := model.Inbox{}
	err := coll.FindOne(ctx, filter).Decode(&inbox)
	if err != nil {
		return nil, err
	}

	var pbMsgs []*pb.ChatMsg
	for _, msg := range inbox.Inbox {
		pbMsg := pb.ChatMsg{}
		copier.Copy(&pbMsg, msg)
		pbMsgs = append(pbMsgs, &pbMsg)
	}

	return &pb.GetOfflineMsgResp{Msgs: pbMsgs}, nil
}

func (s *OfflineService) PutMsg(ctx context.Context, in *pb.PutMsgReq) (*pb.PutMsgResp, error) {
	var msg imodel.ChatMessage
	copier.Copy(&msg, in.Msg)
	filter := bson.M{INBOX_KEY: in.UserId}
	update := bson.M{"$push": bson.M{INBOX_FIELD: msg}}

	ctx, cancelFunc := context.WithTimeout(ctx, 2*time.Second)
	defer cancelFunc()
	coll := global.DB.Database(DB_Name).Collection(COLL_NAME)

	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	err := coll.FindOneAndUpdate(ctx, filter, update, opts).Err()
	if err != nil {
		return nil, err
	}
	return &pb.PutMsgResp{Succ: true}, nil
}

func (s *OfflineService) RemoveMsg(ctx context.Context, in *pb.RemoveMsgReq) (*pb.RemoveMsgResp, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, 2*time.Second)
	defer cancelFunc()
	coll := global.DB.Database(DB_Name).Collection(COLL_NAME)

	filter := bson.M{INBOX_KEY: in.Uid}
	// update := bson.M{
	// 	"$pull": bson.M{
	// 		"inbox": bson.M{"id": bson.M{"$in": in.MsgIds}},
	// 	},
	// }

	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &pb.RemoveMsgResp{}, nil
}
