package sqlstorage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"hw12_13_14_15_calendar/internal/logger"
	"hw12_13_14_15_calendar/internal/storage"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type config struct {
	host     string
	port     string
	baseName string
	user     string
	password string
}

// Storage ...
type StorageDB struct {
	db     *sqlx.DB
	config config
	logger *logger.Logger
}

// New ...
func New(logger *logger.Logger, host, port, baseName, user, password string) (*StorageDB, error) {

	// "postgres://username:password@localhost:5432/database_name"
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user, password, host, port, baseName)
	// Create connection pool
	db, err := sqlx.Connect("pgx", connString)
	if err != nil {
		return nil, err
	}

	storage := &StorageDB{
		db:     db,
		config: config{host: host, port: port, baseName: baseName, user: user, password: password},
		logger: logger,
	}

	return storage, nil
}

// Connect ...
func (s *StorageDB) Connect(ctx context.Context) error {
	return nil
}

// Close ...
func (s *StorageDB) Close(ctx context.Context) error {
	return nil
}

// CreateEvent ...
func (s *StorageDB) CreateEvent(evnt storage.Event) error {
	if (evnt == storage.Event{}) {
		return errors.New("error CreateEvent: empty event")
	}

	_, err := s.db.Exec(
		`INSERT INTO events (id, title, starttime, endtime, descript, userid) VALUES ($1, $2, $3, $4, $5, $6)`,
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
func (s *StorageDB) UpdateEvent(evnt storage.Event) error {
	if (evnt == storage.Event{}) {
		return errors.New("error UpdateEvent: empty event")
	}

	_, err := s.db.Exec(
		`UPDATE events SET (title, starttime, endtime, descript, userid) = (
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
func (s *StorageDB) DeleteEvent(evnt storage.Event) error {
	if (evnt == storage.Event{}) {
		return errors.New("error DeleteEvent: empty event")
	}

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
func (s *StorageDB) GetEvents(startDT, endDT time.Time) ([]storage.Event, error) {
	rows, err := s.db.Query(`SELECT * FROM events
			WHERE startTime >=$1 AND startTime <=$2`,
		startDT,
		endDT,
	)
	if err != nil {
		s.logger.Error("error getevebt event")
	}
	defer rows.Close()

	var events []storage.Event
	for rows.Next() {
		e := storage.Event{}
		err = rows.Scan(&e.ID, &e.Title, &e.StartTime, &e.EndTime, &e.Descript, &e.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	return events, nil
}
