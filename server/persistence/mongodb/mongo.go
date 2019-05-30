package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Connection struct {
	conn *mongo.Client
}

func (c Connection) GetDatabase(s string) *mongo.Database{
	return c.conn.Database(s)
}

func NewConnection() *Connection {
	m := new(Connection)
	clientOptions := options.Client().ApplyURI("mongodb://jl:vwdilab@localhost:27017/pairot")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	m.conn = client
	return m
}