# TODO API

Simple TODO API with JWT Authentication


## API Documentation

Current end-point: http://todo-api.dlavieri.com

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
