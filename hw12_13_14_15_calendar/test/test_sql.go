package test

import (
	_ "github.com/jackc/pgx/stdlib"
)

// Event ...
// type Event struct {
// 	ID        int64     // uuid.UUID //- уникальный идентификатор события (можно воспользоваться UUID);
// 	Title     string    // * Заголовок - короткий текст;
// 	StartTime time.Time // * Дата и время события;
// 	EndTime   time.Time // * Длительность события (или дата и время окончания);
// 	Descript  string    // * Описание события - длинный текст, опционально;
// 	UserID    string    // * ID пользователя, владельца события;
// }

// func main() {
// 	connString := "postgres://root:12345@127.0.0.1:5438/calendar"

// 	fmt.Println(connString)
// 	db, err := sqlx.Connect("pgx", connString)
// 	if err != nil {
// 		log.Println("error connect DB")
// 	}
// 	fmt.Println(db)

// 	_, err = db.Exec(
// 		`INSERT INTO events(id, title, descript, userid) VALUES (44, 'fff', 'ggghg', '44')`,
// 	)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// timeStart := time.Now()
// 	// timeEnd := time.Now().Add(time.Hour * 23)
// 	rows, err := db.Queryx(`SELECT * FROM events`)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	for rows.Next() {
// 		e := Event{}
// 		if err := rows.StructScan(&e); err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(e)
// 	}
// }
