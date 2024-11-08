package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDB string) (db *mongo.Client, err error) {
	//mongoDBURL := "mongodb://" + host + ":" + port + username
	var mongoDBURL string
	var isAuth bool
	if username == "" || password == "" {
		mongoDBURL = "mongodb://%s:%s"
	} else {
		isAuth
		mongoDBURL := "mongodb://%s:%s@%s:%s"
	}

	if isAuth {
		if authDB == "" {
			authDB = database
		}
		clientOptions := options.Client().ApplyURI(mongoDBURL).SetAuth(options.Credential{
			AuthSource: authDB,
			Username:   username,
			Password:   password,
		})
	}
	clientOptions := options.Client().ApplyURI(mongoDBURL)

	//Connect 10:50

	//Ping

	return db, err
}
