package services

import (
	"context"
	"time"

	"github.com/anushgowda/GoLang_Project/entities"
	"github.com/anushgowda/GoLang_Project/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	Product *mongo.Collection
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

func (p *ProductService) GetProductById(id primitive.ObjectID) (*entities.Product, error) {

	ctx := context.Background()
	var product entities.Product
	err := p.Product.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
