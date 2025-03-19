package initialize

import (
	"github.com/adnpa/IM/app/presence/global"
	"github.com/adnpa/IM/pkg/common/storage"
	"github.com/redis/go-redis/v9"
)

func InitDB() {
	// c := global.ServerConfig.MysqlInfo
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	c.User, c.Password, c.Host, c.Port, c.Name)
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold: time.Second,   // 慢 SQL 阈值
	// 		LogLevel:      logger.Silent, // Log level
	// 		Colorful:      true,          // 禁用彩色打印
	// 	},
	// )

	// var err error
	// global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		SingularTable: true,
	// 	},
	// 	Logger: newLogger,
	// })
	// if err != nil {
	// 	panic(err)
	// }

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		Protocol: 3,  // specify 2 for RESP 2 or 3 for RESP 3
	})
	global.RedisPool = storage.NewPool(rdb)
}
