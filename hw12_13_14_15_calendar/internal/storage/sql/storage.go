package sqlstorage

import (
	"context"
	"fmt"
	"io"
	"time"

	"hw12_13_14_15_calendar/internal/storage"
	"hw12_13_14_15_calendar/internal/util/logger"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// type config struct {
// 	host     string
// 	port     string
// 	baseName string
// 	user     string
// 	password string
// }

// Storage ...
type Storage struct {
	pgpool *pgxpool.Pool
	logger *logger.Logger
}

// New ...
func New(logger *logger.Logger, host, port, baseName, user, password string) (storage.Storage, error) {
	// "postgres://username:password@localhost:5432/database_name"
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, baseName)
	pgp, err := newPool(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("failed create sql storage: %w", err)
	}

	storage := &Storage{
		pgpool: pgp,
		logger: logger,
	}

	return storage, nil
}

func newPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	config.ConnConfig.RuntimeParams["timezone"] = "UTC"
	config.ConnConfig.RuntimeParams["application_name"] = "Calendar"

	p, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("pg.NewPool: %w", err)
	}

	return p, nil
}

// CreateEvent ...
func (s *Storage) CreateEvent(ctx context.Context, event storage.Event) error {
	if m, ok := storage.EnoughData(event); !ok {
		// m := strings.Join(m, ", ")
		return fmt.Errorf("not enough data: %+v", m)
	}

	// var id int64
	// err := s.pgpool.QueryRow(ctx, `SELECT id FROM events WHERE id=$1`, event.ID).Scan(&id)
	// if err != nil {
	// 	return err
	// }

	// if id != 0 {
	// 	e := fmt.Sprintf("exists event with id=%d", event.ID)
	// 	return errors.New(e)
	// }

	sql := `INSERT INTO events (id, title, starttime, endtime, descript, userid) 
				VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.pgpool.Exec(ctx, sql,
		event.ID,
		event.Title,
		event.StartTime,
		event.EndTime,
		event.Descript,
		event.User,
	)
	if err != nil {
		return fmt.Errorf("failed create sql event: %w", err)
	}

	return nil
}

// UpdateEvent ...
func (s *Storage) UpdateEvent(ctx context.Context, event storage.Event) error {
	if m, ok := storage.EnoughData(event); !ok {
		// m := strings.Join(m, ", ")
		return fmt.Errorf("not enough data: %+v", m)
	}

	sql := `UPDATE events SET (title, starttime, endtime, descript, userid) = 
				($1, $2, $3, $4, $5) WHERE id=$6`
	_, err := s.pgpool.Exec(ctx, sql,
		event.Title,
		event.StartTime,
		event.EndTime,
		event.Descript,
		event.User,
		event.ID,
	)
	if err != nil {
		return fmt.Errorf("failed update sql event: %w", err)
	}

	return nil
}

// DeleteEvent ...
func (s *Storage) DeleteEvent(ctx context.Context, event storage.Event) error {
	if event.ID == 0 {
		return fmt.Errorf("cannot be id=%d", event.ID)
	}

	var id int64
	err := s.pgpool.QueryRow(ctx, `SELECT id FROM events WHERE id=$1`, event.ID).Scan(&id)
	if err != nil {
		return fmt.Errorf("failed delete sql event: %w", err)
	}

	if id == 0 {
		return fmt.Errorf("no such event id=%d", event.ID)
	}

	sql := `DELETE FROM events WHERE id=$1`
	_, err = s.pgpool.Exec(ctx, sql,
		event.ID,
	)
	if err != nil {
		return fmt.Errorf("failed delete sql event: %w", err)
	}

	return nil
}

// GetEvents ...
func (s *Storage) GetEvents(ctx context.Context, start, end time.Time) (result []storage.Event, err error) {
	sql := `SELECT * FROM events WHERE startTime >=$1 AND startTime <=$2`

	rows, err := s.pgpool.Query(ctx, sql,
		start,
		end,
	)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("failed get sql events: %w", err)
	}

	for {
		r, err := nextRow(rows)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed get sql events: %w", err)
		}
		result = append(result, *r)
	}

	return
}

func nextRow(rows pgx.Rows) (event *storage.Event, err error) {
	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}
		return nil, io.EOF
	}

	r := storage.Event{}
	if err := rows.Scan(&r.ID, &r.Title, &r.StartTime, &r.EndTime, &r.Descript, &r.User); err != nil {
		return nil, err
	}

	return &r, nil
}
