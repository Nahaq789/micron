package usercontroller

import (
	"context"
	"log/slog"
	"user_service/internal/application/commands"
	"user_service/internal/application/use_cases/user"
	"user_service/proto"
)

type UserController struct {
	logger   slog.Logger
	register user.RegisterAdminUser
	proto.UnimplementedUserServiceServer
}

func NewUserController(l slog.Logger, register user.RegisterAdminUser) *UserController {
	return &UserController{
		logger:   l,
		register: register,
	}
}

func (uc *UserController) CreateAdminUser(ctx context.Context, req *proto.CreateAdminUserRequest) (*proto.CreateAdminUserResponse, error) {
	uc.logger.InfoContext(ctx, "管理者ユーザ作成処理を開始します。")
	command := commands.NewRegisterAdminUserCommand(req.GetEmail(), req.GetUsername(), req.GetBio(), req.GetOrganizationId())
	if err := uc.register.RegisterAdmin(ctx, command); err != nil {
		return nil, err
	}

	res := &proto.CreateAdminUserResponse{
		Message: "管理者ユーザの作成に成功しました。",
	}

	uc.logger.InfoContext(ctx, "管理者ユーザ作成処理を終了します。")
	return res, nil
}
