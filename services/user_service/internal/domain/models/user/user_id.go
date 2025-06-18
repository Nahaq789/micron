package user

type UserId struct {
	value int
}

func NewUserId(value int) UserId {
	return UserId{value: value}
}

func Init() UserId {
	return UserId{value: 0}
}

func (u UserId) GetValue() int {
	return u.value
}
