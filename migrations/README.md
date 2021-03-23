# DB Migrations

SQL migrations are done with [Golang Migrate](https://github.com/golang-migrate/migrate) installed with `brew install golang-migrate`.

You can create a migration with:
`migrate create -dir migrations/ -ext sql <migration_name>`

You can run the migrations with for example:
`migrate -source file://migrations -database 'postgres://<user>:<password>@127.0.0.1:5432/<database>?sslmode=disable' up`