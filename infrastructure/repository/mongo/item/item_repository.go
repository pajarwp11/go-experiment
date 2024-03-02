package item

import (
	"context"

	model "pajarwp11/go-experiment/core/models/item"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collectionItem = "item"
)

type (
	Repository struct {
		DB *mongo.Database
	}
)

func New(db *mongo.Database) *Repository {
	return &Repository{
		DB: db,
	}
}

func (repo *Repository) GetItemByID(id string) (doc model.Item, err error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error("error get item: " + err.Error())
	}
	filter := bson.M{
		"_id": objectID,
	}

	err = repo.DB.Collection(collectionItem).FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		log.Error("error get item: " + err.Error())
	}
	return
}

func (repo *Repository) InsertItem(itemData *model.InsertItem) error {
	_, err := repo.DB.Collection(collectionItem).InsertOne(context.TODO(), itemData)
	if err != nil {
		log.Error("error insert item: " + err.Error())
	}
	return err
}

func (repo *Repository) UpdateItem(itemData *model.UpdateItem) error {
	objectID, err := primitive.ObjectIDFromHex(itemData.ID)
	if err != nil {
		log.Error("error update item: " + err.Error())
	}
	where := bson.M{
		"_id": objectID,
	}

	update := bson.M{"$set": bson.M{
		"name":     itemData.Name,
		"quantity": itemData.Quantity,
		"weight":   itemData.Weight,
	}}

	_, err = repo.DB.Collection(collectionItem).UpdateOne(context.TODO(), where, update)
	if err != nil {
		log.Error("error update item: " + err.Error())
	}
	return err
}

func (repo *Repository) DeleteItem(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error("error delete item: " + err.Error())
	}
	filter := bson.M{
		"_id": objectID,
	}
	_, err = repo.DB.Collection(collectionItem).DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Error("error delete item: " + err.Error())
	}
	return err
}
