package entity

type Bike struct {
	Name  string
	Type  string
	Model string
}

func (bike Bike) CanDrive() bool {
	return true
}
