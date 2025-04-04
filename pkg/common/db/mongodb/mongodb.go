package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DB_Name = "im"
)

// var conn *Conn

// func init() {
// 	conn, _ = NewConn()
// }

func GetById(name string, id int64) (*mongo.SingleResult, error) {
	filter := bson.M{"id": id}
	return Get(name, filter)
}

func Get(name string, filter interface{}) (*mongo.SingleResult, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	c, _ := NewConn()
	defer c.Close()
	coll := c.delegate.Database(DB_Name).Collection(name)

	res := coll.FindOne(ctx, filter)
	return res, nil
}

func GetAll(name string, filter interface{}) (*mongo.Cursor, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	c, _ := NewConn()
	defer c.Close()
	coll := c.delegate.Database(DB_Name).Collection(name)

	return coll.Find(ctx, filter)
}

func GetDecode(name string, filter interface{}, result interface{}) error {
	r, err := Get(name, filter)
	if err != nil {
		return err
	}
	return r.Decode(result)
}

func Exist(name string, filter interface{}) bool {
	one, _ := Get(name, filter)
	r, _ := one.Raw()
	return r.Validate() == nil
}

func Insert(name string, data interface{}) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	c, err := NewConn()
	if err != nil {
		return err
	}
	defer c.Close()

	coll := c.delegate.Database(DB_Name).Collection(name)
	_, err = coll.InsertOne(ctx, data)
	return err
}

func Delete(name string, filter interface{}) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	c, err := NewConn()
	if err != nil {
		return err
	}
	defer c.Close()

	coll := c.delegate.Database(DB_Name).Collection(name)
	_, err = coll.DeleteOne(ctx, filter)
	return err
}

func Update(name string, filter interface{}, update interface{}) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	c, err := NewConn()
	if err != nil {
		return err
	}
	defer c.Close()

	coll := c.delegate.Database(DB_Name).Collection(name)
	_, err = coll.UpdateOne(ctx, filter, update)
	return err
}

type Conn struct {
	delegate *mongo.Client
}

func NewConn() (*Conn, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	//mongodb://foo:bar@localhost:27017
	user := "root"
	pwd := "123456"

	uri := fmt.Sprintf("mongodb://%s:%s@localhost:27017", user, pwd)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &Conn{
		delegate: client,
	}, nil
}

func (c *Conn) Close() {
	err := c.delegate.Disconnect(context.TODO())
	if err != nil {
		return
	}
}

// func (c *Conn) Find(filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
// 	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancelFunc()
// 	coll := c.delegate.Database("IM").Collection(CollectionChat)

// 	return coll.Find(ctx, filter, opts...)
// }

// func (c *Conn) FindOne(filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
// 	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancelFunc()
// 	coll := c.delegate.Database("IM").Collection(CollectionChat)

// 	return coll.FindOne(ctx, filter, opts...)
// }

// func (c *Conn) FindOneAndUpdate(filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) error {
// 	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancelFunc()
// 	coll := c.delegate.Database("IM").Collection(CollectionChat)

// 	return coll.FindOneAndUpdate(ctx, filter, update, opts...).Err()
// }

// func (c *Conn) InsertOne(document interface{}, opts ...*options.InsertOneOptions) (interface{}, error) {
// 	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancelFunc()
// 	coll := c.delegate.Database("IM").Collection(CollectionChat)
// 	one, err := coll.InsertOne(ctx, document, opts...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return one.InsertedID, err
// }

// func (c *Conn) Insert(documents []interface{}, opts ...*options.InsertManyOptions) ([]interface{}, error) {
// 	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancelFunc()
// 	coll := c.delegate.Database("IM").Collection(CollectionChat)

// 	many, err := coll.InsertMany(ctx, documents, opts...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return many.InsertedIDs, err
// }

// func (c *Conn) UpdateOne(filter interface{}, update interface{}, opts ...*options.UpdateOptions) error {
// 	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancelFunc()
// 	coll := c.delegate.Database("IM").Collection(CollectionChat)

// 	_, err := coll.UpdateOne(ctx, filter, update, opts...)
// 	return err
// }

// func (c *Conn) Update(filter interface{}, update interface{}, opts ...*options.UpdateOptions) error {
// 	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancelFunc()
// 	coll := c.delegate.Database("IM").Collection(CollectionChat)

// 	_, err := coll.UpdateMany(ctx, filter, update, opts...)
// 	return err
// }
