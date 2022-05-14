package sqlstorage

import (
	"context"
	"fmt"
	"time"

	"github.com/fenogentov/OTUS-HW-Go/hw12_13_14_15_calendar/internal/logger"
	"github.com/fenogentov/OTUS-HW-Go/hw12_13_14_15_calendar/internal/storage"

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
type Storage struct {
	db     *sqlx.DB
	config config
	logger *logger.Logger
}

// New ...
func New(logger *logger.Logger, host, port, baseName, user, password string) *Storage {
	return &Storage{
		config: config{
			host:     host,
			port:     port,
			baseName: baseName,
			user:     user,
			password: password,
		},
		logger: logger,
	}
}

// Connect ...
func (s *Storage) Connect(ctx context.Context) error {
	// "postgres://username:password@localhost:5432/database_name"
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		s.config.user, s.config.password, s.config.host, s.config.port, s.config.baseName)
	// Create connection pool
	db, err := sqlx.Connect("pgx", connString)
	if err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close ...
func (s *Storage) Close(ctx context.Context) error {
	// TODO
	return nil
}

// CreateEvent ...
func (s *Storage) CreateEvent(evnt storage.Event) {
	if (evnt == storage.Event{}) {
		return
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
		s.logger.Error("error create event")
	}
}

// UpdateEvent ...
func (s *Storage) UpdateEvent(evnt storage.Event) {
	if (evnt == storage.Event{}) {
		return
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
		s.logger.Error("error update event")
	}
}

// DeleteEvent ...
func (s *Storage) DeleteEvent(evnt storage.Event) {
	if (evnt == storage.Event{}) {
		return
	}

	_, err := s.db.Exec(
		`DELETE FROM events WHERE id=$1`,
		evnt.ID,
	)
	if err != nil {
		s.logger.Error("error delete event")
	}
}

// GetEvents ...
func (s *Storage) GetEvents(startDT, endDT time.Time) []storage.Event {
	rows, err := s.db.Query(`SELECT * FROM events
			WHERE startTime >=$1 AND startTime <=$2`,
		startDT,
		endDT,
	)
	if err != nil {
		s.logger.Error("error GetEvents")
	}
	defer rows.Close()

	var events []storage.Event
	for rows.Next() {
		e := storage.Event{}
		err = rows.Scan(&e.ID, &e.Title, &e.StartTime, &e.EndTime, &e.Descript, &e.UserID)
		if err != nil {
			return nil
		}
		events = append(events, e)
	}

	return events
}
