package controllers

import (
	"fmt"
	"net/http"

	"github.com/anushgowda/GoLang_Project/entities"
	"github.com/anushgowda/GoLang_Project/interfaces"
	"github.com/gin-gonic/gin"
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

// func (p ProductController) GetProductById(c *gin.Context) {
// 	result, err := p.ProductService.GetProductById()
// 	if err != nil {
// 		return
// 	} else {
// 		c.IndentedJSON(http.StatusCreated, result)
// 	}
// }
