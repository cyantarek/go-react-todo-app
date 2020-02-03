package main

import (
	"backend/internal/db/mongodb"
	"backend/internal/todo"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var (
	MongoUrl = ""
	DbName = "todo_db"
	helloCollection = "todo"
	Port = "5000"
)

func main() {
	if os.Getenv("PORT") == "" {
		Port = "5000"
	} else {
		Port = os.Getenv("PORT")
	}
	
	app := echo.New()
	
	app.Use(middleware.Logger())
	app.Use(middleware.CORS())
	
	if os.Getenv("MONGO_URL") != "" {
		MongoUrl = os.Getenv("MONGO_URL")
	}
	
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoUrl))
	if err != nil {
		log.Fatal(err)
	}
	
	todoDb := mongodb.NewTodoService(client.Database(DbName).Collection(helloCollection))
	todo.RegisterHandlers(todoDb, app)
	
	log.Fatal(app.Start(fmt.Sprintf(":%s", Port)))
}
