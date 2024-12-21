package test

import (
	"context"
	"github.com/adnpa/IM/pkg/common/db/mongodb"
	"github.com/adnpa/IM/pkg/common/db/mysql"
	"github.com/adnpa/IM/pkg/common/db/redis"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestMysql(t *testing.T) {
	mysql.Close()
}

func TestRedis(t *testing.T) {
	ctx := context.Background()
	cli, err := redis.GetPool().Get(ctx)
	if err != nil {
		panic(t)
	}
	defer func(cli redis.Conn) {
		err := cli.Close()
		if err != nil {
			panic(t)
		}
	}(cli)

	exists, err := cli.Exists("abc", "myhash")
	t.Log(exists)
	t.Log(err)
	//err = cli.Set("abc", "abc")
	//t.Log(err == nil)
	//res, err := cli.Get("abc")
	//t.Log(res)
	//err = cli.Set("abc", "bcd")
	//t.Log(err)
	//res, err = cli.Get("abc")
	//t.Log(res)

	//err = cli.HSet("myhash", "key1", "value1", "key2", "value2")
	//err = cli.HSet("myhash", []string{"key1", "value1", "key2", "value2"})
	//t.Log(err)
	//hGet, err := cli.HGet("myhash", "key1")
	//t.Log(hGet)
}

func TestMongo(t *testing.T) {
	conn, err := mongodb.NewConn()
	defer conn.Close()
	if err != nil {
		t.Fatal()
	}
	//_, err = conn.InsertOne(bson.M{"uid": "111", "hello": "world"})
	//_, err = conn.InsertOne(bson.M{"uid": "111", "you": "haha"})
	//t.Log("insertOne", insert, err)
	filter := bson.D{{"uid", "222"}}
	err = conn.FindOneAndUpdate(filter, bson.M{"$push": bson.M{"msg": "bbb"}})
	if err != nil {
		t.Log(err)
		_, err := conn.InsertOne(bson.D{{"uid", "222"}, {"nnn", "nnn"}})
		if err != nil {
			t.Fatal()
		}
		t.Log("insert")
	}

	one := conn.FindOne(filter)
	raw, err := one.Raw()
	if err != nil {
		t.Fatal()
	}
	t.Log(raw)
	//cursor, err := one
	//if err != nil {
	//	t.Fatal()
	//}
	//for cursor.Next(context.Background()) {
	//	t.Log(cursor.Current)
	//}

}

func TestA(t *testing.T) {
	t.Log(1001 / 1000)
}
