package todo

import "github.com/labstack/echo/v4"

func (t *todo) initRoutes(r *echo.Echo) {
	api := r.Group("/api")
	
	api.GET("/todo", t.getAllTodo)
	api.POST("/todo", t.insertOneTodo)
	api.PUT("/todo/:id/complete", t.todoComplete)
	api.PUT("/todo/:id/undo", t.todoUndo)
	api.DELETE("/todo/:id", t.deleteTodo)
	api.DELETE("/todo", t.deleteAllTodo)
}
