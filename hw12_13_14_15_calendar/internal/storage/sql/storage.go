package sqlstorage

import (
	"context"
	"fmt"
	"time"

	"hw12_13_14_15_calendar/internal/logger"
	"hw12_13_14_15_calendar/internal/storage"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

// Storage ...
type Storage struct { // TODO
	db     *sqlx.DB
	logger *logger.Logger
}

// New ...
func New(logger *logger.Logger, user, password, host, port, nameDB string) (*Storage, error) {
	// "postgres://username:password@localhost:5432/database_name"
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user, password, host, port, nameDB)

	logger.Debug(connString)
	// Create connection pool
	db, err := sqlx.Connect("pgx", connString)
	if err != nil {
		logger.Error("error connect DB")
		return nil, err
	}

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
func (s *Storage) CreateEvent(evnt storage.Event) error {
	_, err := s.db.Exec(
		`INSERT INTO events (id, title, startTime, endTime, descript, userid) VALUES ($1, $2, $3, $4, $5, $6)`,
		evnt.ID,
		evnt.Title,
		evnt.StartTime,
		evnt.EndTime,
		evnt.Descript,
		evnt.UserID,
	)
	if err != nil {
		return err
	}
	return nil
}

// UpdateEvent ...
func (s *Storage) UpdateEvent(evnt storage.Event) error {
	_, err := s.db.Exec(
		`UPDATE events SET (title, startTime, endTime, descript, userid) = (
			$1, $2, $3, $4, $5) WHERE id=$6`,
		evnt.Title,
		evnt.StartTime,
		evnt.EndTime,
		evnt.Descript,
		evnt.UserID,
		evnt.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteEvent ...
func (s *Storage) DeleteEvent(evnt storage.Event) error {
	_, err := s.db.Exec(
		`DELETE FROM events WHERE id=$1`,
		evnt.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetEvents ...
func (s *Storage) GetEvents(timeStart, timeEnd time.Time) ([]storage.Event, error) {
	rows, err := s.db.Queryx(`SELECT id, title, startTime, endTime, descript, userid
			FROM events
			WHERE startTime >=$1 AND startTime <=$2`,
		timeStart,
		timeEnd,
	)
	if err != nil {

		return nil, err
	}
	defer rows.Close()
	var events []storage.Event
	for rows.Next() {
		e := storage.Event{}
		if err := rows.StructScan(&e); err != nil {
			s.logger.Info(err.Error())
		}
		events = append(events, e)
	}
	return events, nil
}
