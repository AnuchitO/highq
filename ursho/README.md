## Synopsis

#### URL Shortener Service

## Code Example

```
# execute the program
go run main.go
```

Using CURL

Generate shortener\
`curl -H "Content-Type: application/json" -X POST -d '{"url":"www.google.com"}' http://localhost:8080/encode/`

Response:
`{"success":true,"response":"http://localhost:8080/bN"}`

Redirect\
Open url in your browser [http://localhost:8080/bN](http://localhost:8080/bN)

OR\
`curl http://localhost:8080/bN

Get info for url shortener\
`curl http://localhost:8080/info/bN `

Response:

```json
{
  "success": true,
  "response": {
    "url": "www.google.com",
    "visited": true,
    "count": 1
  }
}
```

## Motivation

..

## Installation

You can install it using 'go get' or cloning the repository.

#### Select method of persistence

select the method of persistence, in which you going to work.\
`storage := &storages.Postgres{}`

If selected Postgresql as Storage, create database

```sql
CREATE DATABASE ursho_db;
```

Add your config for the method of persistence and other options in file config.json\

```json
{
  "server": {
    "host": "0.0.0.0",
    "port": "8080"
  },
  "options": {
    "prefix": "http://localhost:8080/"
  },
  "posgres": {
    "user": "postgres",
    "password": "mysecretpassword",
    "db": "ursho_db"
  }
}
```

## API Reference

..

## Tests

```
go test -v ./...
```

## License

A short snippet describing the license (MIT, Apache, etc.)
