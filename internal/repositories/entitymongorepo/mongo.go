package entitymongorepo

import (
	"context"
	"entity/internal/core/domain"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type entityRepository struct {
	c *mongo.Collection
}

const EntityCollection = "collection_name_here"

func NewEntityRepository(conn *conn) *entityRepository {
	return &entityRepository{c: conn.DB().Collection(EntityCollection)}
}

func (r *entityRepository) Insert(t *domain.Entity) error {
	_, err := r.c.InsertOne(context.TODO(), t)
	if err != nil {
		log.Println("Error in Repository -> Create()", err)
		return err
	}
	return nil
}

func (r *entityRepository) Set(t *domain.Entity) error {
	_, err := r.c.UpdateByID(context.TODO(), t.Id, bson.M{"$set": t})
	if err != nil {
		log.Println("Error in Repository -> Update()", err)
		return err
	}
	return nil
}

func (r *entityRepository) SelectAll() (t []*domain.Entity, err error) {
	cur, err := r.c.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println("Error in Repository -> ShowAll()", err)
		return []*domain.Entity{}, err
	}

	var results []*domain.Entity
	for cur.Next(context.TODO()) {
		var elem domain.Entity
		err := cur.Decode(&elem)
		if err != nil {
			log.Println("Error in Repository -> ShowAll()", err)
			return []*domain.Entity{}, err
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Println("Error in Repository -> ShowAll()", err)
		return []*domain.Entity{}, err
	}

	cur.Close(context.TODO())

	return results, nil
}

func (r *entityRepository) SelectById(id string) (t *domain.Entity, err error) {
	findResult := r.c.FindOne(context.TODO(), bson.M{"_id": id})
	decodeErr := findResult.Decode(&t)
	if decodeErr != nil {
		log.Println("Error in Repository -> ShowById()", decodeErr)
		return &domain.Entity{}, decodeErr
	}

	return t, nil
}

func (r *entityRepository) Delete(id string) error {
	_, err := r.c.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Println("Unable to delete element of id", id, "\nError", err)
		return err
	}
	return nil
}
