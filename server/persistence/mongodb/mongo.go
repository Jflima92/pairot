package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func NewConnection() *Connection {
	m := new(Connection)
	dbUsername := os.Getenv("MONGO_USERNAME")
	dbPassword := os.Getenv("MONGO_PASSWORD")
	dbName := os.Getenv("MONGO_DATABASE")
	dbPort := os.Getenv("MONGO_PORT")
	uri := fmt.Sprintf("mongodb://%s:%s@localhost:%s/%s", dbUsername, dbPassword, dbPort, dbName)
	clientOptions := options.Client().ApplyURI(uri)
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

type Connection struct {
	conn *mongo.Client
}

func (c Connection) GetDatabase(s string) *mongo.Database{
	return c.conn.Database(s)
}

func (c Connection) FindTeamByName(teamName string) ([]byte, error){
	db := c.conn.Database("pairot")
	var collection = db.Collection("Teams")
	filter := bson.D{{"name", teamName}}
	b, err := collection.FindOne(context.TODO(), filter).DecodeBytes()
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (c Connection) Decode (data []byte, val interface{}) error {
	return bson.Unmarshal(data, val)
}

func (c Connection) UpdateTeamMembers(teamName string, members interface{}) error {
	db := c.conn.Database("pairot")
	var collection = db.Collection("Teams")
	filter := bson.D{{"name", teamName}}
	update := bson.M{
		"$set": bson.M{
			"members": &members,
		},
	}
	_, e := collection.UpdateOne(context.TODO(), filter, update)
	return e
}