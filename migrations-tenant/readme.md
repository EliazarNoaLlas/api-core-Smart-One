## install
go install github.com/pressly/goose/v3/cmd/goose@latest

## add migration (example)
goose create create_tenant_settings sql

## migrate up
make migrate-tenant-up

## migrate down
make migrate-tenant-down
