package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"id,omitempty" bason:"_id" -line:"id"`
	Username    string             `json:"username" -line:"username"`
	Password    string             `json:"password,omitempty" -line:"password"`
	Permissions int                `json:"permissions,omitempty" -line:"permissions"`
}
