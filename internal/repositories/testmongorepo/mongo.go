package testmongorepo

import (
	"context"
	"log"
	"test/internal/core/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type testRepository struct {
	c *mongo.Collection
}

const TestCollection = "testing"

func NewTestRepository(conn Connection) *testRepository {
	return &testRepository{c: conn.DB().Collection(TestCollection)}
}

func (r *testRepository) Create(t *domain.Test) error {
	insertResult, err := r.c.InsertOne(context.TODO(), t)
	log.Println(insertResult)
	if err != nil {
		log.Fatal("Error in Repository -> Create()", err)
	}
	return err
}

func (r *testRepository) Update(t *domain.Test) error {
	updateResult, err := r.c.UpdateByID(context.TODO(), t.Id, t)
	log.Println(updateResult)
	if err != nil {
		log.Fatal("Error in Repository -> Update()", err)
	}
	return err
}

func (r *testRepository) ShowAll() (t []*domain.Test, err error) {
	cur, err := r.c.Find(context.TODO(), bson.D{})
	log.Println(cur)
	if err != nil {
		log.Fatal("Error in Repository -> ShowAll()", err)
	}

	var results []*domain.Test
	for cur.Next(context.TODO()) {
		var elem domain.Test
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return results, err
}

func (r *testRepository) ShowById(id string) (t *domain.Test, err error) {
	findResult := r.c.FindOne(context.TODO(), bson.M{"id": uuid.MustParse(id)})
	log.Println(findResult)

	decodeErr := findResult.Decode(&t)
	if decodeErr != nil {
		log.Fatal("Error in Repository -> ShowById()", decodeErr)
	}

	return t, decodeErr
}

func (r *testRepository) Delete(id string) error {
	deleteResult, err := r.c.DeleteOne(context.TODO(), bson.M{"id": uuid.MustParse(id)})
	log.Println(deleteResult)
	if err != nil {
		log.Fatal("Unable to delete element of id", id, "\nError", err)
	}
	return err
}
