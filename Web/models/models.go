package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	Id primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Movie string `bson:"movie,omitempty"`
	Watched bool `bson:"watched,omitempty"`
}

