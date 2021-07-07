package entity

type Bus struct {
	Name  string
	Type  string
	Model string
}

func (bus Bus) CanDrive() bool {
	return false
}
