# Aplikasi 4cRypto

## Getting Source
Download the source code from gitlab to a local folder of your choice by executing:

	git clone https://git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-17-golang/final-project/group-3/4cRypto.git

## Build Environment
In order to build 4cRypto, you will need the following tools installed in your system:

* **Go** (recent version) - http://golang.org/doc/install
* **Git** (optional) - http://git-scm.com/downloads
    
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

Make database with db name:db4crypto and create database with 4cRypto_ddl.sql and 4cRypto_dml.sql

* **Login**
Login functions to get a jwt authorization token.


## How to Run the Application

### User API

#### Login

Request : - Method : `POST`
- Endpoint : `/api/v1`
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
