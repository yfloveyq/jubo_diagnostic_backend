package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Patient struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string
	OrderList []primitive.ObjectID `bson:"orderList"`
	Orders    []Order
}
