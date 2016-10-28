# TODO API

Simple TODO API with JWT Authentication


## API Documentation

Current end-point: https://todo-api.falmar.com.ve

Rrequests and Responses will always be Content-Type: application/json

## Login

**Request**
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

**Response**

```
200 Created
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
  "token": "jwt_token"
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

## User

#### Register

**Request**
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

**Response**

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
    "created_at": "2016-10-27T22:19:30.255354812-04:00",
    "updated_at": "2016-10-27T22:19:30.25535506-04:00"
  }
}
```

Possibles responses:

| HTTP Code | reponse | description |
|------|---------|-------------|
| 201  | message,user | successfully created user|
| 400  | message | content-type is not application/json, request body empty, missing field |
| 500  | message | - |

#### Update

**Request**
```
PUT end-point/user/7/
Accept: application/json
Content-Type: application/json
Authorization: Bearer jwt_token
```

```json
{
	"name": "Super Weird",
	"email": "weird@user.com"
}
```

> Note: If additional field password is send, then the password will be changed as well

**Response**

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
    "created_at": "2016-10-27T22:19:30.255354812-04:00",
    "updated_at": "2016-10-28T16:12:31.756548215-04:00"
  }
}
```

Possibles responses:

| HTTP Code | reponse | description |
|------|---------|-------------|
| 200  | message,user | successfully updated user |
| 400  | message | content-type is not application/json, request body empty, missing field |
| 403 | message | forbidden access |
| 404 | message | user does not exist |
| 500  | message | - |

#### Delete

**Request**
```
DELETE end-point/user/:user_id/
Accept: application/json
Content-Type: application/json
Authorization: Bearer jwt_token
```

`no body required`

> Note: it will not delete the current logged user, and is required to have scope user:delete to delete other users

**Response**

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
