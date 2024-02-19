# Aplikasi 4cRypto

## Getting Source
Download the source code from gitlab to a local folder of your choice by executing:

	git clone https://git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-17-golang/final-project/group-3/4cRypto.git

## Build Environment
In order to build 4cRypto, you will need the following tools installed in your system:

* **Go** (recent version) - http://golang.org/doc/install
* **Git** (optional) - http://git-scm.com/downloads

* **GITHUB**

  * github.com/gin-gonic/gin

  * github.com/golang-jwt/jwt/v5

  * github.com/joho/godotenv

  * github.com/lib/pq

  * github.com/stretchr/testify
    
## Features

Features available in 4cRypto :

* **Testing**

  Testing functions so that existing code runs as expected.

  For run a testing can use 


	  go test ./... -coverprofile cover.out && go tool cover -html=cover.out
  
  or


	  go test ./... -coverprofile cover.out; go tool cover -html cover.out


* **Authorization**

  Authorization functions to provide access to users using jwt tokens.

* **Logger**

  Logger functions to record every existing request.

* **Database**

  Database functions to store and manage data.

  Make database with database name:db4crypto and create database with 4cRypto_ddl.sql and 4cRypto_dml.sql

* **Login**

  Login functions to get a jwt authorization token.


## How to Run the Application

### User API

#### Login

Request : - Method : `POST`
- Endpoint : `/api/v1/auth/login`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :
```json
{
  "username": "string",
  "password": "string"
}
```

Response :

- Status : 200 OK
- Body :

```json
{
    "Status": {
        "Code"           : 200,
        "Description"    : "Successfully logged in",
    },
    "Data": {
        "Token" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
    }
}
```
