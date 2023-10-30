package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Patient struct {
	Id        primitive.ObjectID `bson:"_id"`
	Name      string
	orderList []primitive.ObjectID `bson:"orderList"`
	Orders    []Order
}
