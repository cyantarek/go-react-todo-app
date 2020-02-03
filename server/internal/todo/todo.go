package todo

import (
	"backend/internal/domain"
	"github.com/labstack/echo/v4"
)

type todo struct {
	todoRepo domain.TodoRepository
	//
	// other things related to this handler
}

func RegisterHandlers(todoRepo domain.TodoRepository, r *echo.Echo) {
	h := newTodo(todoRepo)
	
	h.initRoutes(r)
}

func newTodo(todoRepo domain.TodoRepository) *todo {
	return &todo{todoRepo:todoRepo}
}
