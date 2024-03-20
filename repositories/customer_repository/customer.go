package customer_repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/adrian7123/rinha-backend-2024-q1-gin/configs"
	"github.com/adrian7123/rinha-backend-2024-q1-gin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var coll = configs.GetCollection("customers")

func Get(filter interface{}) []models.Customer {
	customers := []models.Customer{}

	cursor, err := coll.Aggregate(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

	if err := cursor.All(context.TODO(), &customers); err != nil {
		panic(err)
	}

	return customers
}

func GetOne(filter interface{}, projection interface{}) (*models.Customer, error) {
	var customer *models.Customer

	doc := coll.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection))

	if doc.Err() != nil {
		fmt.Println(doc.Err())
		return nil, errors.New("documento nao encontrado")
	}

	err := doc.Decode(&customer)

	if err != nil {
		log.Fatal(err)
	}

	return customer, nil
}

func Update(customer models.Customer) error {
	if _, err := coll.UpdateByID(context.TODO(), bson.D{{Key: "id", Value: customer.Id}}, customer); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func CreateMany(customers []interface{}) {
	if _, err := coll.InsertMany(context.TODO(), customers); err != nil {
		panic(err)
	}
}

func DeleteAll() {
	if _, err := coll.DeleteMany(context.TODO(), bson.D{}); err != nil {
		panic(err)
	}
}
