package entity

type Car struct {
	Name  string
	Type  string
	Model string
}

func (car Car) CanDrive() bool {
	return true
}
