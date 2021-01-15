package repository

type Repository interface {
	Close()
	FinAll()
	FindByID()
}

type Ride struct {
	ID       int
	Link     string
	Name     string
	Date     string
	Distance string
	Time     string
	AvgSpeed string
	Route    string
}
