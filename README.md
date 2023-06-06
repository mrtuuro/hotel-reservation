# Hotel Reservation Backend

## Project Outline
- users -> book from an hotel
- admins -> going to check reservations/bookings
- Authentication and authorization -> JWT Tokens
- Hotels -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> database management -> seeding database/migrating

## Resources
### MongoDB driver

Documentation 
```
https://mongodb.com/docs/drivers/go/current/quick-start
```

Installing mongodb driver
````
go get go.mongodb.org/mongo-driver/mongo
````

### Go fiber
Documentation 
```
https://gofiber.io
```

Installing gofiber
```
go get github.com/gofiber/fiber/v2
```

## Docker
### Installing mongodb as a Docker container
```
docker run --name mongodb -d mongo:latest -p 27017:27017
```