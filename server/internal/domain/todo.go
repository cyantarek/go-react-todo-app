package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID     primitive.ObjectID `bson:"_id" json:"id"`
	Task   string             `json:"task"`
	Status bool               `json:"status"`
}

type TodoRepository interface {
	GetAll() []Todo
	InsertOne(t Todo) error
	Update(filter, update primitive.M) error
	DeleteOne(id primitive.ObjectID) error
	DeleteAll() error
}
