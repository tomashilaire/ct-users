package testmongorepo

import (
	"context"
	"log"
	"test/internal/core/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type testRepository struct {
	c *mongo.Collection
}

const TestCollection = "testing"

func NewTestRepository(conn *conn) *testRepository {
	return &testRepository{c: conn.DB().Collection(TestCollection)}
}

func (r *testRepository) Create(t *domain.Test) error {
	_, err := r.c.InsertOne(context.TODO(), t)
	if err != nil {
		log.Println("Error in Repository -> Create()", err)
		return err
	}
	return nil
}

func (r *testRepository) Update(t *domain.Test) error {
	_, err := r.c.UpdateByID(context.TODO(), t.Id, t)
	if err != nil {
		log.Println("Error in Repository -> Update()", err)
		return err
	}
	return nil
}

func (r *testRepository) ShowAll() (t []*domain.Test, err error) {
	cur, err := r.c.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println("Error in Repository -> ShowAll()", err)
		return []*domain.Test{}, err
	}

	var results []*domain.Test
	for cur.Next(context.TODO()) {
		var elem domain.Test
		err := cur.Decode(&elem)
		if err != nil {
			log.Println("Error in Repository -> ShowAll()", err)
			return []*domain.Test{}, err
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Println("Error in Repository -> ShowAll()", err)
		return []*domain.Test{}, err
	}

	cur.Close(context.TODO())

	return results, nil
}

func (r *testRepository) ShowById(id string) (t *domain.Test, err error) {
	findResult := r.c.FindOne(context.TODO(), bson.M{"id": id})

	decodeErr := findResult.Decode(&t)
	if decodeErr != nil {
		log.Println("Error in Repository -> ShowById()", decodeErr)
		return &domain.Test{}, decodeErr
	}

	return t, nil
}

func (r *testRepository) Delete(id string) error {
	_, err := r.c.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		log.Println("Unable to delete element of id", id, "\nError", err)
		return err
	}
	return nil
}
