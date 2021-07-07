// package app

// import (
// 	"time"

// 	"github.com/fenogentov/OTUS-HW-Go/hw12_13_14_15_calendar/internal/storage"
// )

// // App ...
// type App struct { // TODO
// 	storage *Storage
// }

// // Logger ...
// type Logger interface {
// 	Info(msg string)
// 	Error(msg string)
// 	Warn(msg string)
// 	Debug(msg string)
// }

// // Storage ...
// type Storage interface { // TODO
// 	CreateEvent(evnt storage.Event)     // добавление события в хранилище;
// 	UpdateEvent(evnt storage.Event)     // изменение события в хранилище;
// 	DeleteEvent(evnt storage.Event)     // удаление события из хранилища;
// 	GetEvents(startDT, endDT time.Time) // листинг событий;
// }

// // New ...
// func New(logger Logger, storage *Storage) *App {
// 	return &App{}
// }

// // CreateEvent ...
// func (a *App) CreateEvent(e storage.Event) error {
// 	// TODO
// 	return nil
// 	// return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
// }

// // UpdateEvent ...
// func (a *App) UpdateEvent(evnt storage.Event) {

// }

// // DeleteEvent ...
// func (a *App) DeleteEvent(evnt storage.Event) {

// }

// // GetEvents ...
// func (a *App) GetEvents(startDT, endDT time.Time) {

// }

// // TODO
