# 🪵 Chop

Simplistic URL shortener that helps chop off extra length from your URLs.

Please note this was just a proof of concept hobby project. If you intend to use any of this code in an actual production application, ensure you switch the seond stage of the Docker build to alpine or copy certs manually. Additionally confer with the security middlewares to ensure they fit your needs.

## Technologies

Chop uses Go alongside a couple of other libraries for better DX:

- Chi: idiomatic and elegant API routing
- SQLC: type-safe SQL querying
- Goose: database migrations

## Developer Setup

To setup your developer environment, you need a few prerequisites:

- Go (1.27.1 is used for Chop): https://go.dev/dl/
- PostgreSQL (I use Docker/K8s but anything should do as long as you can connect to it)
- sqlc: https://sqlc.dev/

Once the prerequisites have been setup on your system, you can start developing:

```sh
git clone https://github.com/usrbinbryce/chop.git
cd chop
go get ./...
```

Make sure to set the required environment variables either in .env (ref .env.example) or via your shell.

Additionally, you should create a `chop` database in your PostgreSQL instance.

### Migrations

Running migrations is simple with Goose:

```sh
goose up/down
```

To create a migration:

```sh
goose -s create my_migration_here sql
```

### Database Queries

Database queries should only be added/removed/edited inside of `./internal/adapters/postgresql/sqlc/queries.sql`. To generate interfaces for queries:

```sh
sqlc generate
```
