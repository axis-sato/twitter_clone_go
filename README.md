# twitter_clone_go

## Getting Started

Start the db container.

```bash
docker-compose up -d
```

Run migrations.

```bash
goose up
```

Run seeding.

```bash
go run db/seeds/*.go
```

Generate models

```bash
go generate
```
