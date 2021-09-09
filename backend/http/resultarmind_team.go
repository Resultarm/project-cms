package http

import (
	"backend/entities"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getMembers(c *gin.Context) {
	r, err := dbCollections.Members.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}

	var members = []entities.Member{}
	if err := r.All(context.Background(), &members); err != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}

	c.JSON(200, members)
}

func createMember(c *gin.Context) {
	name := c.PostForm("name")
	function := c.PostForm("function")
	facebook := c.PostForm("facebook")
	twitter := c.PostForm("twitter")
	youtube := c.PostForm("youtube")
	instagram := c.PostForm("instagram")
	file, err := c.FormFile("image")

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

	var member = entities.Member{
		ID:        primitive.NewObjectID(),
		Name:      name,
		Function:  function,
		Facebook:  facebook,
		Instagram: instagram,
		Youtube:   youtube,
		Twitter:   twitter,
		ImagePath: fmt.Sprintf("/media/%s", idImage),
	}

	_, err = dbCollections.Members.InsertOne(context.Background(), member)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	c.JSON(201, gin.H{})
	return
}

func deleteMember(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	_, err = dbCollections.Members.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	c.JSON(204, gin.H{})
}

func getTeam(c *gin.Context) {
	r := dbCollections.Team.FindOne(context.Background(), bson.M{})
	var data = &entities.Team{}
	if err := r.Decode(&data); err != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}

	c.JSON(200, data)
}

func createUpdateTeam(c *gin.Context) {
	var data = entities.Team{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	ctx := context.Background()

	s, err := checkExists(ctx, dbCollections.Team)
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	var old = entities.Team{}
	err = s.Decode(&old)
	if err != nil {
		data.ID = primitive.NewObjectID()
		_, err = dbCollections.Team.InsertOne(ctx, data)
		if err != nil {
			c.JSON(400, gin.H{"message": "Bad Request"})
			return
		}

		c.JSON(201, gin.H{})
		return
	}

	_, err = dbCollections.Team.UpdateOne(ctx, bson.M{"_id": bson.M{"$eq": old.ID}}, bson.M{"$set": bson.M{"title": data.Title, "content": data.Content}})
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}
	c.JSON(201, gin.H{})
	return

}
