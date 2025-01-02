<hr />

## Project Overview
This project is test challenges as a backend, a RESTful API designed for an library application. The following features are currently under development:

- Demonstrates the use of all HTTP methods: GET, POST, PUT, PATCH, and DELETE.

- Utilizes an ORM to define and manage relationships between tables, ensuring efficient data storage and retrieval.

- Implements database transactions for scenarios involving multiple data inserts.

- Provides a unified approach to error management to simplify debugging and enhance application stability.

- Adopts a project structure pattern designed for maintainability and scalability.

## Running the project

- Make sure you have docker installed.
- Copy `.env.example` to `.env`
- Run `docker-compose up -d`
- Go to `localhost:5000` to verify if the server works.

If you are running without docker be sure database configuration is provided in `.env` file and run `go run . app:serve`

#### Environment Variables

<details>
    <summary>Variables Defined in the project </summary>

| Key            | Value                    | Desc                                        |
| -------------- | ------------------------ | ------------------------------------------- |
| `SERVER_PORT`  | `5000`                   | Port at which app runs                      |
| `ENV`          | `development,production` | App running Environment                     |
| `LOG_OUTPUT`   | `./server.log`           | Output Directory to save logs               |
| `LOG_LEVEL`    | `info`                   | Level for logging (check lib/logger.go:172) |
| `DB_USER`      | `username`               | Database Username                           |
| `DB_PASS`      | `password`               | Database Password                           |
| `DB_HOST`      | `0.0.0.0`                | Database Host                               |
| `DB_PORT`      | `3306`                   | Database Port                               |
| `DB_NAME`      | `test`                   | Database Name                               |
| `JWT_SECRET`   | `secret`                 | JWT Token Secret key                        |
| `ADMINER_PORT` | `5001`                   | Adminer DB Port                             |
| `DEBUG_PORT`   | `5002`                   | Port that delve debugger runs in            |

</details>

#### Migration Commands

> if you want to run the migration runner `sql-migrate up` from the host environment.

<details>
    <summary>Migration commands available</summary>

| Command             | Desc                                           |
| ------------------- | ---------------------------------------------- |
| `sql-migrate up`    | runs migration up command                      |
| `sql-migrate dqon`  | runs migration down command                    |
| `sql-migrate status`| Check state migration                          |

</details>

## Implemented Features

- Dependency Injection (go-fx)
- Routing (gin web framework)
- Environment Files
- Logging (file saving on `production`) [zap](https://github.com/uber-go/zap)
- Middlewares (cors)
- Database Setup (mysql)
- Models Setup and Automigrate (gorm)
- Repositories
- Implementing Basic CRUD Operation
- Authentication (JWT)
- Migration Runner Implementation
- Live code refresh
- Dockerize Application with Debugging Support Enabled. Debugger runs at `5002`. Vs code configuration is at `.vscode/launch.json` which will attach debugger to remote application. [Learn More](https://medium.com/wesionary-team/docker-debug-environment-for-go-and-gin-framework-36df80e061ac?source=friends_link&sk=35c9d856852944083dd30059200d87f0)
- Cobra Commander CLI Support. try: `go run . --help`

## Contributing

- At this stage, contributions are not required.
