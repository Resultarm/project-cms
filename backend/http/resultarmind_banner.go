package http

import (
	"backend/entities"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getBanner(c *gin.Context) {
	r := dbCollections.Banner.FindOne(context.Background(), bson.M{})
	var data = &entities.Banner{}
	if err := r.Decode(&data); err != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}

	c.JSON(200, data)
}

func createUpdateBanner(c *gin.Context) {
	var data = entities.Banner{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	if data.Content == "" || data.Title == "" || data.Youtube == "" {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	ctx := context.Background()

	s, err := checkExists(ctx, dbCollections.Banner)
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	var old = entities.Banner{}
	err = s.Decode(&old)
	if err != nil {
		data.ID = primitive.NewObjectID()
		_, err = dbCollections.Banner.InsertOne(ctx, data)
		if err != nil {
			c.JSON(400, gin.H{"message": "Bad Request"})
			return
		}

		c.JSON(201, gin.H{})
		return
	}

	_, err = dbCollections.Banner.UpdateOne(ctx, bson.M{"_id": bson.M{"$eq": old.ID}}, bson.M{"$set": bson.M{"content": data.Content, "title": data.Title, "youtube": data.Youtube}})
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}
	c.JSON(201, gin.H{})
	return
}
