package routes

import (
	"github.com/anushgowda/GoLang_Project/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, a controllers.AuthController) {
	user := r.Group("/api/user")
	user.POST("/register", a.Register)
	user.POST("/login", a.Login)
	user.POST("/logout", a.Logout)
}

func ProductRoutes(r *gin.Engine, p controllers.ProductController) {
	product := r.Group("/api/product") //localhost:4000/api/product/
	product.POST("/insert", p.Addproduct)
	product.GET("/getproduct/:id", p.GetProductById)
	product.GET("/getproducts/:name", p.SearchProducts)
}