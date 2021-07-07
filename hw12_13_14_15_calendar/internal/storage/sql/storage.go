package sqlstorage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fenogentov/OTUS-HW-Go/hw12_13_14_15_calendar/internal/logger"
	"github.com/fenogentov/OTUS-HW-Go/hw12_13_14_15_calendar/internal/storage"

	// _ "github.com/jackc/pgx/stdlib" .
	"github.com/jmoiron/sqlx"
)

// Storage ...
type Storage struct { // TODO
	db     *sqlx.DB
	logger logger.Logger
}

// New ...
func New(logger logger.Logger, user, password, host, port, nameDB string) (*Storage, error) {
	// "postgres://username:password@localhost:5432/database_name"
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user, password, host, port, nameDB)

	// Create connection pool
	db, err := sqlx.Connect("pgx", connString)
	if err != nil {
		logger.Error("error connect DB")
	}

	rows, err := db.Query("SELECT * FROM events")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		fmt.Println(rows)
	}
	// sqlStatement := `INSERT INTO tst (isbn, title, author, price)
	//	VALUES ('qertu', 'jon@calhogu.oq', 'ontghanq', 55)`
	//	_, err = db.Exec(sqlStatement)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	return &Storage{db: db, logger: logger}, nil
}

// Connect ...
func (s *Storage) Connect(ctx context.Context) error {
	// TODO
	return nil
}

// Close ...
func (s *Storage) Close(ctx context.Context) error {
	// TODO
	return nil
}

// CreateEvent ...
func (s *Storage) CreateEvent(evnt storage.Event) {
	_, err := s.db.Exec(
		`INSERT INTO events (id, title, startTime, endTime, description, userid) VALUES ($1, $2, $3, $4, $5, $6)`,
		evnt.ID,
		evnt.Title,
		evnt.StartTime,
		evnt.EndTime,
		evnt.Description,
		evnt.UserID,
	)
	if err != nil {
		fmt.Println(err)
		s.logger.Error("error create event")
	}
}

// UpdateEvent ...
func (s *Storage) UpdateEvent(evnt storage.Event) {
	_, err := s.db.Exec(
		`UPDATE events SET (title, startTime, endTime, description, userid) = (
			$1, $2, $3, $4, $5) WHERE id=$6`,
		evnt.Title,
		evnt.StartTime,
		evnt.EndTime,
		evnt.Description,
		evnt.UserID,
		evnt.ID,
	)
	if err != nil {
		s.logger.Error("error update event")
	}
}

// DeleteEvent ...
func (s *Storage) DeleteEvent(evnt storage.Event) {
	_, err := s.db.Exec(
		`DELETE FROM events WHERE id=$1`,
		evnt.ID,
	)
	if err != nil {
		s.logger.Error("error delete event")
	}
}

// GetEvents ...
func (s *Storage) GetEvents(startDT, endDT time.Time) {
	result, err := s.db.Query(`SELECT id, title, startTime, endTime, description, userid
			FROM events
			WHERE startTime >=$1 AND startTime <=$2`,
		startDT,
		endDT,
	)
	if err != nil {
		s.logger.Error("error GetEvents")
	}
	defer result.Close()
}
