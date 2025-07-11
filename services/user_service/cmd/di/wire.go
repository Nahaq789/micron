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

// var registerAdminUserSet = wire.NewSet(
// 	user.NewRegisterAdminUser,
// 	wire.Bind(new(user.RegisterAdminUser), new())
// )

func ProviderUserRepository(logger *slog.Logger, db *sql.DB) *infrastructure.UserRepositoryImpl {
	repository := infrastructure.NewUserRepositoryImpl(logger, db)
	return repository
}

var userRepositorySet = wire.NewSet(
	ProviderUserRepository,
	wire.Bind(new(repositories.UserRepository), new(*infrastructure.UserRepositoryImpl)),
)

var emailDuplicateChecker = wire.NewSet(
	services.NewEmailDuplicateService,
	wire.Bind(new(services.EmailDuplicateChecker), new(*services.EmailDuplicateService)),
)

var controllerSet = wire.NewSet(usercontroller.NewUserController)

type ControllerSet struct {
	UserController            *usercontroller.UserController
	EditUserProfileController *userprofile.EditUserProfileController
}
