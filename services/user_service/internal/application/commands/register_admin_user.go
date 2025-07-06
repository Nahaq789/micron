package commands

type RegisterAdminUserCommand struct {
	email          string
	userName       string
	bio            string
	organizationId string
}

func NewRegisterAdminUserCommand(email, userName, bio, organizaqtionId string) RegisterAdminUserCommand {
	return RegisterAdminUserCommand{
		email:          email,
		userName:       userName,
		bio:            bio,
		organizationId: organizaqtionId,
	}
}

func (r RegisterAdminUserCommand) GetEmail() string {
	return r.email
}

func (r RegisterAdminUserCommand) GetUserName() string {
	return r.userName
}

func (r RegisterAdminUserCommand) GetBio() string {
	return r.bio
}

func (r RegisterAdminUserCommand) GetOrganizationId() string {
	return r.organizationId
}
