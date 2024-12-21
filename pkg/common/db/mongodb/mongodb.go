package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	CollectionChat = "msg"
)

var conn *Conn

func init() {
	conn, _ = NewConn()
}

func GetConn() *Conn {
	return conn
}

type Conn struct {
	delegate *mongo.Client
}

func NewConn() (*Conn, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	//mongodb://foo:bar@localhost:27017
	user := "root"
	pwd := "example"

	uri := fmt.Sprintf("mongodb://%s:%s@192.168.1.129:27017", user, pwd)
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

func (c *Conn) Find(filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	coll := c.delegate.Database("IM").Collection(CollectionChat)

	return coll.Find(ctx, filter, opts...)
}

func (c *Conn) FindOne(filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	coll := c.delegate.Database("IM").Collection(CollectionChat)

	return coll.FindOne(ctx, filter, opts...)
}

func (c *Conn) FindOneAndUpdate(filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	coll := c.delegate.Database("IM").Collection(CollectionChat)

	return coll.FindOneAndUpdate(ctx, filter, update, opts...).Err()
}

func (c *Conn) InsertOne(document interface{}, opts ...*options.InsertOneOptions) (interface{}, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	coll := c.delegate.Database("IM").Collection(CollectionChat)
	one, err := coll.InsertOne(ctx, document, opts...)
	if err != nil {
		return nil, err
	}
	return one.InsertedID, err
}

func (c *Conn) Insert(documents []interface{}, opts ...*options.InsertManyOptions) ([]interface{}, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	coll := c.delegate.Database("IM").Collection(CollectionChat)

	many, err := coll.InsertMany(ctx, documents, opts...)
	if err != nil {
		return nil, err
	}
	return many.InsertedIDs, err
}

func (c *Conn) UpdateOne(filter interface{}, update interface{}, opts ...*options.UpdateOptions) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	coll := c.delegate.Database("IM").Collection(CollectionChat)

	_, err := coll.UpdateOne(ctx, filter, update, opts...)
	return err
}

func (c *Conn) Update(filter interface{}, update interface{}, opts ...*options.UpdateOptions) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	coll := c.delegate.Database("IM").Collection(CollectionChat)

	_, err := coll.UpdateMany(ctx, filter, update, opts...)
	return err
}
