package commands

type EditProfileCommand struct {
	userId   int
	userName string
	bio      string
}

func NewEditProfileCommand(userId int, userName string, bio string) EditProfileCommand {
	return EditProfileCommand{
		userId:   userId,
		userName: userName,
		bio:      bio,
	}
}

func (c EditProfileCommand) GetUserId() int {
	return c.userId
}

func (c EditProfileCommand) GetUserName() string {
	return c.userName
}

func (c EditProfileCommand) GetBio() string {
	return c.bio
}
