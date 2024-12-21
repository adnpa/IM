package mongodb

import (
	"github.com/adnpa/IM/pkg/pb/pb_chat"
	"github.com/adnpa/IM/pkg/pb/pb_ws"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/proto"
	"strconv"
)

const (
	singleGocMsgNum = 5000
)

type MsgInfo struct {
	SendTime int64
	Msg      []byte
}

type UserChat struct {
	UID string
	Msg []MsgInfo
}

func SaveUserChat(uid string, sendTime int64, m *pb_chat.MsgDataToDB) error {
	seqUid := getSeqUid(uid, m.MsgData.Seq)
	filter := bson.M{"uid": seqUid}
	mBytes, err := proto.Marshal(m)
	if err != nil {
		return err
	}
	sMsg := MsgInfo{
		SendTime: sendTime,
		Msg:      mBytes,
	}
	err = conn.FindOneAndUpdate(filter, bson.M{"$push": bson.M{"msg": sMsg}})
	if err != nil {
		sChat := UserChat{
			UID: seqUid,
			Msg: []MsgInfo{sMsg},
		}
		_, err := conn.InsertOne(sChat)
		if err != nil {
			return err
		}
	}
	return nil
}

// singleGocMsgNum条一组，select时一个seq会查到5000个文档，组内有序
func getSeqUid(uid string, seq uint32) string {
	seqSuffix := seq / singleGocMsgNum
	return indexGen(uid, seqSuffix)
}

func indexGen(uid string, seqSuffix uint32) string {
	return uid + ":" + strconv.FormatInt(int64(seqSuffix), 10)
}

// GetMsgBySeqList 获取写扩散消息
func GetMsgBySeqList(uid string, seqList []uint32) ([]*pb_ws.MsgData, error) {
	var msgL []*pb_ws.MsgData

	//去重
	// 所有组seq-seqL 的map
	//m := func(uid string, seqList []uint32) map[string][]uint32 {
	//	t := make(map[string][]uint32)
	//	for i := 0; i < len(seqList); i++ {
	//		seqUid := getSeqUid(uid, seqList[i])
	//		if value, ok := t[seqUid]; !ok {
	//			var temp []uint32
	//			t[seqUid] = append(temp, seqList[i])
	//		} else {
	//			t[seqUid] = append(value, seqList[i])
	//		}
	//	}
	//	return t
	//}(uid, seqList)
	//
	//sChat := UserChat{}
	//for seqUid, value := range m {
	//	if err := conn.FindOne(bson.M{"uid": seqUid}).Decode(&sChat); err != nil { //过期/丢失
	//		continue
	//	}
	//
	//}

	return msgL, nil
}
