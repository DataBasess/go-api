package utils

import (
	"context"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

func DocumentList(cur *mongo.Cursor) []interface{} {
	var list []interface{}
	for cur.Next(context.TODO()) {
		var model bson.M
		err := cur.Decode(&model)
		if err != nil {
			log.Fatal(err)
			return list
		}
		// log.Println("model", model)
		list = append(list, model)
	}
	// once exhausted, close the cursor
	cur.Close(context.TODO())

	return list
}

func StructToMap(data interface{}) (map[string]interface{}, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mapData := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}
