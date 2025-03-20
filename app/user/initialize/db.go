package initialize

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adnpa/IM/app/user/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB() {
	c := global.ServerConfig.MysqlInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.Name)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // 禁用彩色打印
		},
	)

	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// for i := 0; i < 100; i++ { // 生成 100 条假数据
	// 	user := generateFakeUser()
	// 	result := global.DB.Create(&user)
	// 	if result.Error != nil {
	// 		log.Printf("插入用户失败: %v", result.Error)
	// 	}
	// }
}

// 测试,生成假数据
// func generateFakeUser() model.User {
// 	return model.User{
// 		Mobile:   gofakeit.Phone(),
// 		Email:    gofakeit.Email(),
// 		Passwd:   gofakeit.Password(true, true, true, true, false, 12),
// 		Salt:     []byte(gofakeit.UUID()),
// 		Nickname: gofakeit.Username(),
// 		// Avatar:    gofakeit.ImageURL(100, 100),
// 		// Sex:       int8(gofakeit.Number(0, 2)),
// 		Memo:      gofakeit.HackerPhrase(),
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}
// }
