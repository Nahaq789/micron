package usertype

type Member struct {
	id     int
	name   string
	roleId int
}

func NewMember(roleId int) Member {
	return Member{id: 1, name: "Member", roleId: roleId}
}

func (m Member) DecideRole() int {
	return m.roleId
}
