## install
go install github.com/pressly/goose/v3/cmd/goose@latest

## add migration (example)
goose create create_table_name sql

## migrate up
make migrate-up

## migrate down
make migrate-down
