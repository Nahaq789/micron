package aggregate

type Organization struct {
	organizationId   int
	organizationName string
}

func NewOrganization(organizationId int, organizationName string) Organization {
	return Organization{
		organizationId:   organizationId,
		organizationName: organizationName,
	}
}
