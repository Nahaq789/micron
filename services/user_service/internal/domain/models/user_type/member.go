package usertype

type Member struct {
	id   int
	name string
}

func NewMember() Member {
	return Member{id: 1, name: "Member"}
}
