package role

type Admin struct {
	roleId int
}

func DetermineAdminRole() Admin {
	return Admin{roleId: 1}
}

func (a Admin) GetRoleId() int {
	return a.roleId
}
