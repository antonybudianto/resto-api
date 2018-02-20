# resto-api
REST API Backend for RestoHub

## Prerequisites
- Go
- MySQL
  - SQL dump [here](https://github.com/antonybudianto/resto-api/blob/master/files/db/create_table.sql)

## Start manually
```
go run cmd/apiapp/main.go
```

## Start using Docker
If you have Docker and docker compose, you can boot up using:
```
docker-compose up
```

### Get top nearest restaurants
**GET** http://localhost:8000/restaurants?lat=-6.2237107&lng=106.8203731

### Book restaurants
**POST** http://localhost:8000/books/

Request body:
```json
{
	"restaurantId": 1,
	"totalPeople": 2,
	"bookDatetime": "2019-10-12 07:00:00"
}
```

### Adminer
http://localhost:8080/

### Credits
For GCD: https://developers.google.com/maps/articles/phpsqlsearch_v3
