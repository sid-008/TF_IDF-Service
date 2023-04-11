package collection

import "go.mongodb.org/mongo-driver/mongo"

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("DB1").Collection("Posts")
	return collection
}
