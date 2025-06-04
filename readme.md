# API DASHBOARD 409

Backend system for SIEM dashboard of PBL409.

## Project Directory Structure

```bash
PBL409/
├── .env
├── Dockerfile
├── config
│   └── database_connection.go
├── database
│   ├── migration
│   │   ├── migrateDatabase.go
│   │   ├── service_migration.go
│   │   └── user_migration.go
│   └── seeder
│       ├── service_seeder.go
│       └── user_seeder.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── middleware
│   └── middleware.go
├── pkg
│   ├── auth
│   │   ├── handler.go
│   │   └── service.go
│   ├── services
│   │   ├── dtos.go
│   │   ├── handler.go
│   │   ├── model.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── users
│   │   ├── dtos.go
│   │   ├── handler.go
│   │   ├── model.go
│   │   ├── repository.go
│   │   └── service.go
│   └── utils
│       ├── response.go
│       ├── validate_id.go
│       └── validate_json.go
├── readme.md
└── router
    └── router.go
```

## Requirements

1. Make sure you have Go installed on your PC [Click this to download](https://go.dev/doc/install)
2. Also you need docker to run the database or you can install the database manually.

## How to run it?

1. Pull the project

```bash
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
JWT_SECRET_KEY="JWT KEY"
```

3. use `docker-compose up -d` to run the API

### Send feedback for this project

`discord : athkr`
