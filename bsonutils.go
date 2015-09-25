package bsonutils

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

// ObjectId returns the bson.ObjectId value represented by str. Any invalid value causes to return an error
func ObjectId(str string) (id bson.ObjectId, err error) {
	if !bson.IsObjectIdHex(str) {
		return id, errors.New("string doesn't represents a bson.ObjectId")
	}
	return bson.ObjectIdHex(str), nil
}
