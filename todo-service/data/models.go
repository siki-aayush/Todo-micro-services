package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func New(mongo *mongo.Client) Models {
	client = mongo

	return Models{
		TodoEntry: TodoEntry{},
	}
}

type Models struct {
	TodoEntry TodoEntry
}

type TodoEntry struct {
	ID          string    `json:"id,omitEmpty" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

func (t *TodoEntry) Insert(entry TodoEntry) (*mongo.InsertOneResult, error) {
	collection := client.Database("todo").Collection("todo")

	todo, err := collection.InsertOne(context.TODO(), TodoEntry{
		ID:          entry.ID,
		Name:        entry.Name,
		Description: entry.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		log.Println("Error inserting Todo:", err)
		return nil, err
	}

	return todo, err
}

func (t *TodoEntry) All() ([]*TodoEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("todo").Collection("todo")

	opts := options.Find()
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println("Error finding Todo:", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	var Todos []*TodoEntry

	for cursor.Next(ctx) {
		var item TodoEntry

		err := cursor.Decode(&item)
		if err != nil {
			log.Println("Error decoding Todos:", err)
			return nil, err
		} else {
			Todos = append(Todos, &item)
		}
	}

	return Todos, nil
}

func (t *TodoEntry) GetOne(id string) (*TodoEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("todo").Collection("todo")

	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error converting id to ObjectID:", err)
		return nil, err
	}

	var entry TodoEntry
	err = collection.FindOne(ctx, bson.M{"_id": docId}).Decode(&entry)
	if err != nil {
		log.Println("Error finding Todo:", err)
		return nil, err
	}

	return &entry, nil
}

func (t *TodoEntry) Update() (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("todo").Collection("todo")

	docId, err := primitive.ObjectIDFromHex(t.ID)
	if err != nil {
		log.Println("Error converting id to ObjectID:", err)
		return nil, err
	}

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": docId},
		bson.D{
			{
				"$set", bson.D{
					{"name", t.Name},
					{"description", t.Description},
					{"updated_at", time.Now()},
				}},
		},
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *TodoEntry) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("todo").Collection("todo")

	err := collection.Drop(ctx)
	if err != nil {
		log.Println("Error dropping Todo collection:", err)
		return err
	}

	return nil
}

func (t *TodoEntry) Delete(entry TodoEntry) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("todo").Collection("todo")

	docId, err := primitive.ObjectIDFromHex(entry.ID)
	if err != nil {
		log.Println("Error converting id to ObjectID:", err)
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": docId})
	if err != nil {
		log.Println("Error deleting Todo:", err)
		return err
	}

	return nil
}
