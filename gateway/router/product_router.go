package router

import (
	"github.com/SomeHowMicroservice/shm-be/gateway/config"
	"github.com/SomeHowMicroservice/shm-be/gateway/handler"
	"github.com/SomeHowMicroservice/shm-be/gateway/middleware"
	"github.com/SomeHowMicroservice/shm-be/services/user/model"
	userpb "github.com/SomeHowMicroservice/shm-be/services/user/protobuf"
	"github.com/gin-gonic/gin"
)

func ProductRouter(rg *gin.RouterGroup, cfg *config.AppConfig, userClient userpb.UserServiceClient, productHandler *handler.ProductHandler) {
	accessName := cfg.Jwt.AccessName
	secretKey := cfg.Jwt.SecretKey

	product := rg.Group("/products")
	{
		product.GET("/:slug", productHandler.ProductDetails)
	}

	category := rg.Group("/categories")
	{
		category.GET("/tree", productHandler.GetCategoryTree)
	}

	admin := rg.Group("/admin", middleware.RequireAuth(accessName, secretKey, userClient), middleware.RequireMultiRoles([]string{model.RoleAdmin}))
	{
		admin.POST("/categories", productHandler.CreateCategory)
		admin.POST("/products", productHandler.CreateProduct)
	}
}
