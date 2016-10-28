# TODO API

Simple TODO API with JWT Authentication


## API Documentation

Current end-point: https://todo-api.falmar.com.ve

Rrequests and Responses will always be Content-Type: application/json

## Login

```
POST end-point/login/
Accept: application/json
Content-Type: application/json
```

** Request body **

```json
{
	"email": "weird@user.com",
	"password": "super-password"
}
```

**Response body**

possibles responses codes:

| Code | reponse | description |
|------|---------|-------------|
| 200  | message, claims, token | successfully created user|
| 404 | message | user does not exist |
| 401 | message | authentication failed |
| 400  | message | content-type is not application/json, request body empty, missing field |
| 500  | message | - |

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

I'm returning the claims as well because probably the client application needs some of them basically the user and scope


## User

#### Register

```
POST end-point/user/
Accept: application/json
Content-Type: application/json
```

** Request body **

```json
{
	"name": "some weird user",
	"email": "weird@user.com",
	"password": "super-password"
}
```

**Response body**

possibles responses codes:

| Code | reponse | description |
|------|---------|-------------|
| 201  | message,user | successfully created user|
| 400  | message | content-type is not application/json, request body empty, missing field |
| 500  | message | - |

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
