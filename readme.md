# API DASHBOARD 409

Backend system for SIEM dashboard of PBL409.

## Project Directory Structure

```
PBL409/
├── .env
├── config
│   └── database_connection.go
├── database
│   ├── docker-compose.yml
│   ├── migration
│   │   ├── migrateDatabase.go
│   │   ├── service_migration.go
│   │   └── user_migration.go
│   └── seeder
│       ├── service_seeder.go
│       └── user_seeder.go
├── dtos
│   └── service.go
├── go.mod
├── go.sum
├── handler
│   └── service_handler.go
├── main.go
├── middleware
│   └── middleware.go
├── models
│   ├── service.go
│   └── user.go
├── repositories
│   ├── service_repository.go
│   └── user_repository.go
├── router
│   └── router.go
├── services
│   ├── service_service.go
│   └── user_service.go
└── utils
    ├── response.go
    └── validate_id.go
```

## Requirement

1. Make sure you have Go installed on your PC [Click this to download]("https://go.dev/doc/install")
2. Also you need docker to run the database or you can install the database manually.

## How to run it?

1. Pull the project

```
git pull https://github.com/spexf/pbl409-dashboard
```

2. Copy this and put on `.env` file in the project root.

```env
DATABASE_USER="dbuser"
DATABASE_HOST="127.0.0.1"
DATABASE_PASSWORD="dbpassword"
DATABASE_NAME="dbname"
DATABASE_PORT="5432"
DATABASE_SSL="disable"
DATABASE_TIMEZONE="dbtimezone"

# DOCKER COMPOSE ENV
POSTGRES_USER="dbuser"
POSTGRES_PASSWORD="dbpassword"
POSTGRES_DB="dbname"
```

3. Run the database by change your directory to database directory, and then run `docker compose up -d`.
4. Run `go mod tidy` to install the package that needed to run this project.
5. Run the project with `go run main.go`.

### Send feedback for this readme

`discord : athkr`
