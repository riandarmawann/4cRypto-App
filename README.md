# Aplikasi 4cRypto

## Getting Source
Download the source code from gitlab to a local folder of your choice by executing:

	git clone https://git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-17-golang/final-project/group-3/4cRypto.git

## Build Environment
In order to build 4cRypto, you will need the following tools installed in your system:

* **Go** (recent version) - http://golang.org/doc/install

* **Git** (optional) - http://git-scm.com/downloads

* **Ganache** (optional) - https://trufflesuite.com/ganache/

* **Packages**

  * github.com/gin-gonic/gin

  * github.com/golang-jwt/jwt/v5

  * github.com/joho/godotenv

  * github.com/lib/pq

  * github.com/stretchr/testify
    
## Features

Features available in 4cRypto :

### Feature Application
* **Testing**

  Testing functions so that existing code runs as expected.

  For run a testing can use 


	  go test ./... -coverprofile cover.out && go tool cover -html=cover.out
  
  or


	  go test ./... -coverprofile cover.out; go tool cover -html cover.out

* **Authorization**

  Authorization functions to provide access to users using jwt tokens for a moment.

* **Logger**

  Logger functions to record every existing request.

* **Database**

  Database functions to store and manage data.

  Make database with database name:db4crypto and create database with 4cRypto_ddl.sql and 4cRypto_dml.sql

### Feature User

* **Login**

  Login functions to get a jwt authorization token.

* **Refresh Token**

  Refresh Token to got new Token if the last Token has been expired

* **Create/Register User**

  Register functions to create new user

* **Update User**

  Update functions to update user 

* **Delete User**

  Delete functions to delete user

* **GetByID**

  GetByID functions to get user information by id

### Feature Crypto

* **Bid**

  Bid function to user can bid some coin for the high price

* **Ask**

  Ask function to user can ask some coin for the lower price

* **Cancel Order**

  Cancel Order Function to user can cancel the order from orderbook

* **Matching Engine**

  The cryptocurrency exchange matching engine is software that decentralised exchanges and brokerage companies use to fulfil market orders.

* **OrderBook**

  In the context of cryptocurrency, the term "order books" refers to the record of buy and sell orders for a particular cryptocurrency or trading pair on an exchange platform.

* **Ganache**

  Ganache, as provided by the Truffle Suite, is a popular tool for Ethereum development that offers a personal blockchain for Ethereum development purposes. It provides a local Ethereum blockchain environment that developers can use to deploy contracts, develop applications, and run tests. 

## How to Run the Application

### User API

#### Login

Request : 
- Method : `POST`
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

#### Refresh Token

Request : 
- Method : `GET`
- Endpoint : `/api/v1/auth/refresh-token`

Response :

- Status : 201 Created
- Body :

```json
{
    "Status": {
        "Code"           : 201,
        "Description"    : "Successfully refresh token",
    },
    "Data": {
        "Token" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
    }
}
```

#### Create/Register

Request : 
- Method : `POST`
- Endpoint : `/users`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :
```json
{
  "name": "string",
  "email": "string",
  "username": "string",
  "password": "string",
  "role": "string"
}
```

Response :

- Status : 201 Created
- Body :

```json
{
    "Status": {
        "Code"           : 201,
        "Description"    : "Created",
    },
    "Data": {
      "name": "string",
      "email": "string",
      "username": "string",
      "password": "string",
    }
}
```

#### Update

Request : 
- Method : `PUT`
- Endpoint : `/users/update/:id`


Response :

- Status : 200 OK
- Body :

```json
{
    "Status": {
        "Code"           : 200,
        "Description"    : "Updated",
    },
      "Data": {
        "id": "string",
        "name": "string",
        "email": "string",
        "username": "string",
        "password": "string",
        "role": "string",
        "created_at": "time.Time",
        "updated_at": "time.Time",
    }
}
```

#### Delete

Request : 
- Method : `DELETE`
- Endpoint : `/users/delete/:id`


Response :

- Status : 200 OK
- Body :

```json
{
    "Status": {
        "Code"           : 200,
        "Description"    : "Success",
    },
    "Data": {
        "Data": {
          ""
    }
    }
}
```

#### GetById

Request : 
- Method : `GET`
- Endpoint : `/users/:id`


Response :

- Status : 200 OK
- Body :

```json
{
    "Status": {
        "Code"           : 200,
        "Description"    : "",
    },
    "Data": {
          "id": "string",
          "name": "string",
          "email": "string",
          "username": "string",
          "password": "string",
          "role": "string",
          "created_at": "time.Time",
          "updated_at": "time.Time",
    }
}
```

### Crypto API

#### Bids

Request : 
- Method : `POST`
- Endpoint : `/order`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :
```json
{
  "bids": "bool(true)",
  "size" : "int",
  "price": "string",
  "limit": "LIMIT",
  "market": "MARKET",
  "timestamp": "time.Time"
}
```

Response :

- Status : 200 OK
- Body :

```json
{
    "Status": {
        "Code"           : 200,
        "Description"    : "Successfully Order",
    },
    "Data": {
        "msg" : "Success"
    }
}
```

#### Asks

Request : 
- Method : `POST`
- Endpoint : `/order`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :
```json
{
  "bids": "bool(true)",
  "size" : "int",
  "price": "string",
  "limit": "LIMIT",
  "market": "MARKET",
  "timestamp": "time.Time"
}
```

Response :

- Status : 200 OK
- Body :

```json
{
    "Status": {
        "Code"           : 200,
        "Description"    : "Successfully Order",
    },
    "Data": {
        "msg" : "Success"
    }
}
```

#### OrderBook

Request : 
- Method : `GET`
- Endpoint : `/book/ETH`

Response :

- Status : 200 OK
- Body :

```json
{
    "Status": {
        "Code"           : 200,
        "Description"    : "Successfully Get Data",
    },
    "Data": {
        "Ask" : {
          "price": "int",
          "size": "int",
          "bid": "bool(false)",
          "timestamp": "time.Time"
        },
        "Bids": {
          "price": "int",
          "size": "int",
          "bid": "bool(false)",
          "timestamp": "time.Time"
        }
    }
}
```

#### Cancel Order

Request : 
- Method : `DELETE`
- Endpoint : `/order/:id`


Response :

- Status : 200 OK
- Body :

```json
{
    "Status": {
        "Code"           : 200,
        "Description"    : "Success",
    },
    "Data": {
        "Data": {
          ""
    }
    }
}
```
