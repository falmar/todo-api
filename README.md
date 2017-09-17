# TODO API

Simple TODO API with JWT Authentication

## Index
- [API Documentation](#api-documentation)
	- [Login](#login)
	- [Users](#user)
	- [TODOs](#todo)
- [Local Install](#local-install)


## API Documentation

~Current end-point: http://todo-api.dlavieri.com~
> NOTE: end-point is not longer live

Requests and Responses will always be Content-Type: application/json

## Login

Request
```
POST end-point/login/
Accept: application/json
Content-Type: application/json
```

```json
{
	"email": "weird@user.com",
	"password": "super-password"
}
```

Response

```
200 OK
Content-Type: application/json
```
```json
{
  "claims": {
    "user": {
      "id": 7,
      "name": "some weird user",
      "email": "weird@user.com",
      "created_at": "2016-10-28T02:19:30.255355Z",
      "updated_at": "2016-10-28T02:19:30.255355Z"
    },
    "scope": "todo:create,todo:update,todo:delete"
  },
  "message": {
    "title": "User successfully log in",
    "type": "SUCCESS"
  },
  "token": "{jwt_token}"
}
```

Possibles responses:

| HTTP Code | reponse | description |
|------|---------|-------------|
| 200  | message, claims, token | successfully created user|
| 400  | message | content-type is not application/json, request body empty, missing field |
| 401 | message | authentication failed |
| 404 | message | user does not exist |
| 500  | message | - |

> Note: I'm returning the claims as well because probably the client application needs some of them basically the user and scope

***

## User

### Register

Request
```
POST end-point/user/
Accept: application/json
Content-Type: application/json
```

```json
{
	"name": "some weird user",
	"email": "weird@user.com",
	"password": "super-password"
}
```

Response

```
201 Created
Content-Type: application/json
```
```json
{
  "message": {
    "title": "User successfully created",
    "type": "SUCCESS"
  },
  "user": {
    "id": 7,
    "name": "some weird user",
    "email": "weird@user.com",
    "created_at": "2016-10-27T22:19:30.255354812Z",
    "updated_at": "2016-10-27T22:19:30.25535506Z"
  }
}
```

Possibles responses:

| HTTP Code | reponse | description |
|------|---------|-------------|
| 201  | message, user | successfully created user|
| 400  | message | content-type is not application/json, request body empty, missing field |
| 500  | message | - |

### Update

Request

```
PUT end-point/user/7/
Accept: application/json
Content-Type: application/json
Authorization: Bearer {jwt_token}
```

```json
{
	"name": "Super Weird",
	"email": "weird@user.com"
}
```

> Note: If additional field password is send, then the password will be changed as well

Response

```
200 OK
Content-Type: application/json
```
```json
{
  "message": {
    "title": "User successfully updated",
    "type": "SUCCESS"
  },
  "user": {
    "id": 7,
    "name": "Super Weird",
    "email": "weird@user.com",
    "created_at": "2016-10-27T22:19:30.255354812Z",
    "updated_at": "2016-10-28T16:12:31.756548215Z"
  }
}
```

Possibles responses:

| HTTP Code | reponse | description |
|------|---------|-------------|
| 200  | message, user | successfully updated user |
| 400  | message | content-type is not application/json, request body empty, missing field |
| 403 | message | forbidden access |
| 404 | message | user does not exist |
| 500  | message | - |

### Delete

Request

```
DELETE end-point/user/:user_id/
Accept: application/json
Content-Type: application/json
Authorization: Bearer {jwt_token}
```

`no body required`

> Note: it will not delete the current logged user, and is required to have scope user:delete to delete other users

Response

```
200 OK
Content-Type: application/json
```
```json
{
  "message": {
    "title": "User successfully deleted",
    "type": "SUCCESS"
  }
}
```

Possibles responses:

| HTTP Code | reponse | description |
|------|---------|-------------|
| 200  | message | successfully deleted user |
| 400  | message | content-type is not application/json, request body empty, missing field |
| 403 | message | forbidden access |
| 404 | message | user does not exist |
| 500  | message | - |

***

## TODO

### List

**Filter**

filter format: `filter_name:value`

e.g: `GET end-point/todo/?filters=completed:true,title=first`

Available filters:

- completed

**Sort**

sort format: `updated_at:+|-`

- \+ = descendant [SQL: DESC]
- \- = ascendant [SQL: ASC]

e.g: `GET end-point/todo/?sort=completed:+,title:-`

> NOTE: you need to ENCODE the url, using title:- will turn out as "title: ". Enconding the URL will ensure that it will be the plus symbol encoded such as "title:%2B"

Available sorts:

- id
- title
- completed
- created_at
- updated_at

Request
```
GET end-point/todo/?current_page=1&page_size=5
Accept: application/json
Content-Type: application/json
Authorization: Bearer {jwt_token}
```
`no body required`

Response

```
200 OK
Content-Type: application/json
```
```json
{
  "pagination": {
    "current_page": 1,
    "page_size": 5,
    "pages": [
      {
        "page": 1,
        "link": "localhost:8080/todo/?current_page=1"
      }
    ],
    "results": 1,
    "total_results": 1,
    "total_pages": 1,
    "links": {
      "current": {
        "page": 1,
        "link": "localhost:8080/todo/?current_page=1"
      },
      "next": {
        "page": 1,
        "link": "localhost:8080/todo/?current_page=1"
      },
      "previous": {
        "page": 1,
        "link": "localhost:8080/todo/?current_page=1"
      }
    }
  },
  "todos": [
    {
      "id": 1,
      "user_id": 7,
      "title": "first TODO",
      "completed": false,
      "created_at": "2016-10-29T02:47:02.682465Z",
      "updated_at": "2016-10-29T02:47:02.682465Z",
      "link": "end-point/todo/1/"
    }
  ]
}
```

Possible responses:

| HTTP Code | reponse | description |
|------|---------|-------------|
| 200  | pagination, list | list of todos |
| 403 | message | forbidden access |
| 500  | message | - |

### Get

Request
```
GET end-point/todo/1/
Accept: application/json
Authorization: Bearer {jwt_token}
```
`no body required`

Response

```
200 OK
Content-Type: application/json
```
```json
{
  "todo": {
    "id": 1,
    "user_id": 7,
    "title": "First TODO",
    "completed": false,
    "created_at": "2016-11-04T16:51:28.013889102Z",
    "updated_at": "2016-11-04T16:51:28.013889102Z",
    "link": "end-point/todo/1/"
  }
}
```

Possible responses:

| HTTP Code | reponse | description |
|------|---------|-------------|
| 201  | todo | get a single todo |
| 403 | message | forbidden access |
| 500  | message | - |

### Add

Request
```
POST end-point/todo/
Accept: application/json
Content-Type: application/json
Authorization: Bearer {jwt_token}
```

```json
{
	"title": "First TODO",
	"completed": false
}
```

Response

```
201 Created
Content-Type: application/json
```
```json
{
  "message": {
    "title": "TODO successfully created",
    "type": "SUCCESS"
  },
  "todo": {
    "id": 1,
    "user_id": 7,
    "title": "First TODO",
    "completed": false,
    "created_at": "2016-11-04T16:51:28.013889102Z",
    "updated_at": "2016-11-04T16:51:28.013889102Z",
    "link": "end-point/todo/1/"
  }
}
```

Possible responses:

| HTTP Code | reponse | description |
|------|---------|-------------|
| 201  | message, todo | successfully created todo |
| 400  | message | content-type is not application/json, request body empty, missing field |
| 403 | message | forbidden access |
| 500  | message | - |

### Update

Request
```
PUT end-point/todo/1/
Accept: application/json
Content-Type: application/json
Authorization: Bearer {jwt_token}
```

```json
{
	"title": "First TODO",
	"completed": true
}
```

Response

```
200 OK
Content-Type: application/json
```
```json
{
  "message": {
    "title": "TODO successfully updated",
    "type": "SUCCESS"
  },
  "todo": {
    "id": 46,
    "user_id": 8,
    "title": "First TODO",
    "completed": true,
    "created_at": "2016-11-04T19:26:27.194859Z",
    "updated_at": "2016-11-04T20:54:41.73513Z",
    "link": "localhost:8080/todo/46/"
  }
}
```

Possible responses:

| HTTP Code | reponse | description |
|------|---------|-------------|
| 200  | message, todo | successfully updated todo |
| 400  | message | content-type is not application/json, request body empty, missing field |
| 403 | message | forbidden access |
| 404 | message | todo/user not found |
| 500  | message | - |

### Delete

Request
```
PUT end-point/todo/1/
Accept: application/json
Content-Type: application/json
Authorization: Bearer {jwt_token}
```

`no body required`

Response

```
200 OK
Content-Type: application/json
```
```json
{
  "message": {
    "title": "TODO successfully deleted",
    "type": "SUCCESS"
  }
}
```

Possible responses:

| HTTP Code | reponse | description |
|------|---------|-------------|
| 200  | message | successfully deleted todo |
| 403 | message | forbidden access |
| 404 | message | todo/user not found |
| 500  | message | - |

***

## Local Install

**Requirements**

- Docker >= 1.10
- PostgreSQL >= 9.3
- Golang >= 1.7

### Prepare PostgreSQL

Assuming you are using docker...

```
$ docker pull postgres:9.5
$ docker run --name todo-api-postgres -e POSTGRES_PASSWORD=superpassword -d -p 5432:5432 postgres:9.5
$ docker run -it --rm --link todo-api-postgres:postgres -e PGPASSWORD=superpassword postgres:9.5 psql -h postgres -U postgres
```

Expose port 5432 to connect to postgres if building the api locally

On posgres shell

```
psql (9.5.5)
Type "help" for help.

postgres=# DROP TABLE IF EXISTS public.todo;
postgres=# DROP TABLE IF EXISTS public.user;

postgres=# CREATE TABLE public.user (
  id serial PRIMARY KEY,
  name VARCHAR(45) NOT NULL,
  email VARCHAR(90) UNIQUE NOT NULL,
  password VARCHAR(512) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE,
  updated_at TIMESTAMP WITH TIME ZONE
);

postgres=# CREATE TABLE public.todo (
  id serial PRIMARY KEY,
  user_id int4 NOT NULL,
  title VARCHAR(256) NOT NULL,
  completed BOOL NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE,
  updated_at TIMESTAMP WITH TIME ZONE
);

psotgres=# ALTER TABLE public.todo
ADD CONSTRAINT todo_user_id FOREIGN KEY (user_id) REFERENCES public.user (id);

postgres=# \q
```

Once you quit the postgres shell the container will be inmediatly removed

Now postgres is ready

### Building Golang todo-api on docker

Not recommended if you wanna change the project files and run the API right away

Assuming your current directory is the project folder

```
$ docker build -t your-user/todo-api .
$ docker run --name todo-api -p 8080:80 -d \
--link todo-api-postgres:postgres \
-e PORT=80 \
-e JWT_KEY=secret-key \
-e JWT_ISSUER=anything \
-e DB_HOST=postgres \
-e DB_NAME=postgres \
-e DB_USER=postgres \
-e DB_SCHEMA=public \
-e DB_PASSWORD=superpassword \
-e DB_PORT=5432 \
-e DB_SSLMODE=disable \
your-user/todo-api
```

expose port 8080 to connect to the API

### Building Golang todo-api locally

This is a better option if you want to modify the project and run the API after changes are made

Assuming you have Golang locally installed and your current directory is the project folder

```
$ cp .env.example .env
$ vi .env
```

Edit the .env file that should look similar to this. Change whatever is neccessary

```
# BASIC
PORT=8080
JWT_KEY=secret-key
JWT_ISSUER=anything

# DB - PostgreSQL
DB_HOST=localhost
DB_NAME=postgres
DB_USER=postgres
DB_SCHEMA=test
DB_PASSWORD=superpassword
DB_PORT=5432
DB_SSLMODE=disable
```

Now run the API

```
// get dependencies
$ go get ./

// build and execute
$ go build && ./todo-api

// or install... if you have already setup $GOPATH/bin to your $PATH
$ go install && todo-api
```
