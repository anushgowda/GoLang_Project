package services

import (
	"context"
	"time"
	// "fmt"

	"github.com/anushgowda/GoLang_Project/entities"
	"github.com/anushgowda/GoLang_Project/interfaces"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	Product *mongo.Collection
}

// Addproduct implements interfaces.IProduct.
// func (*ProductService) Addproduct(p *entities.Product) error {
// 	panic("unimplemented")
// }

// GetProductById implements interfaces.IProduct.
func (*ProductService) GetProductById(id primitive.ObjectID) (*entities.Product, error) {
	panic("unimplemented")
}

// SearchProducts implements interfaces.IProduct.
func (*ProductService) SearchProducts(name string) (*entities.Product, error) {
	panic("unimplemented")
}

func InitProductService(collection *mongo.Collection) interfaces.IProduct {

	return &ProductService{Product: collection}
}

func (p *ProductService) Addproduct(product *entities.Product) (string, error) {
	product.ID = primitive.NewObjectID()
	product.CreatedAt = time.Now()
	product.UpdatedAt = product.CreatedAt
	_, err := p.Product.InsertOne(context.Background(), product)
	if err != nil {
		return "", err
	} else {
		return "Record Inserted Successfully", nil
	}
}
