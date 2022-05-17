package storage

import (
	"time"
)

// Event ...
type Event struct {
	ID        int64     `json:"id"`       // uuid.UUID //- уникальный идентификатор события (можно воспользоваться UUID);
	Title     string    `json:"title"`    // * Заголовок - короткий текст;
	StartTime time.Time `json:"start"`    // * Дата и время события;
	EndTime   time.Time `json:"end"`      // * Длительность события (или дата и время окончания);
	Descript  string    `json:"descript"` // * Описание события - длинный текст, опционально;
	User      string    `json:"user"`     // * ID пользователя, владельца события;
}
