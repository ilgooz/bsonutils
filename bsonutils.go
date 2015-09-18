package bsonutils

import "gopkg.in/mgo.v2/bson"

func ObjectId(s string) (bson.ObjectId, bool) {
	var id bson.ObjectId
	if !bson.IsObjectIdHex(s) {
		return id, false
	}
	return bson.ObjectIdHex(s), true
}
