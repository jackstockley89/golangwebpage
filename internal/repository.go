package repository

//Repository interface
type Repository interface {
	Close()
	FindByID(id string) (*Ride, error)
}

//Ride struct
type Ride struct {
	ID       string
	Link     string
	Name     string
	Date     string
	Distance string
	Time     string
	AvgSpeed string
	Route    string
}
