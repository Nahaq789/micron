package usercontroller

import (
	"context"
	"log/slog"
	"net/http"
	"user_service/internal/application/commands"
	"user_service/internal/application/dtos"
	"user_service/internal/application/use_cases/user"
	"user_service/proto"
)

type UserController struct {
	logger   *slog.Logger
	register *user.RegisterAdminUser
	get      *user.GetUserById
	proto.UnimplementedUserServiceServer
}

func NewUserController(l *slog.Logger, register *user.RegisterAdminUser, get *user.GetUserById) *UserController {
	return &UserController{
		logger:   l,
		register: register,
		get:      get,
	}
}

func (uc *UserController) CreateAdminUser(ctx context.Context, req *proto.CreateAdminUserRequest) (*proto.CreateAdminUserResponse, error) {
	uc.logger.InfoContext(ctx, "管理者ユーザ作成処理を開始します。")
	command := commands.NewRegisterAdminUserCommand(req.GetEmail(), req.GetUsername(), req.GetBio(), req.GetOrganizationId())
	if err := uc.register.RegisterAdmin(ctx, command); err != nil {
		uc.logger.ErrorContext(ctx, "管理者ユーザ作成に失敗しました。")
		res := &proto.CreateAdminUserResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "管理者ユーザ作成に失敗しました。",
		}
		return res, nil
	}

	res := &proto.CreateAdminUserResponse{
		StatusCode: http.StatusCreated,
		Message:    "管理者ユーザの作成に成功しました。",
	}

	uc.logger.InfoContext(ctx, "管理者ユーザ作成処理を終了します。")
	return res, nil
}

func (uc *UserController) GetUserById(ctx context.Context, id string) (*dtos.UserDto, error) {
	uc.logger.InfoContext(ctx, "ユーザ情報取得処理を開始します。")
	return nil, nil
}
