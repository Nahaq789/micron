package dtos

type UserDto struct {
	userId     int    `json:"user_id"`
	uuidUserId string `json:"uuid_user_id"`
	username   string `json:"username"`
	email      string `json:"email"`
	roleId     int    `json:"role_id"`
	userTypeId int    `json:"user_type_id"`
	bio        string `json:"bio"`
}

func NewUserDto(userId int, uuidUserId string, username string, email string, roleId int, userTypeId int, bio string) *UserDto {
	return &UserDto{
		userId:     userId,
		uuidUserId: uuidUserId,
		username:   username,
		email:      email,
		roleId:     roleId,
		userTypeId: userTypeId,
		bio:        bio,
	}
}
func (u *UserDto) GetUserId() int {
	return u.userId
}
func (u *UserDto) GetUuidUserId() string {
	return u.uuidUserId
}
func (u *UserDto) GetUsername() string {
	return u.username
}
func (u *UserDto) GetEmail() string {
	return u.email
}
func (u *UserDto) GetRoleId() int {
	return u.roleId
}
func (u *UserDto) GetUserTypeId() int {
	return u.userTypeId
}
func (u *UserDto) GetBio() string {
	return u.bio
}
