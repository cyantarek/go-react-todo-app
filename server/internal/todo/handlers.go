package todo

import (
	"backend/internal/domain"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (t *todo) getAllTodo(c echo.Context) error {
	allTodo := t.todoRepo.GetAll()
	
	return c.JSON(200, allTodo)
}

func (t *todo) insertOneTodo(c echo.Context) error {
	var in domain.Todo
	
	err := c.Bind(&in)
	if err != nil {
		return err
	}
	
	in.ID = primitive.NewObjectID()
	
	err = t.todoRepo.InsertOne(in)
	if err != nil {
		return err
	}
	
	return c.JSON(200, in)
}

func (t *todo) todoComplete(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return err
	}
	
	err = t.todoRepo.Update(primitive.M{"_id": id}, primitive.M{"$set": primitive.M{"status": true}})
	if err != nil {
		return err
	}
	
	return c.JSON(200, echo.Map{"status": "completed"})
}

func (t *todo) todoUndo(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return err
	}
	
	err = t.todoRepo.Update(primitive.M{"_id": id}, primitive.M{"$set": primitive.M{"status": false}})
	if err != nil {
		return err
	}
	
	return c.JSON(200, echo.Map{"status": "undo done"})
}

func (t *todo) deleteTodo(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return err
	}
	
	err = t.todoRepo.DeleteOne(id)
	if err != nil {
		return err
	}
	
	return c.JSON(200, echo.Map{"status": "deleted"})
}

func (t *todo) deleteAllTodo(c echo.Context) error {
	err := t.todoRepo.DeleteAll()
	if err != nil {
		return err
	}
	
	return c.JSON(200, echo.Map{"status": "deleted all"})
}