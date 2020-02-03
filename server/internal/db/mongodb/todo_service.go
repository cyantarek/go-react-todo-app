package mongodb

import (
	"backend/internal/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoService struct {
	todo *mongo.Collection
}

func (h *TodoService) GetAll() []domain.Todo {
	cur, err := h.todo.Find(context.Background(), primitive.D{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	
	var out []domain.Todo
	var t domain.Todo
	for cur.Next(context.Background()) {
		err := cur.Decode(&t)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		
		out = append(out, t)
	}
	
	if cur.Err() != nil {
		fmt.Println(err)
		return nil
	}
	
	return out
}

func (h *TodoService) InsertOne(t domain.Todo) error {
	result, err := h.todo.InsertOne(context.Background(), t)
	if err != nil {
		return err
	}
	
	fmt.Println(result.InsertedID)
	
	return nil
}

func (h *TodoService) Update(filter, update primitive.M) error {
	result, err := h.todo.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	
	fmt.Println(result.ModifiedCount)
	
	return nil
}

func (h *TodoService) DeleteOne(id primitive.ObjectID) error {
	result, err := h.todo.DeleteOne(context.Background(), primitive.M{"_id": id})
	if err != nil {
		return err
	}
	
	fmt.Println(result.DeletedCount)
	
	return nil
}

func (h *TodoService) DeleteAll() error {
	result, err := h.todo.DeleteMany(context.Background(), primitive.D{})
	if err != nil {
		return err
	}
	
	fmt.Println(result.DeletedCount)
	
	return nil
}

func NewTodoService(coll *mongo.Collection) *TodoService {
	return &TodoService{todo: coll}
}
