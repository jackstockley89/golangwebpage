package postgres

import (
	"context"
	"database/sql"
	"time"

	repo "github.com/jackstockley89/golangwebpage/repository"
)

// repository represent the repository model
type repository struct {
	db *sql.DB
}

// NewRepository will create a variable that represent the Repository struct
func NewRepository(dialect, dsn string, idleConn, maxConn int) (repo.Repository, error) {
	db, err := sql.Open(dialect, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(idleConn)
	db.SetMaxOpenConns(maxConn)

	return &repository{db}, nil
}

// Close attaches the provider and close the connection
func (r *repository) Close() {
	r.db.Close()
}

func (r *repository) FindByID(id string) (*repo.Ride, error) {
	ride := new(repo.Ride)
	sql := "SELECT id, link, name, date, distance, time, avgspeed, route FROM rides_table WHERE id = $1"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, sql, id).Scan(&ride.ID, &ride.Link, &ride.Name, &ride.Date, &ride.Distance, &ride.Time, &ride.AvgSpeed, &ride.Route)
	if err != nil {
		return nil, err
	}
	return ride, nil
}
