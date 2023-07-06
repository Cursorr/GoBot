package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Instance *MongoDatabase

func init() { 
	Instance = NewMongoDatabase()
	Instance.connectDB()
}

type UserData struct {
	ID      	string			`bson:"_id"`
	GuildID 	string          `bson:"guild_id"`
	InviterID 	string 			`bson:"inviter_id"`
	Invites 	int            	`bson:"invites"`
	Left 		int				`bson:"left"`
}

type Database interface {
	UpdateUserData(guildID, userID string, query bson.D)
	GetUserData(guildID, userID string) (UserData, error)
}

type MongoDatabase struct {
	client             *mongo.Client
	projectDB          *mongo.Database
	invitesCollection  *mongo.Collection
	ctx                context.Context
}

func NewMongoDatabase() *MongoDatabase {
	return &MongoDatabase{}
}

func (db *MongoDatabase) connectDB() {
	clientOptions := options.Client()
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	db.client = client
	db.projectDB = client.Database("gobot")
	db.invitesCollection = db.projectDB.Collection("invites")
	db.ctx = context.TODO()
}

func (db *MongoDatabase) UpdateUserData(guildID, userID string, query bson.D) {
	filter := bson.D{
		{Key: "_id", Value: userID}, 
		{Key: "guild_id", Value: guildID}}
	opts := options.Update().SetUpsert(true)

	_, err := db.invitesCollection.UpdateOne(db.ctx, filter, query, opts)
	if err != nil {
		panic(err)
	}
}

func (db *MongoDatabase) GetUserData(guildID, userID string) (UserData, error) {
	filter := bson.D{
		{Key: "_id", Value: userID}, 
		{Key: "guild_id", Value: guildID}}

	var data UserData

	err := db.invitesCollection.FindOne(db.ctx, filter).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return UserData{userID, userID, "", 0, 0}, err
		}
		return data, err
	}

	return data, nil
}
