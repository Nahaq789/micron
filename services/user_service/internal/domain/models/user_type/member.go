package usertype

type Member struct {
	id   int
	name string
}

func NewMember() Member {
	return Member{id: 1, name: "Member"}
}

func (m Member) GetTypeName() string {
	return m.name
}

func (m Member) GetTypeId() int {
	return m.id
}
