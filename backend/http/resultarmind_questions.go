package http

import (
	"backend/entities"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getQuestions(c *gin.Context) {
	r := dbCollections.Questions.FindOne(context.Background(), bson.M{})
	var data = &entities.Questions{}
	if err := r.Decode(&data); err != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}

	c.JSON(200, data)
}

func createUpdateQuestions(c *gin.Context) {
	var data = entities.Questions{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	ctx := context.Background()

	s, err := checkExists(ctx, dbCollections.Questions)
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	var old = entities.Questions{}
	err = s.Decode(&old)
	if err != nil {
		data.ID = primitive.NewObjectID()
		_, err = dbCollections.Questions.InsertOne(ctx, data)
		if err != nil {
			c.JSON(400, gin.H{"message": "Bad Request"})
			return
		}

		c.JSON(201, gin.H{})
		return
	}

	_, err = dbCollections.Questions.UpdateOne(ctx, bson.M{"_id": bson.M{"$eq": old.ID}}, bson.M{"$set": bson.M{"title": data.Title, "content": data.Content}})
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}
	c.JSON(201, gin.H{})
	return
}

func getQuestion(c *gin.Context) {
	r, err := dbCollections.Question.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}

	var question = []entities.Question{}
	if err := r.All(context.Background(), &question); err != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}

	c.JSON(200, question)
}

func createQuestion(c *gin.Context) {
	question := c.PostForm("question")
	answer := c.PostForm("answer")
	file, err := c.FormFile("image")

	if err != nil {
		fmt.Println("Here?")
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	var idImage string = primitive.NewObjectID().Hex()

	err = c.SaveUploadedFile(file, fmt.Sprintf("media/%s", idImage))
	if err != nil {
		fmt.Println("Here?2")
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	var member = entities.Question{
		ID:        primitive.NewObjectID(),
		Question:  question,
		Answer:    answer,
		ImagePath: fmt.Sprintf("/media/%s", idImage),
	}

	_, err = dbCollections.Question.InsertOne(context.Background(), member)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	c.JSON(201, gin.H{})
	return
}

func deleteQuestion(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	_, err = dbCollections.Question.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	c.JSON(204, gin.H{})
}
