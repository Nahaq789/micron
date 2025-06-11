package user

type UserType int

const (
	Member UserType = iota
	Guest
)

var userTypeName = map[UserType]string{
	Member: "Member",
	Guest:  "Guest",
}

var userTypeKey = map[UserType]int{
	Member: 1,
	Guest:  2,
}

func (ut UserType) String() string {
	return userTypeName[ut]
}

func (ut UserType) Int() int {
	return userTypeKey[ut]
}
