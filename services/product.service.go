package services

import (
	"context"
	"time"
	"fmt"
	"github.com/anushgowda/GoLang_Project/entities"
	"github.com/anushgowda/GoLang_Project/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	Product *mongo.Collection
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


func (prod *ProductService) SearchProducts(name string) ([]*entities.Product, error) {
	var products []*entities.Product
	cursor, err := prod.Product.Find(context.TODO(), bson.M{"Name": name})
	if err != nil {
		return nil, err
	} else {
		fmt.Println(cursor)
		for cursor.Next(context.TODO()) {
			product := &entities.Product{}
			err := cursor.Decode(product)

			if err != nil {
				return nil, err
			}
			products = append(products, product)
		}
		if err := cursor.Err(); err != nil {
			return nil, err
		}
		if len(products) == 0 {
			return []*entities.Product{}, nil
		}
		return products, nil
	}
}