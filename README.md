# Aplikasi 4cRypto

## Getting Source
Download the source code from gitlab to a local folder of your choice by executing:

git clone https://git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-17-golang/final-project/group-3/4cRypto.git

## Features
Features available in 4cRypto :
1. Testing
Testing functions so that existing code runs as expected
2. Authorization
Authorization functions to provide access to users using jwt tokens
3. Logger
The logger functions to record every existing request
4. Login
Login functions to get a jwt authorization token

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
