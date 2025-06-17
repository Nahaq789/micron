package commands

type RegisterAdminUserCommand struct {
	email    string
	userName string
	bio      string
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
