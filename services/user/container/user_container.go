package container

import (
	"github.com/SomeHowMicroservice/shm-be/services/user/handler"
	measurementRepo "github.com/SomeHowMicroservice/shm-be/services/user/repository/measurement"
	profileRepo "github.com/SomeHowMicroservice/shm-be/services/user/repository/profile"
	roleRepo "github.com/SomeHowMicroservice/shm-be/services/user/repository/role"
	userRepo "github.com/SomeHowMicroservice/shm-be/services/user/repository/user"
	"github.com/SomeHowMicroservice/shm-be/services/user/service"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Container struct {
	GRPCHandler *handler.GRPCHandler
}

func NewContainer(db *gorm.DB, grpcServer *grpc.Server) *Container {
	userRepo := userRepo.NewUserRepository(db)
	profileRepo := profileRepo.NewProfileRepository(db)
	measurementRepo := measurementRepo.NewMeasurementRepository(db)
	roleRepo := roleRepo.NewRoleRepository(db)
	svc := service.NewUserService(userRepo, roleRepo, profileRepo, measurementRepo)
	hdl := handler.NewGRPCHandler(grpcServer, svc)
	return &Container{
		GRPCHandler: hdl,
	}
}
