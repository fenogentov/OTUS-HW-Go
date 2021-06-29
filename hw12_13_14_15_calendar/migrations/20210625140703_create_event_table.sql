-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS events (
    id serial PRIMARY KEY, 
    title varchar(255) NOT NULL, 
    starttime timestamp without time zone, 
    endtime timestamp without time zone, 
    descript text NOT NULL,
    userid varchar(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE event;
-- +goose StatementEnd