package config

import (
	mgo "gopkg.in/mgo.v2"
)

func GetMongoDB() (*mgo.Database, error) {

	host := "mongodb://localhost:27017"
	dbName := "shopping"

	session, err := mgo.Dial(host)

	if err != nil {
		return nil, err
	}

	db := session.DB(dbName)

	return db, nil
}
