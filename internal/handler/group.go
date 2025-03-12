package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/adnpa/IM/internal/service/group"
	"github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/common/db/mongodb"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type CreateGroupReq struct {
	Uid  int64  `json:"uid"`
	Name string `json:"name"`
}

func CreateGroup(c *gin.Context) {
	req := &CreateGroupReq{}
	err := c.ShouldBind(req)
	if err != nil {
		return
	}

	uid, _ := strconv.ParseInt(c.Query("uid"), 10, 64)
	srv := group.GroupService{}
	srv.CreateGroup(&group.GroupInfo{
		Gid:   utils.NowMilliSecond(),
		Name:  req.Name,
		Owner: uid,
	})
}

func Disband(c *gin.Context) {

}

func ApplyGroup(c *gin.Context) {
	uid, _ := strconv.ParseInt(c.Query("uid"), 10, 64)
	gid, _ := strconv.ParseInt(c.Query("gid"), 10, 64)
	mongodb.Insert("group_member", &group.GroupMember{
		GroupId: gid,
		UserId:  uid,
		LastAck: 0})
	c.JSON(http.StatusOK, "ok")
}

func GetGroups(c *gin.Context) {
	uid, _ := strconv.ParseInt(c.Query("uid"), 10, 64)
	var groups []*group.GroupMember
	cur, _ := mongodb.GetAll("group_member", bson.M{"uid": uid})
	cur.All(context.Background(), &groups)
	var groupInfo []*group.GroupInfo
	cur1, _ := mongodb.GetAll("group_info", bson.M{})
	cur1.All(context.Background(), &groupInfo)
	c.JSON(http.StatusOK, groupInfo)
}

func GetSelfApplyList(c *gin.Context) {

}

func GetGroupApplyList(c *gin.Context) {

}

type HandleApplyReq struct {
	GroupId int64
	UserId  int64
	Flag    group.ApplyFlag
}

func HandleApply(c *gin.Context) {
	req := &HandleApplyReq{}
	err := c.ShouldBind(req)
	if err != nil {
		return
	}

	// uid, _ := strconv.ParseInt(c.Query("uid"), 10, 64)
	srv := group.GroupService{}
	srv.HandleApply(&group.GroupApply{
		GroupId:     req.GroupId,
		ApplyUserId: req.UserId,
		Flag:        req.Flag,
	})
}
