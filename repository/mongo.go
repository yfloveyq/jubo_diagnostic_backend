package repository

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"jubo.com/eric/diagnostic/domain"
	"log"
	"os"
)

type MongoJuboRepository struct {
	uri         string
	ctx         context.Context
	client      *mongo.Client
	patientColl *mongo.Collection
	orderColl   *mongo.Collection
	lookup      bson.A
}

func NewMongoJuboRepository() *MongoJuboRepository {
	ctx := context.TODO()
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	lookup := bson.A{
		bson.M{
			"$lookup": bson.M{
				"from":         "orders",    // The name of the "orders" collection
				"localField":   "orderList", // Field in the "patients" collection
				"foreignField": "_id",       // Field in the "orders" collection
				"as":           "Orders",    // Name for the joined array of orders
			},
		},
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	patientColl := client.Database("jubo").Collection("patients")
	orderColl := client.Database("jubo").Collection("orders")

	return &MongoJuboRepository{
		uri:         uri,
		ctx:         ctx,
		client:      client,
		patientColl: patientColl,
		orderColl:   orderColl,
		lookup:      lookup,
	}
}

func (m *MongoJuboRepository) ListPatients() []domain.Patient {
	cursor, err := m.patientColl.Aggregate(m.ctx, m.lookup)
	if err != nil {
		panic(err)
	}

	var results []domain.Patient
	if err = cursor.All(m.ctx, &results); err != nil {
		panic(err)
	}
	return results
}

func (m *MongoJuboRepository) UpdatePatient(patient domain.Patient) (successful bool, err error) {
	filter := bson.D{{"_id", patient.Id}}
	update := bson.D{{"$set", patient}}
	result, err := m.patientColl.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount == 1, nil
}

func (m *MongoJuboRepository) InsertOrder(order domain.Order) (success bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (m *MongoJuboRepository) UpdateOrder(order domain.Order) (success bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (m *MongoJuboRepository) DeleteOrder(id string) (successful bool, err error) {
	//TODO implement me
	panic("implement me")
}
