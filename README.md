# Azathot

## Requirements

### Config File
Make sure to copy config_dist.json to config.json and fill it with the necessary values. The following fields are numerical, so you would like to remove the quotes and put a number there: crypt_cost jwt_expiration

### Dependencies
There are a couple of dependencies you would like to go get <package> before building the project: golang.org/x/crypto/bcrypt github.com/dgrijalva/jwt-go

## Endpoints
### Protected endpoints
There are some endpoints that are protected under JWT validation. This JWT needs to be sent as a header under the name **Token**. Expect a **401** http status code in case the token was not validated successfully.

### 500 http status code
If something unexpected happens while processing your request, expect a **500** http status code as a response.

### /healthz
Make a **GET** request to this endpoint to check the healthiness of the application, this implies DB connection check as well.

### /signup
Make a **POST** request to this endpoint with the following JSON structure:
```json
{
    "email": "email@domain.com",
    "password": "userPassword"
}
```
Expect to receive the following response on success with a **200** http status code:
```json
"user registered succesfully"
```
Expect a **400** http status code if the email is already registered.

### /login
Make a **POST** request to this endpoint with the following JSON structure:
```json
{
    "email": "email@domain.com",
    "password": "userPassword"
}
```
Expect to receive the following object on success with a **200** http status code:
```json
{
    "message": "login successful",
    "token": "JWT"
}
```

### /players


### /admin/players

## Clone this repository

Use these steps to clone from SourceTree, our client for using the repository command-line free. Cloning allows you to work on your files locally. If you don't yet have SourceTree, [download and install first](https://www.sourcetreeapp.com/). If you prefer to clone from the command line, see [Clone a repository](https://confluence.atlassian.com/x/4whODQ).

1. You’ll see the clone button under the **Source** heading. Click that button.
2. Now click **Check out in SourceTree**. You may need to create a SourceTree account or log in.
3. When you see the **Clone New** dialog in SourceTree, update the destination path and name if you’d like to and then click **Clone**.
4. Open the directory you just created to see your repository’s files.