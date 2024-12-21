package main

import (
	"github.com/adnpa/IM/pkg/common/db/mysql/dao"
	"github.com/adnpa/IM/pkg/common/db/redis"
	"github.com/adnpa/IM/pkg/common/logger"
	"time"
)

func main() {
	logger.L().Info("start delete mongodb expired record")

	go func() {
		timer := time.NewTimer(c.Option.AutoReloadInterval)
		for range timer.C {
			uidL := dao.GetAllUserUid()
			if len(uidL) <= 0 {
			} else {
				for _, uid := range uidL {
					minSeq, err := redis.GetUserMinSeq(uid)
					if err != nil {
						continue
					} else {
						err := redis.SetUserMinSeq(uid, minSeq)

					}
				}
			}

			timer.Reset(7 * 24 * time.Hour)
		}
	}()
}
