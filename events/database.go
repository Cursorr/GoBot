package events

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var projectDB *mongo.Database
var invitesCollection *mongo.Collection

var ctx context.Context

func init() { 
	connectDB() 
	ctx = context.TODO()
}

type UserData struct {
	ID      	string			`bson:"_id"`
	GuildID 	string          `bson:"guild_id"`
	InviterID 	string 			`bson:"inviter_id"`
	Invites 	int            	`bson:"invites"`
	Left 		int				`bson:"left"`
}

func connectDB() {
	clientOptions := options.Client()
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	projectDB = client.Database("gobot")
	invitesCollection = projectDB.Collection("invites")

}

func UpdateUserData(guildID, userID string, query bson.D) {
	filter := bson.D{
		{Key: "_id", Value: userID}, 
		{Key: "guild_id", Value: guildID}}
	opts := options.Update().SetUpsert(true)

	_, err := invitesCollection.UpdateOne(ctx, filter, query, opts)

	if err != nil {
		panic(err)
	}
}

func GetUserData(guildID, userID string) (UserData, error) {
	filter := bson.D{
		{Key: "_id", Value: userID}, 
		{Key: "guild_id", Value: guildID}}

	var data UserData

	err := invitesCollection.FindOne(ctx, filter).Decode(&data)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return UserData{userID, userID, "", 0, 0}, err
		}
		return data, err
	}

	return data, nil
}
