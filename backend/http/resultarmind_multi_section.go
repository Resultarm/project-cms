package http

import (
	"backend/entities"
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getMultiSectionItem(c *gin.Context) {
	r, err := dbCollections.MultiSectionItem.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}

	var multiSectionIcon = []entities.MultiSectionItem{}
	if err := r.All(context.Background(), &multiSectionIcon); err != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}

	c.JSON(200, multiSectionIcon)
}

func createMultiSectionItem(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	tab := c.PostForm("tab")
	btnLink := c.PostForm("button_link")
	file, err := c.FormFile("image")

	var tabStruct = entities.Tab{}
	err = json.Unmarshal([]byte(tab), &tabStruct)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	var btnLinkStruct = entities.Link{}
	err = json.Unmarshal([]byte(btnLink), &btnLinkStruct)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	var idImage = primitive.NewObjectID().String()

	err = c.SaveUploadedFile(file, fmt.Sprintf("media/%s", idImage))
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	var multiSectionItem = entities.MultiSectionItem{
		ID:         primitive.NewObjectID(),
		Title:      title,
		Content:    content,
		Tab:        tabStruct,
		ButtonLink: btnLinkStruct,
		ImagePath:  fmt.Sprintf("/media/%s", idImage),
	}

	_, err = dbCollections.MultiSectionItem.InsertOne(context.Background(), multiSectionItem)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	c.JSON(201, gin.H{})
	return
}

func deleteMultiSectionItem(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	_, err = dbCollections.MultiSectionItem.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	c.JSON(204, gin.H{})
}

func getMultiSection(c *gin.Context) {
	r := dbCollections.MultiSection.FindOne(context.Background(), bson.M{})
	var data = &entities.MultiSections{}
	if err := r.Decode(&data); err != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}

	c.JSON(200, data)
}

func createUpdateMultiSection(c *gin.Context) {
	var data = entities.MultiSections{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	ctx := context.Background()

	s, err := checkExists(ctx, dbCollections.MultiSection)
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	var old = entities.MultiSections{}
	err = s.Decode(&old)
	if err != nil {
		data.ID = primitive.NewObjectID()
		_, err = dbCollections.MultiSection.InsertOne(ctx, data)
		if err != nil {
			c.JSON(400, gin.H{"message": "Bad Request"})
			return
		}

		c.JSON(201, gin.H{})
		return
	}

	_, err = dbCollections.MultiSection.UpdateOne(ctx, bson.M{"_id": bson.M{"$eq": old.ID}}, bson.M{"$set": bson.M{"title": data.Title, "content": data.Content}})
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}
	c.JSON(201, gin.H{})
	return
}
