package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Event ...
type Event struct {
	ID        int64     // uuid.UUID //- уникальный идентификатор события (можно воспользоваться UUID);
	Title     string    // * Заголовок - короткий текст;
	StartTime time.Time // * Дата и время события;
	EndTime   time.Time // * Длительность события (или дата и время окончания);
	Descript  string    // * Описание события - длинный текст, опционально;
	UserID    string    // * ID пользователя, владельца события;
}

func main() {
	connString := "postgres://root:12345@127.0.0.1:5432/calendar"

	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Println("error connect DB")
	}
	defer pool.Close()

	// _, err = db.Exec(
	// 	`INSERT INTO events(id, title, descript, userid) VALUES (44, 'fff', 'ggghg', '44')`,
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // timeStart := time.Now()
	// // timeEnd := time.Now().Add(time.Hour * 23)
	rows, err := pool.Query(context.Background(), `SELECT * FROM events`)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

}
