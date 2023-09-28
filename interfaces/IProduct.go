package interfaces

import (
	"github.com/anushgowda/GoLang_Project/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IProduct interface {
	Addproduct(p *entities.Product) (string, error)
	GetProductById(id primitive.ObjectID) (*entities.Product, error)
	SearchProducts(name string) ([]*entities.Product, error)
}
