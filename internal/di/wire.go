//go:build wireinject
// +build wireinject

package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/yashkadam007/brewkar/internal/config"
	"github.com/yashkadam007/brewkar/internal/controller"
	"github.com/yashkadam007/brewkar/internal/repository"
	"github.com/yashkadam007/brewkar/internal/router"
	"github.com/yashkadam007/brewkar/internal/service"
)

var infraSet = wire.NewSet(
	ProvideConfig,
	ProvideLogger,
	ProvideDatabase,
)

var repoSet = wire.NewSet(
	repository.NewUserRepository,
)

var serviceSet = wire.NewSet(
	wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)),
	provideAuthService,
)

var controllerSet = wire.NewSet(
	controller.NewAuthController,
)

// InitializeApp initializes the complete application
func InitializeApp() (*gin.Engine, error) {
	panic(wire.Build(
		infraSet,
		repoSet,
		serviceSet,
		controllerSet,
		router.SetupRouter,
	))
}

// Provider functions
func provideAuthService(userRepo repository.UserRepository, cfg *config.Config) *service.AuthServiceImpl {
	return service.NewAuthService(userRepo, cfg.JWT).(*service.AuthServiceImpl)
}
