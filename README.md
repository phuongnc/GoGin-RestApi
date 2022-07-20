
  

# Golang CRUD Api

  

A Golang REST API built with Go Gin, connected to SQLite Database. The main functions:

Api for login and authentication:

- Registration

- Login

Api CRUD shapes of type with types:

- Triangle

- Rectangle

- Square

- Diamond

  

Api request for shapes:

- Area

- Perimeter

  
  

## System

  

---

  

This repository was developed with the system versions below:

  

- Go 1.18

- MacOS (or Ubuntu 18.x)

  

## How to use

---
### Run from source code

Clone this repository to your local machine. Ensure that you have Go environment. Go to root project and run:
```
go run cmd/entity-server/main.go
```
### Run by docker

From root project run the script:
```
sh ./scripts/run.dev.sh
```

## Api Specification
Please import postman collection file to view the api: [shapeservice.postman_collection](https://github.com/phuongnc/GoGin-RestApi/blob/master/docs/shapeservice.postman_collection.json)

Sample request:
```
curl --location --request POST 'localhost:7171/api/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
	"email": "nguyencongphuong@gmail.com",
    "password": "Password@123",
    "first_name": "Nguyen",
    "last_name": "Phuong"
}'
```
## Questions / Feedbacks / Bugs

  

---

  

Feel free to reach out to me if you have any questions or feedback on how my code can be improved.

My email:  nguyencongphuong@gmail.com


  

### TODO

  

- [x] REST APIs

- [x] Docker build

- [ ] Unit test

- [ ] Swagger documentation
