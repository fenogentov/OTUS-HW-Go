package storage

import (
	"time"
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
