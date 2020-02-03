package mongodb

import (
	"backend/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type HelloService struct {
	hello *mongo.Collection
}

func (h *HelloService) Get() domain.HelloMessage {
	return domain.HelloMessage{Text: "Hi"}
}

func NewHelloService(coll *mongo.Collection) *HelloService {
	return &HelloService{hello: coll}
}
