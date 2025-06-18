package usertype

type UserType interface {
	GetTypeName() string
	GetTypeId() int
}
