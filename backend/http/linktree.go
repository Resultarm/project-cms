package http

import (
	"backend/db"
	"backend/entities"
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbCollections *db.Db

func GetUserBy(key, value string, users *mongo.Collection) (*entities.User, error) {
	s := users.FindOne(context.Background(), bson.M{key: value})
	if s.Err() != nil {
		return nil, s.Err()
	}

	var decoded = entities.User{}
	err := s.Decode(&decoded)
	if err != nil {
		return nil, err
	}

	return &decoded, nil
}

func HashPassword(password string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
}

func GenerateJWT(username, role string) (string, error) {
	var mySigningKey = []byte(secretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func login(c *gin.Context) {
	var user = &entities.User{}
	err := c.BindJSON(user)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	existentUser, err := GetUserBy("username", user.Username, dbCollections.Users)
	if err != nil {
		c.JSON(400, gin.H{"message": "Wrong credential"})
		return
	}

	if HashPassword(user.Password) != existentUser.Password {
		c.JSON(400, gin.H{"message": "Wrong credential"})
		return
	}

	token, err := GenerateJWT(existentUser.Username, fmt.Sprintf("%d", existentUser.Permissions))
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func createImage(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	image := entities.Image{ID: primitive.NewObjectID(), FilePath: fmt.Sprintf("%s", "/media/"+file.Filename)}
	dbCollections.Logo.InsertOne(context.Background(), image)

	err = c.SaveUploadedFile(file, fmt.Sprintf("%s", "media/"+file.Filename))
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(201, gin.H{})
}

func getImage(c *gin.Context) {
	var image = &entities.Image{}
	err := dbCollections.Logo.FindOne(context.Background(), bson.M{}).Decode(image)
	if err != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}

	c.JSON(200, image)
}

func deleteImage(c *gin.Context) {
	err := dbCollections.Logo.Drop(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"message": "Not Found"})
		return
	}

	c.JSON(204, gin.H{"message": "deleted"})
}

func read(c *gin.Context) {
	r := dbCollections.Data.FindOne(context.Background(), bson.M{})
	var data = &entities.Data{}
	if err := r.Decode(&data); err != nil {
		if err == mongo.ErrNoDocuments {

			c.JSON(404, gin.H{"message": "Not Found"})
			return
		}
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	c.JSON(200, data)

}

func createLinksData(c *gin.Context) {
	var data = entities.Data{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	if err := data.Validate(); err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	result, err := dbCollections.Data.InsertOne(context.Background(), data)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	c.JSON(201, result)

}

func update(c *gin.Context) {
	idHex := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}
	filter := bson.M{"_id": id}

	var data = &entities.Data{}
	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	var updatedData = bson.D{{"$set", bson.D{{"content", data.Content}, {"links", data.Links}, {"title", data.Title}}}}

	_, err = dbCollections.Data.UpdateOne(context.TODO(), filter, updatedData)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	c.JSON(200, gin.H{})
}
func delete(c *gin.Context) {
	idHex := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	_, err = dbCollections.Data.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	c.JSON(204, gin.H{})
}
