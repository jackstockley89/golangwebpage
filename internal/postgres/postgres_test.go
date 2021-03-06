package postgres

import (
	"database/sql"
	"log"
	"testing"

	r "github.com/jackstockley89/golangwebpage/internal"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var u = &r.Ride{
	ID:       uuid.New().String(),
	Link:     "Test Link",
	Name:     "Test Name",
	Date:     "01/01/2021",
	Distance: "100",
	Time:     "05:00:00",
	AvgSpeed: "15",
	Route:    "Test Route",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestFindByID(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	sql := "SELECT id, link, name, date, distance, time, avgspeed, route FROM rides_table WHERE id = \\$1"

	rows := sqlmock.NewRows([]string{"id", "link", "name", "date", "distance", "time", "avgspeed", "route"}).
		AddRow(u.ID, u.Link, u.Name, u.Date, u.Distance, u.Time, u.AvgSpeed, u.Route)

	mock.ExpectQuery(sql).WithArgs(u.ID).WillReturnRows(rows)

	ride, err := repo.FindByID(u.ID)
	assert.NotNil(t, ride)
	assert.NoError(t, err)
}

func TestFindByIDError(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	sql := "SELECT id, link, name, date, distance, time, avgspeed, route FROM rides_table WHERE id = \\$1"

	rows := sqlmock.NewRows([]string{"id", "link", "name", "date", "distance", "time", "avgspeed", "route"})

	mock.ExpectQuery(sql).WithArgs(u.ID).WillReturnRows(rows)

	ride, err := repo.FindByID(u.ID)
	assert.Empty(t, ride)
	assert.Error(t, err)
}
