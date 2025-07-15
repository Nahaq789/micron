package di

import (
	"database/sql"
	"log/slog"
	"user_service/internal/domain/repositories"
	"user_service/internal/domain/services"
	"user_service/internal/infrastructure"
	usercontroller "user_service/internal/presentations/user_controller"
	userprofile "user_service/internal/presentations/user_profile"

	"github.com/google/wire"
)

func ProviderUserRepository(logger *slog.Logger, db *sql.DB) *infrastructure.UserRepositoryImpl {
	repository := infrastructure.NewUserRepositoryImpl(logger, db)
	return repository
}

var userRepositorySet = wire.NewSet(
	ProviderUserRepository,
	wire.Bind(new(repositories.UserRepository), new(*infrastructure.UserRepositoryImpl)),
)

var emailDuplicateCheckerSet = wire.NewSet(
	services.NewEmailDuplicateService,
	wire.Bind(new(services.EmailDuplicateChecker), new(*services.EmailDuplicateService)),
)

var controllerSet = wire.NewSet(usercontroller.NewUserController)

type ControllerSet struct {
	UserController            *usercontroller.UserController
	EditUserProfileController *userprofile.EditUserProfileController
}

func Initialize(logger *slog.Logger, db *sql.DB) *ControllerSet {
	wire.Build(
		userRepositorySet,
		emailDuplicateCheckerSet,
		wire.Struct(new(ControllerSet), "*"),
	)
	return nil
}
