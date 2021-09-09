package http

import (
	"backend/entities"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func checkExists(ctx context.Context, col *mongo.Collection) (*mongo.SingleResult, error) {
	s := col.FindOne(ctx, bson.M{})
	if s.Err() == mongo.ErrNoDocuments {
		return s, nil
	}
	if s.Err() != nil {
		return nil, s.Err()
	}

	return s, nil
}

func getMenu(c *gin.Context) {
	r := dbCollections.Menu.FindOne(context.Background(), bson.M{})
	var data = &entities.Menu{}
	if err := r.Decode(&data); err != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}

	c.JSON(200, data)
}

func createUpdateMenu(c *gin.Context) {
	var data = entities.Menu{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	if len(data.Links) == 0 {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}
	ctx := context.Background()

	s, err := checkExists(ctx, dbCollections.Menu)
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	var old = entities.Menu{}
	err = s.Decode(&old)
	if err != nil {
		data.ID = primitive.NewObjectID()
		_, err = dbCollections.Menu.InsertOne(ctx, data)
		if err != nil {
			c.JSON(400, gin.H{"message": "Bad Request"})
			return
		}

		c.JSON(201, gin.H{})
		return
	}

	_, err = dbCollections.Menu.UpdateOne(ctx, bson.M{"_id": bson.M{"$eq": old.ID}}, bson.M{"$set": bson.M{"links": data.Links}})
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}
	c.JSON(201, gin.H{})
	return
}
