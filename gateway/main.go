package main

import (
	"fmt"
	"log"

	"github.com/SomeHowMicroservice/shm-be/gateway/config"
	"github.com/SomeHowMicroservice/shm-be/gateway/container"
	"github.com/SomeHowMicroservice/shm-be/gateway/initialization"
	"github.com/SomeHowMicroservice/shm-be/gateway/router"
	"github.com/gin-gonic/gin"
)

var (
	authAddr = "localhost:8081"
	userAddr = "localhost:8082"
	productAddr = "localhost:8083"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Tải cấu hình Gateway thất bại: %v", err)
	}

	authAddr = cfg.App.ServerHost + fmt.Sprintf(":%d", cfg.Services.AuthPort)
	userAddr = cfg.App.ServerHost + fmt.Sprintf(":%d", cfg.Services.UserPort)
	productAddr = cfg.App.ServerHost + fmt.Sprintf(":%d", cfg.Services.ProductPort)
	clients := initialization.InitClients(authAddr, userAddr, productAddr)

	appContainer := container.NewContainer(clients.AuthClient, clients.UserClient, clients.ProductClient, cfg)

	r := gin.Default()
	config.CORSConfig(r)

	api := r.Group("/api/v1")
	router.AuthRouter(api, cfg, clients.UserClient, appContainer.Auth.Handler)
	router.UserRouter(api, cfg, clients.UserClient, appContainer.User.Handler)
	router.ProductRouter(api, cfg, clients.UserClient, appContainer.Product.Handler)

	r.Run(fmt.Sprintf(":%d", cfg.App.GRPCPort))
}
