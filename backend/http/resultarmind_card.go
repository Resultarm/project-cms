package http

import (
	"backend/entities"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getCards(c *gin.Context) {
	r := dbCollections.Card.FindOne(context.Background(), bson.M{})
	var data = &entities.Cards{}
	if err := r.Decode(&data); err != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}

	c.JSON(200, data)
}

func createUpdateCards(c *gin.Context) {
	var data = entities.Cards{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	if data.Title == "" || data.Content == "" {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	if len(data.Cards) == 0 {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	ctx := context.Background()

	s, err := checkExists(ctx, dbCollections.Card)
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	// Generate new id for every card
	for index, _ := range data.Cards {
		data.Cards[index].ID = primitive.NewObjectID()
	}

	var old = entities.Cards{}
	err = s.Decode(&old)
	if err != nil {
		data.ID = primitive.NewObjectID()

		_, err = dbCollections.Card.InsertOne(ctx, data)
		if err != nil {
			c.JSON(400, gin.H{"message": "Bad Request"})
			return
		}

		c.JSON(201, gin.H{})
		return
	}

	_, err = dbCollections.Card.UpdateOne(ctx, bson.M{"_id": bson.M{"$eq": old.ID}}, bson.M{"$set": bson.M{"title": data.Title, "content": data.Content, "cards": data.Cards}})
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}
	c.JSON(201, gin.H{})
	return
}
