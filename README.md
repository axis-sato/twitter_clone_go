# twitter_clone_go

## Getting Started

Start the db container.

```bash
docker-compose up -d
docker-compose ps
         Name                       Command               State                 Ports
---------------------------------------------------------------------------------------------------
twitter_clone_db         docker-entrypoint.sh --def ...   Up      0.0.0.0:3307->3306/tcp, 33060/tcp
twitter_clone_go_app_1   realize start                    Up      0.0.0.0:1323->1323/tcp
```

Run migrations.

```bash
docker-compose exec app sql-migrate up -config=db/dbconf.yml
```

Run seeding.

```bash
docker-compose exec app go run db/seeds/*.go
```

Generate models

```bash
docker-compose exec app go generate
```
