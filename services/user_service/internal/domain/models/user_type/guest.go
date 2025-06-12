package usertype

type Guest struct {
	id   int
	name string
}

func NewGuest() Guest {
	return Guest{id: 2, name: "Guest"}
}

func (g Guest) DecideRole() int {
	return 1
}
