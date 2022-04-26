package usersmongorepo

import (
	"context"
	"log"
	"os"
	"root/internal/core/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type usersRepository struct {
	c *mongo.Collection
}

func NewUsersRepository(conn *conn) *usersRepository {
	return &usersRepository{c: conn.DB().Collection(os.Getenv("DB_COLLECTION"))}
}

func (r *usersRepository) Save(t *domain.User) error {
	_, err := r.c.InsertOne(context.TODO(), t)
	if err != nil {
		log.Println("Error in Repository -> Save()", err)
		return err
	}
	return nil
}

func (r *usersRepository) Update(t *domain.User) error {
	_, err := r.c.UpdateByID(context.TODO(), t.Id, bson.M{"$set": t})
	if err != nil {
		log.Println("Error in Repository -> Update()", err)
		return err
	}
	return nil
}

func (r *usersRepository) GetAll() (t []*domain.User, err error) {
	cur, err := r.c.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println("Error in Repository -> ShowAll()", err)
		return []*domain.User{}, err
	}

	var results []*domain.User
	for cur.Next(context.TODO()) {
		var elem domain.User
		err := cur.Decode(&elem)
		if err != nil {
			log.Println("Error in Repository -> ShowAll()", err)
			return []*domain.User{}, err
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Println("Error in Repository -> GetAll()", err)
		return []*domain.User{}, err
	}

	cur.Close(context.TODO())

	return results, nil
}

func (r *usersRepository) GetById(id string) (t *domain.User, err error) {
	findResult := r.c.FindOne(context.TODO(), bson.M{"_id": id})
	decodeErr := findResult.Decode(&t)
	if decodeErr != nil {
		log.Println("Error in Repository -> GetById()", decodeErr)
		return &domain.User{}, decodeErr
	}

	return t, nil
}

func (r *usersRepository) GetByEmail(email string) (t *domain.User, err error) {
	findResult := r.c.FindOne(context.TODO(), bson.M{"email": email})
	decodeErr := findResult.Decode(&t)
	if decodeErr != nil {
		log.Println("Error in Repository -> GetByEmail()", decodeErr)
		return &domain.User{}, decodeErr
	}

	return t, nil
}

func (r *usersRepository) Delete(id string) error {
	_, err := r.c.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Println("Unable to delete element of id", id, "\nError", err)
		return err
	}
	return nil
}
