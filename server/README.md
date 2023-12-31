# Pocket Mate Server

Pocket Mate backend(server) using Go.

# Getting Started

Create .env file:

```
cp .env.sample .env
```

## Run

```bash
make run
```

## Test

```bash
make test
```

To run only short test (unit test)

```bash
make test-short

```

# Database

## Create migrations

Create migration using [migrate CLI](https://github.com/golang-migrate/migrate). Here is an example.

```
migrate create -ext sql -dir ./migration -seq create_trip_table
```

Once you run the command, you should have up/down files.

```
0000001_create_trip_table.up.sql
0000001_create_trip_table.down.sql
```

- `0000001_create_trip_table.up.sql` should contain instructions that you are trying to achieve.
- `0000001_create_trip_table.down.sql` should contain instructions to remove whatever change is made in the `up`.

## Run migrations (docker-compose)

```
migrate -path ./migration -database "mysql://root:password@tcp(localhost:3305)/pm" up
```

## Rollback migrations

```
migrate -path ./migration -database "mysql://root:password@tcp(localhost:3305)/pm" down
```
