package models

import (
	//"gopkg.in/mgo.v2/bson"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

)

type DbModel struct{
	ColName string
	DbName string
	Collection *mongo.Collection
}

func (d DbModel) GetCollection(client *mongo.Client) *mongo.Collection{
	return client.Database(d.DbName).Collection(d.ColName)
}
