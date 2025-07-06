package usercontroller

import (
	"context"
	"user_service/internal/application/commands"
	"user_service/internal/application/use_cases/user"
	"user_service/proto"
)

type UserController struct {
	register user.RegisterAdminUser
	proto.UnimplementedUserServiceServer
}

func NewUserController(register user.RegisterAdminUser) *UserController {
	return &UserController{
		register: register,
	}
}

func (uc *UserController) CreateAdminUser(ctx context.Context, req *proto.CreateAdminUserRequest) (*proto.CreateAdminUserResponse, error) {
	command := commands.NewRegisterAdminUserCommand(req.GetEmail(), req.GetUsername(), req.GetBio(), req.GetOrganizationId())
	if err := uc.register.RegisterAdmin(ctx, command); err != nil {
		return nil, err
	}

	res := &proto.CreateAdminUserResponse{
		Message: "管理者ユーザの作成に成功しました。",
	}

	return res, nil
}
