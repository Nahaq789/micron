package commands

type GetUserCommand struct {
	UserId int
}

func NewGetUserCommand(userId int) GetUserCommand {
	return GetUserCommand{
		UserId: userId,
	}
}

func (c GetUserCommand) GetUserId() int {
	return c.UserId
}
