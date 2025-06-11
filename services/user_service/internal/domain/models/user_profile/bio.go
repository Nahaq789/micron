package userprofile

type Bio struct {
	value string
}

func NewBio(v string) Bio {
	return Bio{value: v}
}
