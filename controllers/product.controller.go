package controllers

import (
	"fmt"
	"net/http"

	"github.com/anushgowda/GoLang_Project/entities"
	"github.com/anushgowda/GoLang_Project/interfaces"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductController struct {
	ProductService interfaces.IProduct
}

func InitProductController(productSvc interfaces.IProduct) *ProductController {
	return &ProductController{ProductService: productSvc}
}

func (p ProductController) Addproduct(c *gin.Context) {
	fmt.Println("Invoked controller")
	var product entities.Product
	err := c.BindJSON(&product)
	if err != nil {
		fmt.Println("Not Binding")
		return
	}
	fmt.Printf("binding")
	result, err := p.ProductService.Addproduct(&product)
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
}

func (p ProductController) GetProductById(c *gin.Context) {

	productId := c.Param("id")
	pid, err := primitive.ObjectIDFromHex(productId)
	product, err := p.ProductService.GetProductById(pid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, product)
}

func (p ProductController) SearchProducts(c *gin.Context) {
	name := c.Param("Name")
	products, err := p.ProductService.SearchProducts(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, products)
}

