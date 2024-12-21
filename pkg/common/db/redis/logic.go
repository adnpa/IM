package redis

import (
	"context"
	"strconv"
)

const (
	AccountTempCode               = "ACCOUNT_TEMP_CODE"
	resetPwdTempCode              = "RESET_PWD_TEMP_CODE"
	userIncrSeq                   = "REDIS_USER_INCR_SEQ:" // user incr seq
	appleDeviceToken              = "DEVICE_TOKEN"
	userMinSeq                    = "REDIS_USER_MIN_SEQ:"
	uidPidToken                   = "UID_PID_TOKEN_STATUS:"
	conversationReceiveMessageOpt = "CON_RECV_MSG_OPT:"
)

//seq设计 从0开始增长 redis.Incr执行

func GetPool() Pool {
	return redisPool
}

// GetUserMaxSeq Get the largest Seq
func GetUserMaxSeq(uid string) (uint64, error) {
	key := userIncrSeq + uid
	cli, err := redisPool.Get(context.Background())
	defer cli.Close()
	if err != nil {
		return 0, err
	}
	return GetUint(cli, key)
}

// IncrUserSeq Perform seq auto-increment operation of user messages
func IncrUserSeq(uid string) (uint64, error) {
	cli, err := redisPool.Get(context.Background())
	defer cli.Close()

	key := userIncrSeq + uid
	if err != nil {
		return 0, err
	}
	val, err := cli.Incr(key)
	return uint64(val), err
}

func GetUserMinSeq(uid string) (uint64, error) {
	cli, err := redisPool.Get(context.Background())
	defer cli.Close()
	if err != nil {
		return 0, err
	}

	key := userMinSeq + uid
	return GetUint(cli, key)
}

// SetUserMinSeq Set the user's minimum seq
func SetUserMinSeq(uid string, minSeq uint64) (err error) {
	cli, err := redisPool.Get(context.Background())
	defer cli.Close()
	if err != nil {
		return err
	}

	key := userMinSeq + uid
	err = cli.Set(key, strconv.FormatUint(minSeq, 10))
	return err
}

func SetSingleConversationRecvMsgOpt(userID, conversationID string, opt int32) error {
	cli, err := redisPool.Get(context.Background())
	defer cli.Close()
	if err != nil {
		return err
	}

	key := conversationReceiveMessageOpt + userID
	err = cli.HSet(key, conversationID, opt)
	return err
}

func GetSingleConversationRecvMsgOpt(userID, conversationID string) (int, error) {
	cli, err := redisPool.Get(context.Background())
	defer cli.Close()
	if err != nil {
		return 0, err
	}

	key := conversationReceiveMessageOpt + userID
	res, err := cli.HGet(key, conversationID)
	if err != nil {
		return 0, err
	}
	resInt, err := strconv.Atoi(res)
	if err != nil {
		return 0, err
	}
	return resInt, err

}

func GetUint(cli Conn, key string) (uint64, error) {
	reply, err := cli.Get(key)
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(reply, 10, 64)
}
